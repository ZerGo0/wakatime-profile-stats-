package wakatime

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

// credits: https://github.com/wakatime/wakatime-cli

const (
	// BaseURL is the base url of the wakatime api.
	BaseURL = "https://api.wakatime.com/api/v1"
	// BaseIPAddrv4 is the base ip address v4 of the wakatime api.
	BaseIPAddrv4 = "143.244.210.202"
	// BaseIPAddrv6 is the base ip address v6 of the wakatime api.
	BaseIPAddrv6 = "2604:a880:4:1d0::2a7:b000"
	// DefaultTimeoutSecs is the default timeout used for requests to the wakatime api.
	DefaultTimeoutSecs = 120
)

// Client communicates with the wakatime api.
type Client struct {
	baseURL string
	client  *http.Client
	// doFunc allows api client options to manipulate request/response handling.
	// default function will be set in constructor.
	//
	// wrapping by api options should be performed as follows:
	//
	//	next := c.doFunc
	//	c.doFunc = func(c *Client, req *http.Request) (*http.Response, error) {
	//		// do something
	//		resp, err := next(c, req)
	//		// do more
	//		return resp, err
	//	}
	doFunc func(c *Client, req *http.Request) (*http.Response, error)
}

// NewClient creates a new Client. Any number of Options can be provided.
func NewClient(baseURL, apiKey string) *Client {
	base64ApiKey := base64.StdEncoding.EncodeToString([]byte(apiKey))

	c := &Client{
		baseURL: baseURL,
		client: &http.Client{
			Transport: NewTransport(),
		},
		doFunc: func(c *Client, req *http.Request) (*http.Response, error) {
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Authorization", "Basic "+base64ApiKey)
			return c.client.Do(req)
		},
	}

	return c
}

// Do executes c.doFunc(), which in turn allows wrapping c.client.Do() and manipulating
// the request behavior of the api client.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.doFunc(c, req)
	if err != nil {
		// don't set alternate host if there's a custom api url
		if !strings.HasPrefix(c.baseURL, BaseURL) {
			return nil, err
		}

		var dnsError *net.DNSError
		if !errors.As(err, &dnsError) {
			return nil, err
		}

		c.client = &http.Client{
			Transport: NewTransportWithHostVerificationDisabled(),
		}

		req.URL.Host = BaseIPAddrv4
		if isLocalIPv6() {
			req.URL.Host = BaseIPAddrv6
		}

		zap.L().Debug("dns error, will retry with host ip", zap.String("host", req.URL.Host), zap.Error(err))

		resp, errRetry := c.doFunc(c, req)
		if errRetry != nil {
			return nil, errRetry
		}

		return resp, nil
	}

	return resp, nil
}

func isLocalIPv6() bool {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:80", BaseIPAddrv4))
	if err != nil {
		zap.L().Warn("failed dialing to detect default local ip address", zap.Error(err))
		return true
	}

	defer func() {
		if err := conn.Close(); err != nil {
			zap.L().Debug("failed to close connection to api wakatime", zap.Error(err))
		}
	}()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.To4() == nil
}

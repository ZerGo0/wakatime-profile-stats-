package wakatime

import "time"

type WakaStats struct {
	Data Data `json:"data"`
}

type BestDay struct {
	Date         string  `json:"date"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type Categories struct {
	Decimal      string  `json:"decimal"`
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type Dependencies struct {
	Decimal      string  `json:"decimal"`
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type Editors struct {
	Decimal      string  `json:"decimal"`
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type Languages struct {
	Decimal      string  `json:"decimal"`
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type Machines struct {
	Decimal       string  `json:"decimal"`
	Digital       string  `json:"digital"`
	Hours         int     `json:"hours"`
	MachineNameID string  `json:"machine_name_id"`
	Minutes       int     `json:"minutes"`
	Name          string  `json:"name"`
	Percent       float64 `json:"percent"`
	Text          string  `json:"text"`
	TotalSeconds  float64 `json:"total_seconds"`
}

type OperatingSystems struct {
	Decimal      string  `json:"decimal"`
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type Projects struct {
	Decimal      string  `json:"decimal"`
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type Data struct {
	BestDay                                         BestDay            `json:"best_day"`
	Categories                                      []Categories       `json:"categories"`
	CreatedAt                                       time.Time          `json:"created_at"`
	DailyAverage                                    float64            `json:"daily_average"`
	DailyAverageIncludingOtherLanguage              float64            `json:"daily_average_including_other_language"`
	DaysIncludingHolidays                           int                `json:"days_including_holidays"`
	DaysMinusHolidays                               int                `json:"days_minus_holidays"`
	Dependencies                                    []Dependencies     `json:"dependencies"`
	Editors                                         []Editors          `json:"editors"`
	End                                             time.Time          `json:"end"`
	Holidays                                        int                `json:"holidays"`
	HumanReadableDailyAverage                       string             `json:"human_readable_daily_average"`
	HumanReadableDailyAverageIncludingOtherLanguage string             `json:"human_readable_daily_average_including_other_language"`
	HumanReadableRange                              string             `json:"human_readable_range"`
	HumanReadableTotal                              string             `json:"human_readable_total"`
	HumanReadableTotalIncludingOtherLanguage        string             `json:"human_readable_total_including_other_language"`
	ID                                              string             `json:"id"`
	IsAlreadyUpdating                               bool               `json:"is_already_updating"`
	IsCached                                        bool               `json:"is_cached"`
	IsCodingActivityVisible                         bool               `json:"is_coding_activity_visible"`
	IsIncludingToday                                bool               `json:"is_including_today"`
	IsOtherUsageVisible                             bool               `json:"is_other_usage_visible"`
	IsStuck                                         bool               `json:"is_stuck"`
	IsUpToDate                                      bool               `json:"is_up_to_date"`
	IsUpToDatePendingFuture                         bool               `json:"is_up_to_date_pending_future"`
	Languages                                       []Languages        `json:"languages"`
	Machines                                        []Machines         `json:"machines"`
	ModifiedAt                                      time.Time          `json:"modified_at"`
	OperatingSystems                                []OperatingSystems `json:"operating_systems"`
	PercentCalculated                               int                `json:"percent_calculated"`
	Projects                                        []Projects         `json:"projects"`
	Range                                           string             `json:"range"`
	Start                                           time.Time          `json:"start"`
	Status                                          string             `json:"status"`
	Timeout                                         int                `json:"timeout"`
	Timezone                                        string             `json:"timezone"`
	TotalSeconds                                    float64            `json:"total_seconds"`
	TotalSecondsIncludingOtherLanguage              float64            `json:"total_seconds_including_other_language"`
	UserID                                          string             `json:"user_id"`
	Username                                        any                `json:"username"`
	WritesOnly                                      bool               `json:"writes_only"`
}

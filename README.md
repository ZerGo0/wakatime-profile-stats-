# Wakatime Profile Stats

This GitHub Action will update your README with your coding stats from [Wakatime](https://wakatime.com/).

## Example

![Wakatime Stats](/assets/preview.png)

## Setup

```yaml
name: Update Readme with Metrics

on:
  schedule:
    - cron: "0 6 * * *"
  workflow_dispatch:
jobs:
  update-readme:
    name: Update Readme with Metrics
    runs-on: ubuntu-latest
    steps:
      - name: wakatime-profile-stats
        uses: ZerGo0/wakatime-profile-stats@main
        with:
          WAKATIME_API_KEY: ${{ secrets.WAKATIME_API_KEY }}
          GH_TOKEN: ${{ secrets.GH_TOKEN }}
```

### Secrets

- `WAKATIME_API_KEY` - **Required**. You can get your API key from [here](https://wakatime.com/settings/account).
- `GH_TOKEN` - **Required**. You can get your GitHub token from [here](https://github.com/settings/tokens). The token should have `repo` scope.

### ReadMe Template

```markdown
<!--START_SECTION:waka-->
<!--END_SECTION:waka-->
```

Anything between `<!--START_SECTION:waka-->` and `<!--END_SECTION:waka-->` will be replaced by the stats. Anything else will remain as it is.

## Development

The project uses `make` to make your life easier. If you're not familiar with Makefiles you can take a look at [this quickstart guide](https://makefiletutorial.com).

Whenever you need help regarding the available actions, just use the following command.

```bash
make help
```

### Setup

To get your setup up and running the only thing you have to do is

```bash
make all
```

This will initialize a git repo, download the dependencies in the latest versions and install all needed tools.
If needed code generation will be triggered in this target as well.

### Run

To run the application you can use the following command

```bash
make run
```

You can find all possible arguments above.

### Test & lint

Run linting

```bash
make lint
```

Run tests

```bash
make test
```

## Credits

Made with [![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)

Inspired by [waka-readme-stats](https://github.com/anmol098/waka-readme-stats)

# Pomodoro-cli

Pomodoro command-line interface.

## Examples

```
# Start pomodoro (Default timer is 25mins)
$ pomodoro-cli start
Started pomodoro 13 Jun 2021 23:12. The pomodoro will end 13 Jun 2021 23:13.

# Start pomodoro with duration 30 mins
$ pomodoro-cli start -d 30
Started pomodoro 13 Jun 2021 23:22. The pomodoro will end 13 Jun 2021 23:52.

# Cancel the current pomodoro
$ pomodoro-cli cancel

# Status of the current pomodoro
$ pomodoro-cli cancel
24:51
```

## Config

The config path is `~/.config/pomodoro-cli`. The folder contains

* `config.yaml` the file to config pomodoro-cli.
* `pomodoros.json` contains the current and all previous pomodoros.

### Hooks

If you want to execute a command when the pomodoro finishes then place an executeable script named `finished` in the current path `~/.config/pomodoro-cli`.

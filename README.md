# powerd

## What it is
Simple daemon to handle battery power levels

## How it looks like
```
NAME:
  powerd - Simple daemon to handle battery power levels

USAGE:
  powerd [global options] command [command options] [arguments...]

VERSION:
  1.0.0

COMMANDS:
    help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
  --check-interval value  Sets the interval to check the battery level in seconds (default: 60)
  --warn-on-level value   Sets the battery level which will trigger a warn message (default: 10)
  --warn-message value    Sets the warning message (default: "Battery level at 10%")
  --exec-on-level value   Sets the battery level which will trigger a command execution (default: 5)
  --exec-command value    Sets the exec command (default: "poweroff")
  --help, -h              show help
  --version, -v           print the version
```

## How to build & install

1. `git clone git@github.com:shizonic/powerd.git`

2. `cd powerd && make build`

3. `sudo make install`

## How to uninstall

1. `sudo make uninstall`

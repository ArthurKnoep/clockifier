# clockifier
A small cli tool to import toggl track time entries into clockify

## Usage

```
usage: clockifier [<flags>] <command> [<args> ...]

Clockifier is a tool to import time entries into clockify

Flags:
  -h, --help     Show context-sensitive help (also try --help-long and
                 --help-man).
  -v, --version  Show application version.
  -c, --flag=~/.config/clockifier/clockifier.conf  
                 The path to the configuration file

Commands:
  help [<command>...]
    Show help.

  config init
    Configure the cli to connect with Clockify and Toggl

  config mapping
    Re-configure the project mapping

  import [<flags>]
    Import your time entries from Toggl to Clockify
```

The first time you start the tool you will need to run `clockifier config init` and follow the instruction. If you need to only re-map the project mapping between Toggl and Clockify, you can run `clockifier config mapping`

Only the time entries with a project mapped between Toggl and Clockify will be imported

`clockifier import` has three flags:
- `--at`, will import all time entries for a given date (default to today)
- `--from`, will import all time entries from a given date
- `--to`, will import all time entries until a given date

`at` is the default flag, `from` and `to` can be used together

All date must be given under the form `DD/MM/YYYY`

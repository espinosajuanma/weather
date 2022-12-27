# 🌳 Go Weather Command Line

## Install

You can just grab the latest binary [release](https://github.com/espinosajuanma/weather/releases).

This command can be installed as a standalone program or composed into a Bonzai command tree.

Standalone

```bash
go install github.com/espinosajuanma/weather/cmd/weather@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/espinosajuanma/weather"
)

var Cmd = &Z.Cmd{
	Name:     `z`,
	Commands: []*Z.Cmd{help.Cmd, weather.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C weather weather
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.

## Add Weather to TMUX

Here's an example of how to add `weather` to your TMUX configuration. Your
mileage may vary.

```tmux
set -g status-interval 1
set -g status-right "#(weather)"
```

## Legal

> Data from MET Norway
>
> Norwegian Meteorological Institute - https://www.met.no

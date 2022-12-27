package weather

import (
	"fmt"
	"strconv"
	"time"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

var (
	Emoji = "⛅"
	Unit  = "celsius"
)

func init() {
	Z.Vars.SoftInit()
}

var Cmd = &Z.Cmd{
	Name: `weather`,
	Commands: []*Z.Cmd{
		getCmd,
		help.Cmd, vars.Cmd, conf.Cmd, // common
		updatedCmd,
	},
	Shortcuts: Z.ArgMap{
		`lat`:   {`var`, `set`, `lat`},
		`lon`:   {`var`, `set`, `lon`},
		`agent`: {`var`, `set`, `agent`},
		`emoji`: {`var`, `set`, `emoji`},
		`unit`:  {`var`, `set`, `unit`},
	},
	Version:     `v0.0.2`,
	Source:      `https://github.com/espinosajuanma/weather`,
	Issues:      `https://github.com/espinosajuanma/weather/issues`,
	Summary:     help.S(_weather),
	Description: help.D(_weather),
}

var getCmd = &Z.Cmd{
	Name:        `get`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     help.S(_get),
	Description: help.D(_get),

	Call: func(x *Z.Cmd, args ...string) error {
		lat, _ := x.Caller.Get("lat")
		if len(args) >= 1 {
			lat = args[0]
		}
		_, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			return fmt.Errorf("Please set a correct latitude. Use help command.")
		}

		lon, _ := x.Caller.Get("lon")
		if len(args) >= 1 {
			lon = args[1]
		}
		_, err = strconv.ParseFloat(lon, 64)
		if err != nil {
			return fmt.Errorf("Please set a correct longitude. Use help command.")
		}

		agent, _ := x.Caller.Get("agent")
		if agent == "" {
			return fmt.Errorf("Please set an User-Agent. Use help command.")
		}

		emoji, _ := x.Caller.Get("emoji")
		if emoji == "" {
			emoji = Emoji
		}

		unit, _ := x.Caller.Get("unit")
		celsius := Unit == "celsius"
		if unit == "fahrenheit" {
			celsius = false
		}

		expires := Z.Vars.Get("expires")
		t, _ := time.Parse(time.RFC1123, expires)
		after := time.Now().After(t)

		current := Z.Vars.Get("temp")
		if expires == "" || after {
			req := NewRequest()
			req.SetCoordinates(lat, lon)
			req.SetAgent(agent)

			res, err := req.Get()
			if err != nil {
				return err
			}

			Z.Vars.Set("expires", res.Expires.Format(time.RFC1123))
			Z.Vars.Set("updated", res.LastModified.Format(time.RFC1123))

			current = res.GetFormat(celsius)
			Z.Vars.Set("temp", current)
		}

		fmt.Printf("%s %s\n", emoji, current)

		return nil
	},
}

var updatedCmd = &Z.Cmd{
	Name:        `updated`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     help.S(_updated),
	Description: help.D(_updated),

	Call: func(x *Z.Cmd, args ...string) error {
		updated := Z.Vars.Get("updated")
		if updated == "" {
			return fmt.Errorf("There is no updated weather")
		}

		t, _ := time.Parse(time.RFC1123, updated)
		fmt.Println(t.Local().Format(time.UnixDate))

		return nil
	},
}

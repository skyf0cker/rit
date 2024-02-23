package main

import (
	"github.com/alecthomas/kong"
)

var CLI struct {
	Browse BrowseCmd `cmd:"browse" help:"browse your reddit homepage"`
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Name("rit"),
		kong.Description("Reddit In Terminal"),
		kong.UsageOnError(),
	)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}

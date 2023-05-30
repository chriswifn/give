package server

import (
  _ "embed"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

//go:embed serve.md
var serveDesc string

//go:embed receive.md
var receiveDesc string

var Cmd = &Z.Cmd{
	Name:      `server`,
	Aliases:   []string{`serv`},
	Usage:     `[help]`,
	Summary:   `Serve and receive files`,
	Copyright: `Copyright 2023 Christian Hageloch`,
	Version:   `v0.1.0`,
	License:   `MIT`,
	Source:    `git@github.com:chriswifn/give.git`,
	Issues:    `github.com/chriswifn/give/issues`,
	Commands:  []*Z.Cmd{help.Cmd, serveCmd, serveQrCmd, receiveCmd, receiveQrCmd},
}

var serveCmd = &Z.Cmd{
	Name:        `serve`,
	Aliases:     []string{`send`},
	Usage:       `help`,
	Summary:     `Serve files in current working directory`,
	Copyright:   `Copyright 2023 Christian Hageloch`,
	Version:     `v0.1.0`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/give.git`,
	Issues:      `github.com/chriswifn/give/issues`,
	MaxArgs:     0,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: serveDesc,

	Call: func(x *Z.Cmd, _ ...string) error {
		serve()
		return nil
	},
}

var serveQrCmd = &Z.Cmd{
	Name:        `serveqr`,
	Aliases:     []string{`sendqr`},
	Usage:       `help`,
	Summary:     `Serve files in current working directory`,
	Copyright:   `Copyright 2023 Christian Hageloch`,
	Version:     `v0.1.0`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/give.git`,
	Issues:      `github.com/chriswifn/give/issues`,
	MaxArgs:     0,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: serveDesc,

	Call: func(x *Z.Cmd, _ ...string) error {
		serveQr()
		return nil
	},
}

var receiveCmd = &Z.Cmd{
	Name:        `receive`,
	Aliases:     []string{`get`},
	Usage:       `help`,
	Summary:     `Downloads files to current working directory`,
	Copyright:   `Copyright 2023 Christian Hageloch`,
	Version:     `v0.1.0`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/give.git`,
	Issues:      `github.com/chriswifn/give/issues`,
	MaxArgs:     0,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: receiveDesc,

	Call: func(x *Z.Cmd, _ ...string) error {
		receive()
		return nil
	},
}

var receiveQrCmd = &Z.Cmd{
	Name:        `receiveqr`,
	Aliases:     []string{`getqr`},
	Usage:       `help`,
	Summary:     `Downloads files to current working directory`,
	Copyright:   `Copyright 2023 Christian Hageloch`,
	Version:     `v0.1.0`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/give.git`,
	Issues:      `github.com/chriswifn/give/issues`,
	MaxArgs:     0,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: receiveDesc,

	Call: func(x *Z.Cmd, _ ...string) error {
		receiveQr()
		return nil
	},
}

package give

import (
	"github.com/chriswifn/give/ix"
	"github.com/chriswifn/give/server"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:      `give`,
	Summary:   `Collection of protocols for sharing files`,
	Commands:  []*Z.Cmd{help.Cmd, server.Cmd, ix.Cmd},
	Copyright: `Copyright 2023 Christian Hageloch`,
	Version:   `v0.1.0`,
	License:   `MIT`,
	Source:    `git@github.com:chriswifn/give.git`,
	Issues:    `github.com/chriswifn/give/issues`,
}

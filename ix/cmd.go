package ix

import (
	_ "embed"
	"fmt"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

//go:embed post.md
var postDesc string

//go:embed get.md
var getDesc string

var Cmd = &Z.Cmd{
	Name:      `ix`,
	Aliases:   []string{`ix`},
	Usage:     `[help|PATH]`,
	Summary:   `ix requests`,
	Copyright: `Copyright 2023 Christian Hageloch`,
	Version:   `v0.1.0`,
	License:   `MIT`,
	Source:    `git@github.com:chriswifn/give.git`,
	Issues:    `github.com/chriswifn/give/issues`,
	Commands:  []*Z.Cmd{help.Cmd, postCmd, getCmd},
}

var postCmd = &Z.Cmd{
	Name:        `post`,
	Aliases:     []string{`up`},
	Usage:       `[help|PATH]`,
	Summary:     `ix post request`,
	Copyright:   `Copyright 2023 Christian Hageloch`,
	Version:     `v0.1.0`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/give.git`,
	Issues:      `github.com/chriswifn/give/issues`,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: postDesc,
	Call: func(x *Z.Cmd, args ...string) error {
		buf := Z.ArgsOrIn(args)
		fmt.Print(post(buf))
		return nil
	},
}

var getCmd = &Z.Cmd{
	Name:        `get`,
	Aliases:     []string{`down`},
	Usage:       `[help|PATH]`,
	Summary:     `ix get request`,
	Copyright:   `Copyright 2023 Christian Hageloch`,
	Version:     `v0.1.0`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/give.git`,
	Issues:      `github.com/chriswifn/give/issues`,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: getDesc,
	Call: func(x *Z.Cmd, args ...string) error {
		in := Z.ArgsOrIn(args)
		fmt.Print(get(in))
		return nil
	},
}

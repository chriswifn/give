# 🌳 Go File Sharing Library

Go Bonzai branch for sharing files privately and publicly.
**Please do not use this it desperately needs a rewrite**.

## Installation

You can grab the latest binary [release](https://github.com/chriswifn/give/releases).
This filter command can be installed as a standalone program or composed into
a Bonzai command tree.

Standalone
```
go install github.com/chriswifn/give/cmd/give@latest
```

Composed

```go
package z

import (
    Z "github.com/rwxrob/bonzai/z"
    "github.com/chriswifn/give"
)

var Cmd = &Z.Cmd{
    Name: `z`,
    Commands: []*Z.Cmd{help.Cmd, give.Cmd},
}
```

## Tab Completion
To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C give give
```


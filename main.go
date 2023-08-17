package main

import (
	_ "botfb/internal/logic"

	"botfb/internal/cmd"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var ctx = gctx.New()
	err := cmd.Main.AddCommand(cmd.Version)
	if err != nil {
		g.Log().Fatal(nil, "add command failed")
	}
	cmd.Main.Run(ctx)
}

package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/lingcoder/prometheus-notifier/internal/cmd"
	_ "github.com/lingcoder/prometheus-notifier/internal/logic"
	_ "github.com/lingcoder/prometheus-notifier/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

package main

import (
	_ "github.com/lingcoder/prometheus-notifier/internal/packed"

	_ "github.com/lingcoder/prometheus-notifier/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/lingcoder/prometheus-notifier/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

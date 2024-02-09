package main

import (
	"context"
	"github.com/liuzhaomax/maxblog-user/internal/app"
	"github.com/liuzhaomax/maxblog-user/internal/core"
)

func main() {
	app.Launch(
		context.Background(),
		app.SetConfigFile(core.LoadEnv()),
		app.SetWWWDir("www"),
	)
}

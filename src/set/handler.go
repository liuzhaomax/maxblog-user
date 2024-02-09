package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-user/src/api_user/handler"
)

var HandlerSet = wire.NewSet(
	handler.HandlerUserSet,
)

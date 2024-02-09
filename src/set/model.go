package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-user/src/api_user/model"
)

var ModelSet = wire.NewSet(
	model.ModelUserSet,
)

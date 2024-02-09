package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-user/src/api_user/business"
)

var BusinessSet = wire.NewSet(
	business.BusinessUserSet,
)

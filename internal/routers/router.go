package routers

import (
	"github.com/cnh1en/lets_go/internal/routers/user"
)

type RouterGroup struct {
	User user.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)

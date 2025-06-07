package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/emailplus/api"

var (
	Router   = new(router)
	apiEmail = api.Api.Email
)

type router struct{ Email email }

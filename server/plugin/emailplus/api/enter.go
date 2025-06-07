package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/emailplus/service"

var (
	Api          = new(api)
	serviceEmail = service.Service.Email
)

type api struct{ Email email }

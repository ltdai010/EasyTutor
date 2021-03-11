package middlerware

import (
	"github.com/beego/beego/v2/server/web/context"
	"strings"
)

func filterAdmin(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/v1/easy-tutor/admin/login") {
		return
	}
	username, typeToken := validateToken(ctx.Input.Header("token"))
	if username != "" && typeToken == "admin" {
		ctx.Request.Header.Set("admin", username)
		return
	}
	ctx.ResponseWriter.WriteHeader(403)
}

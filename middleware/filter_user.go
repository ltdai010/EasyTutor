package middleware

import (
	"github.com/beego/beego/v2/server/web/context"
)

func filterUser(ctx *context.Context) {
	username, tokenType := ValidateToken(ctx.Input.Header("token"))
	if username != "" && tokenType == "user"{
		ctx.Request.Header.Set("username", username)
		return
	}
	if ctx.Input.Method() == "PUT" {
		ctx.ResponseWriter.WriteHeader(403)
	}
	return
}

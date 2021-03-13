package middleware

import "github.com/beego/beego/v2/server/web/context"

func filterNotification(ctx *context.Context) {
	username, tokenType := ValidateToken(ctx.Input.Header("token"))
	if username != ""{
		ctx.Request.Header.Set("username", username)
		ctx.Request.Header.Set("user_type", tokenType)
		return
	}
	return
}

package middleware

import "github.com/beego/beego/v2/server/web/context"

func filterComment(ctx *context.Context) {
	if ctx.Input.Method() == "PUT" || ctx.Input.Method() == "POST" || ctx.Input.Method() == "DELETE" {
		username, typeToken := ValidateToken(ctx.Input.Header("token"))
		if username != "" && typeToken == "user" {
			ctx.Request.Header.Set("username", username)
			return
		} else if username != "" && typeToken == "teacher" {
			ctx.Request.Header.Set("teacher_id", username)
			return
		}
		ctx.ResponseWriter.WriteHeader(403)
	}
	return
}


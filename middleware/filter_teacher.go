package middleware

import "github.com/beego/beego/v2/server/web/context"

func filterTeacher(ctx *context.Context) {
	username, typeToken := ValidateToken(ctx.Input.Header("token"))
	if username != "" && typeToken == "teacher" {
		ctx.Request.Header.Set("teacher_id", username)
		return
	}
	if ctx.Input.Method() == "PUT" {
		ctx.ResponseWriter.WriteHeader(403)
	}
	return
}

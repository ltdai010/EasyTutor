package middlerware

import "github.com/beego/beego/v2/server/web/context"

func filterTeacher(ctx *context.Context) {
	if ctx.Input.Method() == "PUT" {
		username, typeToken := validateToken(ctx.Input.Header("token"))
		if username != "" && typeToken == "teacher" {
			ctx.Request.Header.Set("teacher_id", username)
			return
		}
		ctx.ResponseWriter.WriteHeader(403)
	}
	return
}

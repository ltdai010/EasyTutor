package middlerware

import (
	"github.com/beego/beego/v2/server/web/context"
	"log"
)

func filterRequest(ctx *context.Context) {
	if ctx.Input.Method() == "PUT" || ctx.Input.Method() == "POST" || ctx.Input.Method() == "DELETE" {
		username, typeToken := validateToken(ctx.Input.Header("token"))
		log.Println(username)
		log.Println(typeToken)
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

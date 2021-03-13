package middleware

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {
	web.BConfig.WebConfig.Session.SessionOn = true
	web.InsertFilter("/v1/easy-tutor/*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Connection", "Authorization", "Sec-WebSocket-Extensions", "Sec-WebSocket-Key",
			"Sec-WebSocket-Version", "Access-Control-Allow-Origin", "content-type", "Content-Type", "sessionkey", "token", "Upgrade"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type", "Sec-WebSocket-Accept", "Connection", "Upgrade"},
		AllowCredentials: true,
	}))

	web.InsertFilter("/v1/easy-tutor/user/*", web.BeforeRouter, filterUser)
	web.InsertFilter("/v1/easy-tutor/teacher/*", web.BeforeRouter, filterTeacher)
	web.InsertFilter("/v1/easy-tutor/comment/*", web.BeforeRouter, filterComment)
	web.InsertFilter("/v1/easy-tutor/offer/*", web.BeforeRouter, filterOffer)
	web.InsertFilter("/v1/easy-tutor/request/*", web.BeforeRouter, filterRequest)
	web.InsertFilter("/v1/easy-tutor/admin/*", web.BeforeRouter, filterAdmin)

}

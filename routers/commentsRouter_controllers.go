package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["EasyTutor/controllers/admincontroller:AdminController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/admincontroller:AdminController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/admincontroller:AdminController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/admincontroller:AdminController"],
        beego.ControllerComments{
            Method: "GetPage",
            Router: "/unactivated-teacher",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/admincontroller:AdminController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/admincontroller:AdminController"],
        beego.ControllerComments{
            Method: "ActiveUser",
            Router: "/unactivated-teacher/:teacher_id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/commentcontroller:CommentController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/commentcontroller:CommentController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:comment_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/commentcontroller:CommentController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/commentcontroller:CommentController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:comment_id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/commentcontroller:CommentController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/commentcontroller:CommentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:comment_id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/commentcontroller:CommentController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/commentcontroller:CommentController"],
        beego.ControllerComments{
            Method: "PostComment",
            Router: "/teacher/:teacher_id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/notificationcontroller:NotificationController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/notificationcontroller:NotificationController"],
        beego.ControllerComments{
            Method: "Join",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/notificationcontroller:NotificationController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/notificationcontroller:NotificationController"],
        beego.ControllerComments{
            Method: "GetNotification",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/offercontroller:OfferController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/offercontroller:OfferController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:offer_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/offercontroller:OfferController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/offercontroller:OfferController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:offer_id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/offercontroller:OfferController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/offercontroller:OfferController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:offer_id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"],
        beego.ControllerComments{
            Method: "PostRequest",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"],
        beego.ControllerComments{
            Method: "GetPage",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:request_id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:request_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:request_id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"],
        beego.ControllerComments{
            Method: "GetAvailableTeacher",
            Router: "/:request_id/available-teacher",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/:request_id/offer",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"],
        beego.ControllerComments{
            Method: "GetOffer",
            Router: "/:request_id/offer",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/requestcontroller:RequestController"],
        beego.ControllerComments{
            Method: "AcceptOffer",
            Router: "/accepted-offer/:offer_id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/searchcontroller:SearchController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/searchcontroller:SearchController"],
        beego.ControllerComments{
            Method: "SearchRequest",
            Router: "/request",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/searchcontroller:SearchController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/searchcontroller:SearchController"],
        beego.ControllerComments{
            Method: "SearchTeacher",
            Router: "/teacher",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/storagecontroller:StorageController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/storagecontroller:StorageController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"],
        beego.ControllerComments{
            Method: "GetPage",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:teacher_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"],
        beego.ControllerComments{
            Method: "GetAvailableRequest",
            Router: "/:teacher_id/available-request",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"],
        beego.ControllerComments{
            Method: "GetComment",
            Router: "/:teacher_id/comment",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"],
        beego.ControllerComments{
            Method: "ForgotPass",
            Router: "/forgot-password",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"],
        beego.ControllerComments{
            Method: "ResetPass",
            Router: "/reset-pass",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/teachercontroller:TeacherController"],
        beego.ControllerComments{
            Method: "PutSchedule",
            Router: "/schedule",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"],
        beego.ControllerComments{
            Method: "GetPage",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:user_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"],
        beego.ControllerComments{
            Method: "ForgotPass",
            Router: "/forgot-password",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"] = append(beego.GlobalControllerRouter["EasyTutor/controllers/usercontroller:UserController"],
        beego.ControllerComments{
            Method: "ResetPass",
            Router: "/reset-pass",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

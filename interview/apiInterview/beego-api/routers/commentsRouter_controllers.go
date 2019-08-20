package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["test/beego-api/controllers:IntAuthTokenCacheController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:IntAuthTokenCacheController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:IntAuthTokenCacheController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:IntAuthTokenCacheController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:IntAuthTokenCacheController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:IntAuthTokenCacheController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:IntAuthTokenCacheController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:IntAuthTokenCacheController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:IntAuthTokenCacheController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:IntAuthTokenCacheController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:LokcolController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:LokcolController"],
        beego.ControllerComments{
            Method: "TestApi",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:StudentController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:StudentController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:UserController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:UserController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:UserController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:UserController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test/beego-api/controllers:UserController"] = append(beego.GlobalControllerRouter["test/beego-api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

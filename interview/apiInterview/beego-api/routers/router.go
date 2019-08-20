// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/int_auth_token_cache",
			beego.NSInclude(
				&controllers.IntAuthTokenCacheController{},
			),
		),

		beego.NSNamespace("/student",
			beego.NSInclude(
				&controllers.StudentController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	///v{:version}/test-api /v3/test-api
	beego.Router("/v:version([0-9]+)/test-api", &controllers.LokcolController{}, "get:TestApi")

	// /v{:version}/arithmetic/{:action} /v3/arithmetic/plus?a=3&b=4
	beego.Router("/v:version([0-9]+)/arithmetic/:action", &controllers.LokcolController{}, "get,post:ArithmeticAction")

	// /tutorial/student/list
	beego.Router("/tutorial/student/list", &controllers.LokcolController{}, "get:StudentList")

	// /user/sp100032/wallet/self/detail?intAuthToken=yuqbajnnr
	beego.Router("/user/:userId/wallet/self/detail", &controllers.LokcolController{}, "post:AuthToken")
}

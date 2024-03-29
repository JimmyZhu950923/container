// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gt-container-go/controllers"

	"github.com/astaxie/beego/plugins/cors"

	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ProbjectController{},
			),
		),
		beego.NSNamespace("/label",
			beego.NSInclude(
				&controllers.LabelController{},
			),
		),
		beego.NSNamespace("/node",
			beego.NSInclude(
				&controllers.NodeController{},
			),
		),
		beego.NSNamespace("/repositories",
			beego.NSInclude(
				&controllers.RepositoriesController{},
			),
		),
		beego.NSNamespace("/tag",
			beego.NSInclude(
				&controllers.TagController{},
			),
		),
		beego.NSNamespace("/service",
			beego.NSInclude(
				&controllers.ServicesController{},
			),
		),
		beego.NSNamespace("/storage",
			beego.NSInclude(
				&controllers.StorageController{},
			),
		),
		beego.NSNamespace("/config",
			beego.NSInclude(
				&controllers.ConfigController{},
			),
		),
		beego.NSNamespace("/secret",
			beego.NSInclude(
				&controllers.SecretController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/pods",
			beego.NSInclude(
				&controllers.PodsController{},
			),
		),
		beego.NSNamespace("/namespace",
			beego.NSInclude(
				&controllers.NamespaceController{},
			),
		),
		beego.NSNamespace("/cluster",
			beego.NSInclude(
				&controllers.ClusterController{},
			),
		),
		beego.NSNamespace("/wb",
			beego.NSInclude(
				&controllers.WorkBlanceController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

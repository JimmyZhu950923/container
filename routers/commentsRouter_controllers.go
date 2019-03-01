package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"],
        beego.ControllerComments{
            Method: "Select",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:UserController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:UserController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:UserController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

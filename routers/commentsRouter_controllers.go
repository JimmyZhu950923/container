package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"],
        beego.ControllerComments{
            Method: "FindLabels",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"],
        beego.ControllerComments{
            Method: "AddLabel",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"],
        beego.ControllerComments{
            Method: "FindLabelsById",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"],
        beego.ControllerComments{
            Method: "UpdateLabel",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"],
        beego.ControllerComments{
            Method: "DeleteLabel",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:NamespaceController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:NamespaceController"],
        beego.ControllerComments{
            Method: "GetNamespace",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:NamespaceController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:NamespaceController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:NamespaceController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:NamespaceController"],
        beego.ControllerComments{
            Method: "GetSingle",
            Router: `/:name`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:NamespaceController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:NamespaceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:name`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:NodeController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:NodeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:NodeController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:NodeController"],
        beego.ControllerComments{
            Method: "GetByName",
            Router: `/:name`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"],
        beego.ControllerComments{
            Method: "UpdatedPod",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"],
        beego.ControllerComments{
            Method: "DeletedPod",
            Router: `/del`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/getA`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"],
        beego.ControllerComments{
            Method: "GetPodsInNameSpace",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"],
        beego.ControllerComments{
            Method: "NewPod",
            Router: `/newPod`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"],
        beego.ControllerComments{
            Method: "GetSingle",
            Router: `/singlePod`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"],
        beego.ControllerComments{
            Method: "Project",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:pid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:RepositoriesController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:RepositoriesController"],
        beego.ControllerComments{
            Method: "FindResporities",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:RepositoriesController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:RepositoriesController"],
        beego.ControllerComments{
            Method: "DeleteResporities",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:ServicesController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:ServicesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:ServicesController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:ServicesController"],
        beego.ControllerComments{
            Method: "NewS",
            Router: `/newS`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:TagController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:TagController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:TagController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:TagController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:TagController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:TagController"],
        beego.ControllerComments{
            Method: "FindLabels",
            Router: `/label`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:TagController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:TagController"],
        beego.ControllerComments{
            Method: "RemoveLabels",
            Router: `/label`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:TagController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:TagController"],
        beego.ControllerComments{
            Method: "AddLabels",
            Router: `/label`,
            AllowHTTPMethods: []string{"post"},
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
            AllowHTTPMethods: []string{"get"},
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

    beego.GlobalControllerRouter["gt-container-go/controllers:WorkBlanceController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:WorkBlanceController"],
        beego.ControllerComments{
            Method: "GetDeployments",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:WorkBlanceController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:WorkBlanceController"],
        beego.ControllerComments{
            Method: "CreateDeployment",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:WorkBlanceController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:WorkBlanceController"],
        beego.ControllerComments{
            Method: "UpdateDeployment",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:WorkBlanceController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:WorkBlanceController"],
        beego.ControllerComments{
            Method: "DeleteDeployment",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

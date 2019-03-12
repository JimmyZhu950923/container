package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"],
        beego.ControllerComments{
            Method: "AddLabel",
            Router: `/addLabel`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"],
        beego.ControllerComments{
            Method: "DeleteLabel",
            Router: `/deleteLabel`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"],
        beego.ControllerComments{
            Method: "FindLabels",
            Router: `/findLabels`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"],
        beego.ControllerComments{
            Method: "FindLabelsById",
            Router: `/findLabelsById`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:LabelController"],
        beego.ControllerComments{
            Method: "UpdateLabel",
            Router: `/updateLabel`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

<<<<<<< HEAD
=======
    beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:PodsController"],
        beego.ControllerComments{
            Method: "NewS",
            Router: `/newS`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

>>>>>>> new tag.go
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
<<<<<<< HEAD
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
=======
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
>>>>>>> new tag.go
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:ProbjectController"],
        beego.ControllerComments{
<<<<<<< HEAD
            Method: "Put",
            Router: `/:pid`,
            AllowHTTPMethods: []string{"put"},
=======
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
>>>>>>> new tag.go
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:RepositoriesController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:RepositoriesController"],
        beego.ControllerComments{
            Method: "DeleteResporities",
            Router: `/dr`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:RepositoriesController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:RepositoriesController"],
        beego.ControllerComments{
            Method: "FindResporities",
            Router: `/select`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:TagController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:TagController"],
        beego.ControllerComments{
<<<<<<< HEAD
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
=======
            Method: "AddLabels",
            Router: `/addLabels`,
            AllowHTTPMethods: []string{"post"},
>>>>>>> new tag.go
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:TagController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:TagController"],
        beego.ControllerComments{
<<<<<<< HEAD
            Method: "AddLabels",
            Router: `/addLabels`,
            AllowHTTPMethods: []string{"post"},
=======
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"delete"},
>>>>>>> new tag.go
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:TagController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:TagController"],
        beego.ControllerComments{
            Method: "FindLabels",
            Router: `/findLabels`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:TagController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:TagController"],
        beego.ControllerComments{
            Method: "RemoveLabels",
            Router: `/removeLabels`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gt-container-go/controllers:TagController"] = append(beego.GlobalControllerRouter["gt-container-go/controllers:TagController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/select`,
            AllowHTTPMethods: []string{"get"},
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
<<<<<<< HEAD
            AllowHTTPMethods: []string{"get"},
=======
            AllowHTTPMethods: []string{"post"},
>>>>>>> new tag.go
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

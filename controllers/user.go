package controllers

import (
	"crypto/tls"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @router / [get]
func (u *UserController) GetAll() {
}

// @router /login [get]
func (u *UserController) Login() {
	username := u.Input().Get("username")
	password := u.Input().Get("password")

	fmt.Println("-----", username, password, "-----")

	req := httplib.Post("https://kube.gwunion.cn/c/login")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Param("principal", username)
	req.Param("password", password)

	rep, _ := req.Response()
	coks := rep.Cookies()

	for _, cok := range coks {
		if cok.Name == "sid" {
			fmt.Println(cok)
			u.Ctx.Output.Header("X-Token",cok.Value)
			u.Ctx.SetCookie("sid", cok.Value)
		}
	}

	code := rep.StatusCode

	js := map[string]int{"code": code}

	u.Data["json"] = js
	u.ServeJSON()
}

// @router /logout [get]
func (u *UserController) Logout() {

	req := httplib.Get("https://kube.gwunion.cn/c/log_out")
	u.Ctx.SetCookie("sid", "", "MaxAge=-1")
	rep, _ := req.Response()

	code := rep.StatusCode
	js := map[string]int{"code": code}
	u.Data["json"] = js
	u.ServeJSON()
}

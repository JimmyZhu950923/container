package controllers

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

// Operations about object
type RepositoriesController struct {
	beego.Controller
}

//var cookie = http.Cookie{Name: "sid", Value: "50decf9cef14cf68baaead900b57b33f"}

//查询镜像仓库
//@router /select [get]
func (o *RepositoriesController) FindResporities() {
	// cookie, _ := o.Ctx.Request.Cookie("sid")
	// cok, _ := o.Ctx.Request.Cookie("sid")
	page := o.Input().Get("page")
	pageSize := o.Input().Get("page_size")
	fmt.Println("page,pageSize:", page, pageSize)
	fmt.Println("hhhhhhhhhhhhh", cookie)
	url := "https://kube.gwunion.cn/api/repositories"
	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetCookie(&cookie)
	req.Debug(true)
	//定死项目
	req.Param("project_id", "9")
	req.Param("page", page)
	req.Param("page_size", pageSize)
	response, _ := req.Response()
	total, _ := strconv.Atoi(response.Header.Get("X-Total-Count"))
	fmt.Println(response, total)
	result := []map[string]interface{}{}
	err := req.ToJSON(&result)
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	json := map[string]interface{}{"total": total, "result": result}
	o.Data["json"] = json
	o.ServeJSON()
	return
}

//删除镜像
//@router /dr [delete]
func (o *RepositoriesController) DeleteResporities() {
	//cookie, _ := o.Ctx.Request.Cookie("sid")

	path := o.Input().Get("path")
	url := "https://kube.gwunion.cn/api/repositories/" + path
	req := httplib.Delete(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetCookie(&cookie)
	req.Debug(true)
	fmt.Println(req.Response())
	str, err := req.String()
	if err != nil {
		fmt.Println(err, str)
	}
	fmt.Println(str)
	o.Data["json"] = str
	o.ServeJSON()
	return
}

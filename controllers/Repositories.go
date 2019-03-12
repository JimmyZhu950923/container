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

//@Title Get all
//@Description get all repositories
//@Param q query string false "repositories name"
//@Param page query int false "limit page"
//@Param pageSize query int false "limit page_size"
//@Param project_id query int true "project id"
//@Success 200 {object} model.User
//@router / [get]
func (o *RepositoriesController) FindResporities() {

	//cookie1, _ := o.Ctx.Request.Cookie("sid")
	//fmt.Println("-------------", cookie1)
	q := o.Input().Get("q")
	page := o.Input().Get("page")
	pageSize := o.Input().Get("page_size")
	projectId := o.Input().Get("project_id")
	//fmt.Println("hhhhhhhhhhhhh", cookie)
	url := "https://kube.gwunion.cn/api/repositories"
	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	//req.SetCookie(&cookie)
	req.Debug(true)
	//定死项目
	req.Param("project_id", projectId)
	req.Param("q", q)
	req.Param("page", page)
	req.Param("page_size", pageSize)
	response, _ := req.Response()
	if response.StatusCode == 200 {
		total, _ := strconv.Atoi(response.Header.Get("X-Total-Count"))
		result := []map[string]interface{}{}
		err := req.ToJSON(&result)
		if err != nil {
			fmt.Println(err)
			fmt.Println(result)
		}
		json := map[string]interface{}{"total": total, "result": result, "code": 20000}
		o.Data["json"] = json
	} else {
		json := map[string]int{"code": response.StatusCode}
		o.Data["json"] = json
	}
	o.ServeJSON()
	return
}

//@Title Delete
//@Description delete repositories
//@Param path query string true "repositories path"
//@Success 200 {string} 删除成功
//@router / [delete]
func (o *RepositoriesController) DeleteResporities() {
	//cookie, _ := o.Ctx.Request.Cookie("sid")
	path := o.Input().Get("path")
	url := "https://kube.gwunion.cn/api/repositories/" + path
	req := httplib.Delete(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	//req.SetCookie(&cookie)
	req.Debug(true)
	response, _ := req.Response()
	if response.StatusCode == 200 {
		str, err := req.String()
		if err != nil {
			fmt.Println(err, str)
		}
		fmt.Println(str)
		json := map[string]interface{}{"result": str, "code": 20000}

		o.Data["json"] = json

	} else {
		json := map[string]int{"code": response.StatusCode}
		o.Data["json"] = json

	}
	o.ServeJSON()
	return

}

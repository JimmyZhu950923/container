package controllers

import (
	"crypto/tls"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

// Operations about object
type ProbjectController struct {
	beego.Controller
}

func getBody(o *ProbjectController) map[string]interface{} {
	var body map[string]interface{}
	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &body)
	return body
}

// @Title newProject
// @Description add new project
// @Param name body string true "project for name"
// @Param public body bool true "project for public"
// @Success 200 {string} 提交成功
// @router / [post]
func (o *ProbjectController) Add() {
	body := getBody(o)
	name := body["name"]
	var public string
	if body["public"] == true {
		public = "true"
	} else {
		public = "false"
	}

	//fmt.Println("-----", name, public, "-----")

	//cok, _ := o.Ctx.Request.Cookie("sid")
	js := map[string]interface{}{"project_name": name, "metadata": map[string]string{"public": public}}

	req := httplib.Post("https://kube.gwunion.cn/api/projects")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	_, _ = req.JSONBody(js)

	rep, _ := req.Response()
	if rep.StatusCode == 201 {
		o.Data["json"] = map[string]int{"code": 20000}
		o.ServeJSON()
	} else {
		o.Data["json"] = map[string]interface{}{"code": rep.StatusCode, "message": "项目已存在"}
		o.ServeJSON()
	}

	//fmt.Println(rep)

}

// @Title a Project
// @Description select project by id
// @Param id path string false "project for id"
// @Success 200 {object} model.object
// @router /:id [get]
func (o *ProbjectController) Project() {
	id := o.Ctx.Input.Param(":id")

	url := "https://kube.gwunion.cn/api/projects/" + id
	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")

	var json1 map[string]interface{}
	err := req.ToJSON(&json1)

	if err != nil {

	}
	rep, _ := req.Response()
	if rep.StatusCode == 200 {
		rr := map[string]interface{}{"code": 20000, "data": json1}
		o.Data["json"] = rr
		o.ServeJSON()
	} else {
		o.Data["json"] = map[string]int{"code": rep.StatusCode}
		o.ServeJSON()
	}

}

// @Title ProjectList
// @Description select project of list
// @Param name query string false "project for name"
// @Param public query string false "project for public"
// @Param page query string false "current page"
// @Param page_size query string false "page size"
// @Success 200 {object} model.object
// @router / [get]
func (o *ProbjectController) Select() {

	name := o.Input().Get("name")
	public := o.Input().Get("public")
	page := o.Input().Get("page")
	pageSize := o.Input().Get("page_size")

	url := "https://kube.gwunion.cn/api/projects?page=" + page + "&page_size=" + pageSize
	if name != "" {
		url += "&name=" + name
	}
	if public != "" {
		url += "&public=" + public
	}

	//fmt.Println(url)

	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	req.Param("name", name)

	var json1 []map[string]interface{}
	err := req.ToJSON(&json1)

	if err != nil {

	}

	rep, _ := req.Response()

	if rep.StatusCode == 200 {
		total, _ := strconv.Atoi(rep.Header.Get("X-Total-Count"))
		rr := map[string]interface{}{"code": 20000, "data": json1, "total": total}
		o.Data["json"] = rr
		o.ServeJSON()
	} else {
		o.Data["json"] = map[string]int{"code": rep.StatusCode}
		o.ServeJSON()
	}
}

// @Title modifiedProject
// @Description modify project
// @Param pid path string true "project for id"
// @Param public query string false "project for public"
// @Success 200 {string} 修改成功
// @router /:pid [put]
func (o *ProbjectController) Put() {
	id := o.Ctx.Input.Param(":pid")
	public := o.Input().Get("public")

	url := "https://kube.gwunion.cn/api/projects/" + id
	req := httplib.Put(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	js := map[string]interface{}{"metadata": map[string]string{"public": public}}
	_, _ = req.JSONBody(js)
	resp, _ := req.Response()
	//fmt.Println(resp)

	if resp.StatusCode == 200 {
		o.Data["json"] = map[string]int{"code": 20000}
		o.ServeJSON()
	} else {
		o.Data["json"] = map[string]int{"code": resp.StatusCode}
		o.ServeJSON()
	}

}

// @Title deletedProject
// @Description delete project
// @Param id query string true "project for id"
// @Success 200 {string} 删除成功
// @router / [delete]
func (o *ProbjectController) Delete() {

	id := o.Input().Get("id")
	//fmt.Println("-----------", id)

	url := "https://kube.gwunion.cn/api/projects/" + id
	req := httplib.Delete(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	resp, _ := req.Response()
	//fmt.Println(resp)
	if resp.StatusCode == 200 {
		o.Data["json"] = map[string]int{"code": 20000}
		o.ServeJSON()
	} else {
		o.Data["json"] = map[string]int{"code": resp.StatusCode}
		o.ServeJSON()
	}
}

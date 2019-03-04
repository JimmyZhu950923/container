package controllers

import (
	"crypto/tls"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"strconv"
)

// Operations about object
type ProbjectController struct {
	beego.Controller
}


// @router / [post]
func (o *ProbjectController) Add() {
	name := o.Input().Get("name")
	public := o.Input().Get("public")

	fmt.Println("-----", name, public, "-----")

	cok, _ := o.Ctx.Request.Cookie("sid")
	js := map[string]interface{}{"project_name": name, "metadata": map[string]string{"public": public}}

	req := httplib.Post("https://kube.gwunion.cn/api/projects")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetCookie(cok)
	_, _ = req.JSONBody(js)

	rep, _ := req.Response()
	fmt.Println(rep	)

}

// @router / [get]
func (o *ProbjectController) Select() {

	name := o.Input().Get("name")
	public := o.Input().Get("public")
	page := o.Input().Get("page")
	pageSize := o.Input().Get("page_size")

	cok, _ := o.Ctx.Request.Cookie("sid")
	fmt.Println("------",cok)

	url := "https://kube.gwunion.cn/api/projects?page=" + page + "&page_size=" + pageSize
	if name != "" {
		url += "&name=" + name
	}
	if public != "" {
		url += "&public=" + public
	}

	fmt.Println(url)

	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetCookie(cok)
	req.Param("name", name)

	var json []map[string]interface{}
	err := req.ToJSON(&json)

	//fmt.Println(json)



	rep, _ := req.Response()
	fmt.Println(rep)
	total,_ := strconv.Atoi(rep.Header.Get("X-Total-Count"))
	fmt.Println(total)

	rr := map[string]interface{}{"data":json,"total":total}

	if err != nil {
		o.Ctx.WriteString(err.Error())
	} else {
		o.Data["json"] = rr
		o.ServeJSON()
	}
}

// @router / [put]
func (o *ProbjectController) Put() {
	id := o.Input().Get("pid")
	public := o.Input().Get("public")

	cok, _ := o.Ctx.Request.Cookie("sid")

	url := "https://kube.gwunion.cn/api/projects/" + id
	req := httplib.Put(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	js := map[string]interface{}{"metadata": map[string]string{"public": public}}

	_, _ = req.JSONBody(js)

	req.SetCookie(cok)
	fmt.Println(req.Response())

}

// @router / [delete]
func (o *ProbjectController) Delete() {

	id := o.Input().Get("id")
	fmt.Println("-----------", id)

	cok, _ := o.Ctx.Request.Cookie("sid")
	url := "https://kube.gwunion.cn/api/projects/" + id
	req := httplib.Delete(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetCookie(cok)

	fmt.Println(req.Response())

}

func count(){



}
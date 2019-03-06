package controllers

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

// Operations about object
type TagController struct {
	beego.Controller
}

var cookie = http.Cookie{Name: "sid", Value: "5f046761ef87620768e897bb15e84ee6"}

// @router /select [get]
func (c *TagController) Get() {
	//cookie, _ := c.Ctx.Request.Cookie("sid")
	name := c.Input().Get("name")
	fmt.Println(name)
	req := httplib.Get("https://kube.gwunion.cn/api/repositories/venus/nginx/tags?detail=1")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetCookie(&cookie)
	req.Debug(true)
	//req.Param("repo_name", "z")
	//req.Param("tag", "111")
	rep, _ := req.Response()
	fmt.Println(rep)
	// total, _ := strconv.Atoi(rep.Header.Get("X-Total-Count"))
	// fmt.Println(total)

	json := []map[string]interface{}{}
	err := req.ToJSON(&json)
	if err != nil {
		fmt.Println(err)
		fmt.Println(json)
	}
	// json := map[string]interface{}{"total": total, "result": result}
	c.Data["json"] = json
	c.ServeJSON()
	return
}

// @router / [delete]
func (c *TagController) Delete() {
	//cookie, _ := c.Ctx.Request.Cookie("sid")
	name := c.Input().Get("name")
	url := "https://kube.gwunion.cn/api/repositories/venus/nginx/tags/" + name
	req := httplib.Delete(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetCookie(&cookie)
	req.Debug(true)
	//req.Param("repo_name", "z")
	str, err := req.String()
	if err != nil {
		beego.Info(err)
	}
	fmt.Println(str)
	c.Data["json"] = str
	c.ServeJSON()
	return
}

// @router /findLabels [get]
func (c *TagController) FindLabels() {
	url := "https://kube.gwunion.cn/api/labels?scope=p&project_id=3"
	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Debug(true)
	req.SetCookie(&cookie)
	json := []map[string]interface{}{}
	err := req.ToJSON(&json)
	if err != nil {
		fmt.Println(err)
		fmt.Println(json)
	}
	c.Data["json"] = json
	c.ServeJSON()
	return
}

// @router /removeLabels [delete]
func (c *TagController) RemoveLabels() {
	label_id := c.GetString("label_id")
	name := c.GetString("name")
	fmt.Println("-----", label_id, name, "-----")
	url := "https://kube.gwunion.cn/api/repositories/venus/nginx/tags/" + name + "/labels/" + label_id
	req := httplib.Delete(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Debug(true)
	req.SetCookie(&cookie)
	resp, _ := req.Response()
	fmt.Println("------------------\n", resp)
	return
}

// @router /addLabels [post]
func (c *TagController) AddLabels() {
	name := c.Input().Get("name")
	label_id, _ := strconv.Atoi(c.Input().Get("label_id"))
	fmt.Println("-----", label_id, name, "-----")

	url := "https://kube.gwunion.cn/api/repositories/venus/nginx/tags/" + name + "/labels/"
	req := httplib.Post(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	js := map[string]int{"id": label_id}
	req.JSONBody(js)
	req.Debug(true)
	req.SetCookie(&cookie)

	resp, _ := req.Response()

	fmt.Println(resp)
	return
}

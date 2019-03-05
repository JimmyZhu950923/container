package controllers

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

// Operations about object
type LabelController struct {
	beego.Controller
}

//查询labels
//@router /findLabels [get]
func (o *LabelController) FindLabels() {
	url := "https://kube.gwunion.cn/api/labels"
	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Debug(true)
	req.SetCookie(&cookie)
	req.Param("scope", "p")
	req.Param("project_id", "9")
	response, _ := req.Response()
	total, _ := strconv.Atoi(response.Header.Get("X-Total-Count"))
	json := []map[string]interface{}{}
	err := req.ToJSON(&json)
	if err != nil {
		fmt.Println(err)
		fmt.Println(json)
	}
	result := map[string]interface{}{"total": total, "json": json}
	o.Data["json"] = result
	o.ServeJSON()
	return

}

//根据id单个查询labels
//@router /findLabelsById [get]
func (o *LabelController) FindLabelsById() {
	id := o.Input().Get("id")
	url := "https://kube.gwunion.cn/api/labels/" + id
	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Debug(true)
	req.SetCookie(&cookie)
	response, _ := req.Response()
	fmt.Println(response)
	var json map[string]interface{}
	err := req.ToJSON(&json)
	if err != nil {
		fmt.Println(err)
	}
	o.Data["json"] = json
	o.ServeJSON()
	return
}

//添加labels
//@router /addLabel [post]
func (o *LabelController) AddLabel() {
	name := o.GetString("name")
	color := o.GetString("color")
	description := o.GetString("description")
	url := "https://kube.gwunion.cn/api/labels"
	req := httplib.Post(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetCookie(&cookie)
	label := map[string]interface{}{"name": name, "description": description, "color": color, "scope": "p", "project_id": 9}
	_, _ = req.JSONBody(label)
	response, _ := req.Response()
	fmt.Println(response)
	return

}

//修改labels
//@router /updateLabel [put]
func (o *LabelController) UpdateLabel() {
	id := o.GetString("id")
	name := o.GetString("name")
	color := o.GetString("color")
	description := o.GetString("description")
	fmt.Println(id, name, color, description)
	url := "https://kube.gwunion.cn/api/labels/" + id
	req := httplib.Put(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetCookie(&cookie)
	label := map[string]interface{}{"name": name, "description": description, "color": color, "scope": "p", "project_id": 3}
	_, _ = req.JSONBody(label)
	response, _ := req.Response()
	fmt.Println(response)
	return
}

//删除labels
//@router /deleteLabel [delete]
func (o *LabelController) DeleteLabel() {
	id := o.Input().Get("id")
	url := "https://kube.gwunion.cn/api/labels/" + id
	req := httplib.Delete(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetCookie(&cookie)
	req.Debug(true)
	response, _ := req.Response()
	fmt.Println(response)
	return
}

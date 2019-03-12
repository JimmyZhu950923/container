package controllers

import (
	"crypto/tls"
	"fmt"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

// Operations about object
type TagController struct {
	beego.Controller
}

//@Title allTag
//@Description select all tag
//@Param project_name query string true "project's name"
//@Param repo_name query string true "repository's name"
//@Success 200 {string} 查询成功
// @router /select [get]
func (c *TagController) Get() {
	req := httplib.Get("https://kube.gwunion.cn/api/repositories/venus/nginx/tags?detail=1")
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Debug(true)
	rep, _ := req.Response()
	fmt.Println(rep)
	result := []map[string]interface{}{}
	err := req.ToJSON(&result)
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	json := map[string]interface{}{"result": result, "code": 20000}
	c.Data["json"] = json
	c.ServeJSON()
	return
}

//@Title deleteTag
//@Description delete tag for selected
//@Param project_name query string true "project's name"
//@Param repo_name query string true "repository's name"
//@Param tag_name query string true "tag's name"
//@Success 200 {string} 删除成功
// @router /delete [delete]
func (c *TagController) Delete() {
	name := c.Input().Get("name")
	url := "https://kube.gwunion.cn/api/repositories/venus/nginx/tags/" + name
	req := httplib.Delete(url)
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Debug(true)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	res := map[string]interface{}{"result": str, "code": 20000}
	c.Data["json"] = res
	c.ServeJSON()
	return
}

//@Title allLabels
//@Description select all labels for Tag
//@Param project_id query int true "project's id"
//@Success 200 {string} 查询成功
// @router /findLabels [get]
func (c *TagController) FindLabels() {
	url := "https://kube.gwunion.cn/api/labels?scope=p&project_id=3"
	req := httplib.Get(url)
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Debug(true)
	result := []map[string]interface{}{}
	err := req.ToJSON(&result)
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	json := map[string]interface{}{"result": result, "code": 20000}
	c.Data["json"] = json
	c.ServeJSON()
	return
}

//@Title removeLabels
//@Description remove one label from your selected tag
//@Param project_name query string true "project's name"
//@Param repo_name query string true "repository's name"
//@Param tag_name query string true "tag's name"
//@Param label_id query int true "label's id"
//@Success 200 {string} 删除成功
// @router /removeLabels [delete]
func (c *TagController) RemoveLabels() {
	label_id := c.GetString("label_id")
	name := c.GetString("name")
	fmt.Println("-----", label_id, name, "-----")
	url := "https://kube.gwunion.cn/api/repositories/venus/nginx/tags/" + name + "/labels/" + label_id
	req := httplib.Delete(url)
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Debug(true)
	resp, _ := req.Response()
	fmt.Println("------------------\n", resp)
	res := map[string]interface{}{"code": 20000}
	c.Data["json"] = res
	c.ServeJSON()
	return
}

//@Title addLabels
//@Description add one label from your selected tag
//@Param project_name query string true "project's name"
//@Param repo_name query string true "repository's name"
//@Param tag_name query string true "tag's name"
//@Success 200 {string} 添加成功
// @router /addLabels [post]
func (c *TagController) AddLabels() {
	name := c.Input().Get("name")
	label_id, _ := strconv.Atoi(c.Input().Get("label_id"))
	fmt.Println("-----", label_id, name, "-----")
	url := "https://kube.gwunion.cn/api/repositories/venus/nginx/tags/" + name + "/labels/"
	req := httplib.Post(url)
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	js := map[string]int{"id": label_id}
	req.JSONBody(js)
	req.Debug(true)
	resp, _ := req.Response()
	fmt.Println(resp)
	res := map[string]interface{}{"code": 20000}
	c.Data["json"] = res
	c.ServeJSON()
	return
}

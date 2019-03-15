package controllers

import (
	"crypto/tls"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"strconv"
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
// @router / [get]
func (c *TagController) Get() {
	repoName := c.Input().Get("repoName")
	url := "https://kube.gwunion.cn/api/repositories/" + repoName + "/tags?detail=1"
	req := httplib.Get(url)
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Debug(true)
	rep, _ := req.Response()
	//fmt.Println(rep)
	result := []map[string]interface{}{}
	err := req.ToJSON(&result)
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	if rep.StatusCode == 200 {
		json := map[string]interface{}{"result": result, "code": 20000}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]int{"code": rep.StatusCode}
		c.ServeJSON()
	}
	return
}

//@Title deleteTag
//@Description delete tag for selected
//@Param project_name query string true "project's name"
//@Param repo_name query string true "repository's name"
//@Param tag_name query string true "tag's name"
//@Success 200 {string} 删除成功
// @router / [delete]
func (c *TagController) Delete() {
	repoName := c.Input().Get("repoName")
	name := c.Input().Get("name")
	url := "https://kube.gwunion.cn/api/repositories/" + repoName + "/tags/" + name
	fmt.Println(url)
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
// @router /label [get]
func (c *TagController) FindLabels() {
	projectId := c.Input().Get("projectId")
	url := "https://kube.gwunion.cn/api/labels?scope=p&project_id=" + projectId
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
// @router /label [delete]
func (c *TagController) RemoveLabels() {
	repoName := c.Input().Get("repoName")
	labelId := c.GetString("labelId")
	tagName := c.GetString("tagName")
	fmt.Println("-----", labelId, tagName,repoName, "-----")
	url := "https://kube.gwunion.cn/api/repositories/" + repoName + "/tags/" + tagName + "/labels/" + labelId
	fmt.Println(url)
	req := httplib.Delete(url)
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Debug(true)
	rep, _ := req.Response()
	fmt.Println("------------------\n", rep)
	if rep.StatusCode == 200 {
		res := map[string]interface{}{"code": 20000}
		c.Data["json"] = res
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]int{"code": rep.StatusCode}
		c.ServeJSON()
	}
	return
}

//@Title addLabels
//@Description add one label from your selected tag
//@Param project_name query string true "project's name"
//@Param repo_name query string true "repository's name"
//@Param tag_name query string true "tag's name"
//@Success 200 {string} 添加成功
// @router /label [post]
func (c *TagController) AddLabels() {
	repoName := c.Input().Get("repoName")
	tagName := c.Input().Get("tagName")
	label_id, _ := strconv.Atoi(c.Input().Get("labelId"))
	fmt.Println("-----", label_id, tagName, repoName, "-----")
	url := "https://kube.gwunion.cn/api/repositories/" + repoName + "/tags/" + tagName + "/labels/"
	fmt.Println(url)
	req := httplib.Post(url)
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	js := map[string]int{"id": label_id}
	req.JSONBody(js)
	req.Debug(true)
	rep, _ := req.Response()
	fmt.Println(rep)
	if rep.StatusCode == 200 {
		res := map[string]interface{}{"code": 20000}
		c.Data["json"] = res
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]int{"code": rep.StatusCode}
		c.ServeJSON()
	}
	return
}

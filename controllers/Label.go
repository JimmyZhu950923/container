package controllers

import (
	"crypto/tls"
	"strconv"

	"k8s.io/apimachinery/pkg/util/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

// Operations about object
type LabelController struct {
	beego.Controller
}

// @Title Get all
// @Description  Get All Label
// @Param name query string false "label name"
// @Param scope query string true "p"
// @Param project_id query int true "project id"
// @Success 200 {string} 查询成功
// @router / [get]
func (o *LabelController) FindLabels() {
	name := o.Input().Get("name")
	projectId := o.Input().Get("project_id")
	url := "https://kube.gwunion.cn/api/labels"
	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	req.Debug(true)
	req.Param("scope", "p")
	req.Param("project_id", projectId)
	req.Param("name", name)
	response, _ := req.Response()
	//fmt.Println(response)
	if response.StatusCode == 200 {
		total, _ := strconv.Atoi(response.Header.Get("X-Total-Count"))
		json := []map[string]interface{}{}
		err := req.ToJSON(&json)
		if err != nil {
			//fmt.Println(err)
			//fmt.Println(json)
		}
		result := map[string]interface{}{"total": total, "json": json, "code": 20000}
		o.Data["json"] = result

	} else {
		json := map[string]int{"code": response.StatusCode}
		o.Data["json"] = json
	}
	o.ServeJSON()
	return
}

// @Title Get a label
// @Description   Get label by id
// @Param id query int true "label id"
// @Success 200 {string} 查询成功
//@router /:id [get]
func (o *LabelController) FindLabelsById() {
	id := o.Ctx.Input.Param(":id")
	url := "https://kube.gwunion.cn/api/labels/" + id
	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	req.Debug(true)
	response, _ := req.Response()
	if response.StatusCode == 200 {
		var result map[string]interface{}
		err := req.ToJSON(&result)
		if err != nil {
			//fmt.Println(err)
		}
		json := map[string]interface{}{"result": result, "code": 20000}
		o.Data["json"] = json

	} else {
		json := map[string]int{"code": response.StatusCode}
		o.Data["json"] = json
	}
	o.ServeJSON()
	return
}

//@Title Add Label
//@Description Add a label
//@Param name query string true "label name"
//@Param color query string true "label color"
//@Param description query string false "label description"
//@Success 200 {string} 添加成功
//@router / [post]
func (o *LabelController) AddLabel() {
	name := getRequestBody(o)["name"]
	color := getRequestBody(o)["color"]
	description := getRequestBody(o)["description"]
	url := "https://kube.gwunion.cn/api/labels"
	req := httplib.Post(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	label := map[string]interface{}{"name": name, "description": description, "color": color, "scope": "p", "project_id": 9}
	_, _ = req.JSONBody(label)
	response, _ := req.Response()
	if response.StatusCode == 201 {
		o.Data["json"] = map[string]interface{}{"code": 20000}

	} else {
		json := map[string]int{"code": response.StatusCode}
		o.Data["json"] = json
	}
	o.ServeJSON()
	return

}

//@Title Update label
//@Description update label
//@Param name query string false "label name"
//@Param color query string false "label color"
//@Param description query string false "label description"
//@Success 200 {string} 修改成功
//@router /:id [put]
func (o *LabelController) UpdateLabel() {
	id := o.Ctx.Input.Param(":id")
	name := getRequestBody(o)["name"]
	color := getRequestBody(o)["color"]
	description := getRequestBody(o)["description"]
	url := "https://kube.gwunion.cn/api/labels/" + id
	req := httplib.Put(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	label := map[string]interface{}{"name": name, "description": description, "color": color, "scope": "p", "project_id": 3}
	_, _ = req.JSONBody(label)
	response, _ := req.Response()
	if response.StatusCode == 200 {
		o.Data["json"] = map[string]interface{}{"code": 20000}
	} else {
		json := map[string]int{"code": response.StatusCode}
		o.Data["json"] = json
	}
	o.ServeJSON()
	return

}

//@Title Delete Lable
//@Description Delete label
//@Param id query int true "label_id"
//@Success 200 {string} 删除成功
//@router  /:id  [delete]
func (o *LabelController) DeleteLabel() {
	id := o.Ctx.Input.Param(":id")
	url := "https://kube.gwunion.cn/api/labels/" + id
	req := httplib.Delete(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	req.Debug(true)
	response, _ := req.Response()
	if response.StatusCode == 200 {
		o.Data["json"] = map[string]interface{}{"code": 20000}
	} else {
		json := map[string]int{"code": response.StatusCode}
		o.Data["json"] = json
	}
	o.ServeJSON()
	return

}

func getRequestBody(o *LabelController) (param map[string]interface{}) {
	json.Unmarshal(o.Ctx.Input.RequestBody, &param)
	//fmt.Println(param)
	return
}

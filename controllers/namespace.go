package controllers

import (
	"gt-container-go/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Operations about Users
type NamespaceController struct {
	beego.Controller
}

//var clientset = getClientset()

// @Title Get All Namespace
// @Description get namespace
// @Success 200 {object} models.User
// @router / [get]
func (n *NamespaceController) GetNamespace() {
	// // clientset := getClientset()
	// namespace, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// var json = map[string]interface{}{"code": 20000, "data": namespace}
	// n.Data["json"] = json
	// n.ServeJSON()
	// return
	var namespaces []models.Namespace
	row, _ := models.GetAllNsp(&namespaces)
	n.Data["json"] = map[string]interface{}{"code": 20000, "data": namespaces, "row": row}
	n.ServeJSON()
	return
}

// @Title Get a Namespace
// @Description get  namespace by name
// @Param name path string true "Namespace name"
// @Success 200 {object} models.User
//@router /:name [get]
func (n *NamespaceController) GetSingle() {
	name := n.Ctx.Input.Param(":name")

	// namespace, err := clientset.CoreV1().Namespaces().Get(name, metav1.GetOptions{})
	// if err != nil {
	// 	n.Data["json"] = map[string]interface{}{"code": 400, "message": name + "不存在"}
	// 	n.ServeJSON()
	// } else {
	// 	var json = map[string]interface{}{"code": 20000, "data": namespace}
	// 	n.Data["json"] = json
	// 	n.ServeJSON()
	// 	return
	// }
	var namespaces []models.Namespace
	row, _ := models.GetSingleNsp(name, &namespaces)
	n.Data["json"] = map[string]interface{}{"code": 20000, "data": namespaces, "row": row}
	n.ServeJSON()
	return
}

// @Title Add a Namespace
// @Description set name insert namespace
// @Param name query string true "Namespace name"
// @Success 200 {object} models.User
//@router / [post]
func (n *NamespaceController) Add() {
	name := n.Input().Get("name")
	// 先查询数据库是否存在输入的 name
	err := models.FindNspByName(name)
	// 数据库的 name 为空
	if err == orm.ErrNoRows {
		//插入数据到数据库
		var nsp = models.Namespace{Name: name, Status: "Active", CreateTime: time.Now()}
		_, err := models.InsertNsp(&nsp)
		//数据库事务
		if err != nil {
			//出错回滚
			models.Db.Rollback()
			n.Data["json"] = map[string]interface{}{"code": 405, "message": "数据库未知错误！"}
			n.ServeJSON()
		} else {
			//数据库成功执行 k8s插入namespace
			var namespace v1.Namespace
			namespace.SetName(name)
			ns, err := clientset.CoreV1().Namespaces().Create(&namespace)
			//k8s 出错
			if err != nil {
				//数据库回滚
				models.Db.Rollback()
				n.Data["json"] = map[string]interface{}{"code": 409, "message": name + "已存在!"}
				n.ServeJSON()
			} else {
				//k8s 插入成功，数据库提交事件
				models.Db.Commit()
				var json = map[string]interface{}{"code": 20000, "data": ns}
				n.Data["json"] = json
				n.ServeJSON()
			}
		}
	} else {
		//数据库中存在，返回错误信息
		n.Data["json"] = map[string]interface{}{"code": 409, "message": name + "已存在!"}
		n.ServeJSON()
	}

}

// @Title Delete  Namespace
// @Description delete namespace by name
// @Param name  path string true "Namespace name"
// @Success 200 {object} models.User
//@router /:name [delete]
func (n *NamespaceController) Delete() {
	name := n.Ctx.Input.Param(":name")
	_, err := models.DeleteNsp(name)
	//判断数据库
	if err != nil {
		models.Db.Rollback()
		n.Data["json"] = map[string]interface{}{"code": 405, "message": "数据库未知错误！"}
		n.ServeJSON()
	} else {
		err := clientset.CoreV1().Namespaces().Delete(name, &metav1.DeleteOptions{})
		//判断k8s
		if err != nil {
			panic(err)
			models.Db.Rollback()
			n.Data["json"] = map[string]interface{}{"code": 405, "message": "数据库未知错误！"}
			n.ServeJSON()
		} else {
			models.Db.Commit()
			n.Data["json"] = map[string]interface{}{"code": 20000}
			n.ServeJSON()
		}
	}
}

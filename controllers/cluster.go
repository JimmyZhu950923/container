package controllers

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/astaxie/beego"
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
	// clientset := getClientset()
	namespace, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	var json = map[string]interface{}{"code": 20000, "data": namespace}
	n.Data["json"] = json
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
	fmt.Println(name)
	namespace, err := clientset.CoreV1().Namespaces().Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		n.Data["json"] = map[string]interface{}{"code": 20000}
		n.ServeJSON()
	} else if err != nil {
		n.Data["json"] = map[string]interface{}{"code": 400, "message": name + "不存在"}
		n.ServeJSON()
	} else {
		var json = map[string]interface{}{"code": 20000, "data": namespace}
		n.Data["json"] = json
		n.ServeJSON()
		return
	}
}

// @Title Add a Namespace
// @Description set name insert namespace
// @Param name query string true "Namespace name"
// @Success 200 {object} models.User
//@router / [post]
func (n *NamespaceController) Add() {
	var namespace v1.Namespace
	name := n.Input().Get("name")
	namespace.SetName(name)
	ns, err := clientset.CoreV1().Namespaces().Create(&namespace)
	if err != nil {
		n.Data["json"] = map[string]interface{}{"code": 409, "message": name + "已存在"}
		n.ServeJSON()
	} else {
		var json = map[string]interface{}{"code": 20000, "data": ns}
		n.Data["json"] = json
		n.ServeJSON()
		return
	}
}

// @Title Delete  Namespace
// @Description delete namespace by name
// @Param name  path string true "Namespace name"
// @Success 200 {object} models.User
//@router /:name [delete]
func (n *NamespaceController) Delete() {
	name := n.Ctx.Input.Param(":name")
	err := clientset.CoreV1().Namespaces().Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	n.Data["json"] = map[string]int{"code": 20000}
	n.ServeJSON()
	return
}

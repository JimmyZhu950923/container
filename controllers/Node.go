package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Operations about Users
type NodeController struct {
	beego.Controller
}

// 返回的pod中cpu，memory数据,pod数量
type podData struct {
	CpuRequests    int //cpu请求值
	CpuLimits      int //cpu限制值
	MemoryRequests int //内存请求值
	MemoryLimits   int //内存限制值
	podNum         int //pod数量
}

//var clientset = getClientset()

// @Title GetAll
// @Description get all Nodes
// @Success 200 {object} models.User
// @router / [get]
func (n *NodeController) GetAll() {
	// clientset := getClientset()
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(len(nodes.Items))
	n.Data["json"] = map[string]interface{}{"code": 20000, "data": nodes}
	n.ServeJSON()
	return
}

//@Title Get node by name
//@Description get Node by Name
//@Param name query string false "node's name"
//@Success 200 {object} models.User
//@router /:name [get]
func (n *NodeController) GetByName() {
	name := n.Ctx.Input.Param(":name")
	node, err := clientset.CoreV1().Nodes().Get(name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	pod, err := GetPodByNodeName(name)
	if err != nil {
		panic(err)
	}
	for _, v := range pod.Items {
		req := v.Spec.Containers[0].Resources.Requests
		lim := v.Spec.Containers[0].Resources.Limits
		if req != nil {
			a := req["cpu"]
			cpuR := a.String()[0 : len(a.String())-1]
			fmt.Println(cpuR)
		}
		if lim != nil {

		}

		fmt.Println("cpu>>>:", req["cpu"], "memory>>>:", req["memory"])
		fmt.Println("Lcpu:", lim["cpu"], "Lmemory:", lim["memory"])
	}

	n.Data["json"] = map[string]interface{}{"data": node, "pod": pod, "code": 20000}
	n.ServeJSON()
	return
}

////@router /delete [get]
//func (n *NodesController) Delete(){
//	err:=clientset.CoreV1().Nodes().Delete("test",&metav1.DeleteOptions{})
//	if err!=nil{
//		panic(err)
//	}
//}

package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Operations about Users
type NodeController struct {
	beego.Controller
}

// 返回的pod中cpu，memory数据,pod数量
type allocatedResources struct {
	CpuRequests    int //cpu请求值
	CpuLimits      int //cpu限制值
	MemoryRequests int //内存请求值
	MemoryLimits   int //内存限制值
	PodNum         int //pod数量
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
	// 对应的cpu请求值  内存请求值，cpu限制值 内存限制值
	var cpuR, memoryR, cpuL, memoryL = 0, 0, 0, 0
	for _, v := range pod.Items {
		req := v.Spec.Containers[0].Resources.Requests
		lim := v.Spec.Containers[0].Resources.Limits
		if req != nil {
			//cpuReqVal : request map集合下的所有key为cpu对应的val
			cpuReqVal := req["cpu"]
			//cpuR, _ = strconv.Atoi(a.String()[0 : len(a.String())-1])
			cpu := cpuReqVal.String()[0 : len(cpuReqVal.String())-1]
			//memoryReqVal : request map集合下的所有key为memory对应的val
			memoryReqVal := req["memory"]
			memory := memoryReqVal.String()[0:len(memoryReqVal.String())]
			if memory != "0" {
				memory = memory[0 : len(memory)-2]
			}
			//字符串转数字
			Rcpu, _ := strconv.Atoi(cpu)
			Rmemory, _ := strconv.Atoi(memory)
			cpuR += Rcpu
			memoryR += Rmemory
		}
		if lim != nil {
			cpuLimVal := lim["cpu"]
			//cpuR, _ = strconv.Atoi(a.String()[0 : len(a.String())-1])
			cpu := cpuLimVal.String()[0 : len(cpuLimVal.String())-1]
			memoryLimVal := lim["memory"]
			memory := memoryLimVal.String()[0:len(memoryLimVal.String())]
			if memory != "0" {
				memory = memory[0 : len(memory)-2]
			}
			//字符串转数字
			Lcpu, _ := strconv.Atoi(cpu)
			Lmemory, _ := strconv.Atoi(memory)
			cpuL += Lcpu
			memoryL += Lmemory
		}

	}
	length := len(pod.Items)
	allocatedResources := allocatedResources{CpuRequests: cpuR, CpuLimits: cpuL, MemoryRequests: memoryR, MemoryLimits: memoryL, PodNum: length}
	n.Data["json"] = map[string]interface{}{"data": node, "pod": pod, "allocatedResources": allocatedResources, "code": 20000}
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

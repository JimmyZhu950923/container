package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Operations about Users
type PodsController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Pods
// @Success 200 {object} models.User
// @router /getA		ll [get]
func (p *PodsController) GetAll() {
	//clientset := getClientset()
	//clientset := getInClusterClientset()
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	p.Data["json"] = pods
	p.ServeJSON()
	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message

}

// @Title Get
// @Description get Pod by namespace
// @Success 200 {object} models.User
// @router /singlePod [get]
func (p *PodsController) GetSingle() {
	//clientset := getClientset()

	namespace := p.Input().Get("namespace")
	podName := p.Input().Get("podName")
	pod, err := clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})

	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
			pod, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		//fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		p.Data["json"] = map[string]interface{}{"code": 20000, "data": pod}
		p.ServeJSON()
	}
}

// @Title GetPodsInNameSpace
// @Description get Pods in Namespace default
// @Param nodeName query string true "节点名称"
// @Success 200 {object} models.User
// @router /list [get]
func (p *PodsController) GetPodsInNameSpace() {

	//clientset := getClientset()
	nameSpace := p.Input().Get("nameSpace")
	deploymentName := p.Input().Get("deploymentName")
	nodeName := p.Input().Get("nodeName")
	//fmt.Println(nameSpace)

	var listOptions metav1.ListOptions
	if nodeName != "" {
		listOptions.FieldSelector = "spec.nodeName=" + nodeName
	}

	pods, err := clientset.CoreV1().Pods(nameSpace).List(listOptions)
	var podItems []v1.Pod
	if deploymentName != "" {
		rs, err := clientset.AppsV1().ReplicaSets(nameSpace).List(listOptions)
		if err != nil {
			panic(err.Error())
		} else {
			for _, item1 := range rs.Items {
				if item1.OwnerReferences[0].Name == deploymentName {

					for _, item2 := range pods.Items {
						if item2.OwnerReferences[0].Name == item1.Name {
							podItems = append(podItems, item2)
						}
					}

					break
				}
			}
		}
		pods.Items = podItems
	}

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("There are ", len(pods.Items), " pods in", nameSpace)

		p.Data["json"] = map[string]interface{}{"code": 20000, "data": pods}
		p.ServeJSON()
	}

}

// @Title newPod
// @Description create new pod
// @Success 200 {object} model.user
// @router /newPod [get]
func (p *PodsController) NewPod() {
	var pod v1.Pod
	pod.SetGenerateName("wt-test-")
	pod.Spec.NodeName = "node2"
	pod.Spec.Containers = []v1.Container{{Name: "web", Image: "kube.gwunion.cn/venus/nginx:alpine"}}
	pod1, err := clientset.CoreV1().Pods("default").Create(&pod)
	if err != nil {
		panic(err.Error())
	}
	p.Data["json"] = pod1
	p.ServeJSON()
}

// @Title delete pod
// @Description delete a pod
// @Param name query true "name of the pod"
// @router / [delete]
func (p *PodsController) DeletedPod() {

	name := p.Input().Get("name")
	namespace := p.Input().Get("namespace")
	err := clientset.CoreV1().Pods(namespace).Delete(name, &metav1.DeleteOptions{})

	if err != nil {
		panic(err.Error())
		p.Data["json"] = map[string]int{"message": 111111111}
		p.ServeJSON()
	}
	p.Data["json"] = map[string]int{"code": 20000}
	p.ServeJSON()
}

// @Title UpdatedPod
// @Description update pod
// @Param name query string true "pod's name"
// @Param nameSpace query string true "命名空间"
// @Success 200 {string} 修改成功
// @router / [put]
func (p *PodsController) UpdatedPod() {
	name := p.Input().Get("name")
	nameSpace := p.Input().Get("nameSpace")
	pod, err := clientset.CoreV1().Pods(nameSpace).Get(name, metav1.GetOptions{})
	_, err = clientset.CoreV1().Pods(nameSpace).Update(pod)

	if err != nil {
		p.Data["json"] = map[string]interface{}{"code": 400, "data": err.Error()}
		p.ServeJSON()
	} else {
		p.Data["json"] = map[string]interface{}{"code": 20000}
		p.ServeJSON()
	}

}

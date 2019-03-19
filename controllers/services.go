package controllers

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

// Operations about Users
type ServicesController struct {
	beego.Controller
}


var clientset = getClientset()

// @Title GetAll
// @Description get all Pods
// @Success 200 {object} models.User
// @router / [get]
func (s *ServicesController) GetAll() {
	//clientset := getInClusterClientset()
	services, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("There are %d services in the cluster\n", len(services.Items))
	json := map[string]interface{}{"data": services, "code": 20000}
	s.Data["json"] = json
	s.ServeJSON()
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func getClientset() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func getInClusterClientset() *kubernetes.Clientset {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

// @Title newService
// @Description create new service
// @Success 200 {string} 添加成功
// @router / [post]
func (s *ServicesController)NewS() {
	var service v1.Service
	name := s.Input().Get("name")
	service.SetName(name)
	fmt.Println("name = ", name)
	service.APIVersion = "v1"
	service.Kind = "Service"
	service.Spec = v1.ServiceSpec{
		Ports: []v1.ServicePort{v1.ServicePort{Port: 80}},
	}
	service1, err := clientset.CoreV1().Services("default").Create(&service)
	if err != nil {
		panic(err.Error())
	}
	json := map[string]interface{}{"data": service1, "code": 20000}
	s.Data["json"] = json
	s.ServeJSON()
	return
	}

	//@Title delService
	//@Description delete one service
	//@Success 200 {string} 删除成功
	//@router / [delete]
	func (s *ServicesController)DelS() {
		name := s.Input().Get("name")
		err := clientset.CoreV1().Services("default").Delete(name, &metav1.DeleteOptions{})
		if err != nil {
			panic(err.Error())
		}
		s.Data["json"] = map[string]interface{}{"code": 20000}
		s.ServeJSON()
	}

	//// @Title updService
	//// @Description updeate service
	//// @Success 200 {string} 更新成功
	//// @router /updS [get]
	//func (s *ServicesController)UpdS(){
	//	service, err := clientset.CoreV1().Services("default").Get("example-service", metav1.GetOptions{})
	//	service.SetName("")
	//	service1, err := clientset.CoreV1().Services("default").Update(service)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	s.Data["json"] = service1
	//	s.ServeJSON()
	//}
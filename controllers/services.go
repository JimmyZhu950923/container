package controllers

import (
"flag"
"fmt"
	"k8s.io/api/core/v1"
	"os"
"path/filepath"

"github.com/astaxie/beego"
"k8s.io/apimachinery/pkg/api/errors"
metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
"k8s.io/client-go/kubernetes"
"k8s.io/client-go/rest"
"k8s.io/client-go/tools/clientcmd"
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
	clientset := getInClusterClientset()
	services, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(services.Items))
	s.Data["json"] = services
	s.ServeJSON()
	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	namespace := "default"
	service := "example-service"
	_, err = clientset.CoreV1().Services(namespace).Get(service, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Service %s in namespace %s not found\n", service, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting service %s in namespace %s: %v\n",
			service, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found service %s in namespace %s\n", service, namespace)
	}
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
// @router /newS [get]
func (s *ServicesController)NewS(){
	var service v1.Service
	service.SetName("test")
	service.APIVersion = "v1"
	service.Kind = "Service"
	service.Spec = v1.ServiceSpec{
		Ports:[]v1.ServicePort{v1.ServicePort{Port:80}},
	}
	service1,err := clientset.CoreV1().Services("default").Create(&service)
	if err != nil {
		panic(err.Error())
	}

	s.Data["json"] = service1
	s.ServeJSON()
}

// @Title delService
// @Description delete one service
// @Success 200 {string} 删除成功
// @router /delS [get]
func (s *ServicesController)DelS() {
	err := clientset.CoreV1().Services("default").Delete("test", &metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}
	s.Data["json"] = map[string]string{"data": "成功"}
	s.ServeJSON()
}

// @Title getService
// @Description get one service
// @Success 200 {string} 查找一个成功
// @router /getS [get]
func (s *ServicesController)GetS() {
	namespace := "default"
	service := "example-service"
	_, err := clientset.CoreV1().Services(namespace).Get(service, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Service %s in namespace %s not found\n", service, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting service %s in namespace %s: %v\n",
			service, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found serivce %s in namespace %s\n", service, namespace)
	}
}

// @Title updService
// @Description updeate service
// @Success 200 {string} 更新成功
// @router /updS [get]
func (s *ServicesController)UpdS(){
	service, err := clientset.CoreV1().Services("default").Get("example-service", metav1.GetOptions{})
	service.SetName("")
	service1, err := clientset.CoreV1().Services("default").Update(service)
	if err != nil {
		panic(err.Error())
	}
	s.Data["json"] = service1
	s.ServeJSON()
}
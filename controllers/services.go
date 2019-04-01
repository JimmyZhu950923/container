package controllers

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"strconv"
)

// Operations about Users
type ServicesController struct {
	beego.Controller
}


var clientset = getClientset()

// @Title getServices
// @Description get all Services,
// @Param namespace query string false "namespace for service"
// @Success 200 {object} models.User
// @router / [get]
func (s *ServicesController) GetServices() {
	//clientset := getInClusterClientset()
	namespace := s.Input().Get("namespace")
	fmt.Println("namespace = ", namespace)
	services, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("There are %d services in the cluster\n", len(services.Items))
	json := map[string]interface{}{"data": services, "code": 20000}
	s.Data["json"] = json
	s.ServeJSON()
}

// @Title getSingle
// @Description get one Service,
// @Param namespace path string false "namespace for service"
// @Param name query string false "name for service"
// @Success 200 {object} models.User
// @router /:namespace [get]
func (s *ServicesController) GetSingleService() {
	namespace := s.Ctx.Input.Param(":namespace")
	name := s.Input().Get("name")
	//fmt.Println("namespace = ", namespace)
	//fmt.Println("name = ", name)
	service, err :=clientset.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Service %s in namespace %s not found\n", service, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting service %s in namespace %s: %v\n",
			service, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		//fmt.Printf("Found service %s in namespace %s\n", service, namespace)
		s.Data["json"] = map[string]interface{}{"code": 20000, "data": service}
		s.ServeJSON()
	}
}

// @Title getService
// @Description get service by namespace
// @Param namespace path string false "namespace for service"
// @Param name query string false "name for service"
// @Success 200 {object} models.User
// @router /:name [get]
func (s *ServicesController) GetService() {
	//clientset := getClientset()

	namespace := s.Input().Get("namespace")
	name := s.Ctx.Input.Param(":name")
	service, err := clientset.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})

	if errors.IsNotFound(err) {
		fmt.Printf("Service %s in namespace %s not found\n", service, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting service %s in namespace %s: %v\n",
			service, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		//fmt.Printf("Found service %s in namespace %s\n", service, namespace)
		s.Data["json"] = map[string]interface{}{"code": 20000, "data": service}
		s.ServeJSON()
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

// @Title createService
// @Description create new service
// @Param name body string false "name for service"
// @Param port body string false "port for service"
// @Param namespace body string false "namespace for service"
// @Param type body string false "type for service"
// @Success 200 {string} 添加成功
// @router / [post]
func (s *ServicesController)CreateService() {
	var service v1.Service
	name := s.Input().Get("name")
	port,err := strconv.Atoi(s.Input().Get("port"))
	namespace := s.Input().Get("namespace")
	kind := s.Input().Get("type")
	service.SetName(name)
	//fmt.Println("name = ", name)
	//fmt.Println("namespace = ", namespace)
	//fmt.Println("port = ", port)
	//fmt.Println("type = ", kind)
	service.APIVersion = "v1"
	service.Kind = "Service"
	service.Spec = v1.ServiceSpec{
		Ports: []v1.ServicePort{v1.ServicePort{Port: int32(port)}},
		Type: v1.ServiceType(kind),
	}
	service1, err := clientset.CoreV1().Services(namespace).Create(&service)
	if err != nil {
		panic(err.Error())
	}
	json := map[string]interface{}{"data": service1, "code": 20000}
	s.Data["json"] = json
	s.ServeJSON()
	return
	}

// @Title delService
// @Description delete one service
// @Param name query string false "name for service"
// @Param namespace query string false "namespace for service"
// @Success 200 {string} 删除成功
// @router / [delete]
func (s *ServicesController)DelService() {
	name := s.Input().Get("name")
	namespace := s.Input().Get("namespace")
	//fmt.Println("name = ", name)
	//fmt.Println("namespace = ", namespace)
	err := clientset.CoreV1().Services(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}
	s.Data["json"] = map[string]interface{}{"code": 20000}
	s.ServeJSON()
}

// @Title updService
// @Description updeate service
// @Success 200 {string} 更新成功
// @router / [put]
func (s *ServicesController)UpdS(){
	namespace := s.Ctx.Input.Param(":namespace")
	name := s.Input().Get("name")
	service, err := clientset.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	_, err = clientset.CoreV1().Services(namespace).Update(service)
	if err != nil {
		s.Data["json"] = map[string]interface{}{"code": 400, "data": err.Error()}
		s.ServeJSON()
	} else {
		s.Data["json"] = map[string]interface{}{"code": 20000}
		s.ServeJSON()
	}
}
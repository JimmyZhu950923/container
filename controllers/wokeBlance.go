package controllers

import (
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Operations about WorkBlance
type WorkBlanceController struct {
	beego.Controller
}

// @Title Singel Selete
// @Description get Single Deployment By name
// @Param name path string false "name of the deployment"
// @Param namespace query string false "namespace of the deployment"
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router /:name [get]
func (w *WorkBlanceController) GetSingleDeployment() {

	name := w.Ctx.Input.Param(":name")
	namespace := w.Input().Get("namespace")
	deploymentsClient := clientset.AppsV1().Deployments(namespace)

	deployment, err := deploymentsClient.Get(name, metav1.GetOptions{})

	if errors.IsNotFound(err) {
		w.Data["json"] = map[string]interface{}{"code": 20000}
		w.ServeJSON()
	} else if err != nil {
		panic(err)
	} else {
		w.Data["json"] = map[string]interface{}{"code": 20000, "data": deployment}
		w.ServeJSON()
	}

}

// GetAll ...
// @Title Get All Deployments
// @Description get Deployments
// @Param namespace query string false "Namespace of the deployment"
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router / [get]
func (w *WorkBlanceController) GetDeployments() {

	//clientset := getClientset()
	deploymentsClient := clientset.AppsV1().Deployments(w.Input().Get("namespace"))

	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	} else {
		w.Data["json"] = map[string]interface{}{"code": 20000, "data": list}
		w.ServeJSON()
	}

}

// Create deployment ...
// @Title Create deployment
// @Description get Userinfo
// @Param name query string "Name of the deployment"
// @Param num query string "replicas of the deployment"
// @Param image query string "image of the deployment"
// @Param namespace query string "namespace of the deployment"
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router / [post]
func (w *WorkBlanceController) CreateDeployment() {
	//clientset := getClientset()

	name := w.Input().Get("name")
	num, err := strconv.ParseInt(w.Input().Get("num"), 10, 32)
	image := w.Input().Get("image")
	namespace := w.Input().Get("namespace")

	//fmt.Println(">>>>>>>>>>>>>")
	//fmt.Println(name, num, image, namespace)
	//fmt.Println(">>>>>>>>>>>>>")

	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	var deployment = &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(int32(num)),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "kube.gwunion.cn/" + image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	// Create Deployment
	//fmt.Println("Creating deployment...")
	_, err = deploymentsClient.Create(deployment)

	if err != nil {
		k8Err := err.(*errors.StatusError)
		//fmt.Println(k8Err.ErrStatus.Code)
		w.Data["json"] = map[string]interface{}{"code": k8Err.ErrStatus.Code, "message": "负载" + name + "已存在"}
		w.ServeJSON()
	} else {
		//fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
		w.Data["json"] = map[string]int{"code": 20000}
		w.ServeJSON()
	}
}

// Create deployment ...
// @Title Create deployment
// @Description get Userinfo
// @Param deployment body string "deploment of the yaml or josn"
// @Param namespace query string "namespace of the deployment"
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router /:namespace [post]
func (w *WorkBlanceController) CreateDeploymentByYaml() {

	deployment := w.Ctx.Input.RequestBody
	namespace := w.Ctx.Input.Param(":namespace")
	var err error
	if strings.Index(string(deployment), "{") == -1 {
		deployment, err = yaml.ToJSON(deployment)
		if err != nil {
			panic(err.Error())
		}

	}
	var deploymentEntity appsv1.Deployment
	//fmt.Println(string(deployment))
	err = json.Unmarshal(deployment, &deploymentEntity)
	if err != nil {
		panic(err.Error())
	}

	_, err = clientset.AppsV1().Deployments(namespace).Create(&deploymentEntity)
	if err != nil {
		panic(err.Error())
	} else {
		//fmt.Println("创建陈工")
		w.Data["json"] = map[string]int{"code": 20000}
		w.ServeJSON()
	}

}

// Update deployment
// @Title update deployment
// @Description get Userinfo
// @Param query query string false "Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router / [put]
func (w *WorkBlanceController) UpdateDeployment() {
	//clientset := getClientset()

	namespace := w.Input().Get("namespace")

	deploymentsClient := clientset.AppsV1().Deployments(namespace)

	name := w.Input().Get("name")
	num := w.Input().Get("num")

	if num != "" {
		num, err := strconv.ParseInt(w.Input().Get("num"), 10, 32)
		if err != nil {
			panic(err.Error())
		}
		result, getErr := deploymentsClient.Get(name, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v ", getErr))
		}

		result.Spec.Replicas = int32Ptr(int32(num)) // reduce replica count

		_, updateErr := deploymentsClient.Update(result)
		if updateErr != nil {
			panic(updateErr.Error())
		} else {
			//fmt.Println("Updated deployment...")
			w.Data["json"] = map[string]int{"code": 20000}
			w.ServeJSON()
		}
	} else {

		var deployment1 appsv1.Deployment
		err1 := json.Unmarshal([]byte(name), &deployment1)
		if err1 != nil {
			panic(err1.Error())
		}

		_, err := deploymentsClient.Update(&deployment1)
		if err != nil {
			panic(err.Error())
		} else {
			//fmt.Println("Updated deployment...")
			w.Data["json"] = map[string]int{"code": 20000}
			w.ServeJSON()
		}

	}

	//fmt.Println("Updated deployment...")
	w.Data["json"] = map[string]int{"code": 20000}
	w.ServeJSON()
}

// Delete deployment by name
// @Title Get All
// @Description get Userinfo
// @Param name query string true "Name of deployment"
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router / [delete]
func (w *WorkBlanceController) DeleteDeployment() {
	// Delete Deployment
	//clientset := getClientset()

	name := w.Input().Get("name")
	namespace := w.Input().Get("namespace")
	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	//fmt.Println("Deleted deployment.")
	w.Data["json"] = map[string]int{"code": 20000}
	w.ServeJSON()
}

// Create daemonset ...
// @Title Create daemonset
// @Description get Userinfo
// @Param name query string "Name of the daemonset"
// @Param num query string "replicas of the daemonset"
// @Param image query string "image of the daemonset"
// @Param namespace query string "namespace of the daemonset"
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router /daemonset [post]
func (w *WorkBlanceController) CreateDaemonset() {
	//clientset := getClientset()

	name := w.Input().Get("name")
	image := w.Input().Get("image")
	namespace := w.Input().Get("namespace")

	//fmt.Println(">>>>>>>>>>>>>")
	//fmt.Println(name, image, namespace)
	//fmt.Println(">>>>>>>>>>>>>")

	daemosetsClient := clientset.AppsV1().DaemonSets(namespace)
	var daemonset = &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					//NodeSelector: map[string]string{
					//	"testnode":                "node1",
					//	"beta.kubernetes.io/arch": "arm64",
					//},
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "kube.gwunion.cn/" + image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	// Create daemonset
	//fmt.Println("Creating daemonset...")
	_, err := daemosetsClient.Create(daemonset)

	if err != nil {
		k8Err := err.(*errors.StatusError)
		//fmt.Println(k8Err.ErrStatus.Code)
		w.Data["json"] = map[string]interface{}{"code": k8Err.ErrStatus.Code, "message": "负载" + name + "已存在"}
		w.ServeJSON()
	} else {
		//fmt.Printf("Created daemonset %q.\n", result.GetObjectMeta().GetName())
		w.Data["json"] = map[string]int{"code": 20000}
		w.ServeJSON()
	}
}

// Create daemonset ...
// @Title Create daemonset
// @Description get Userinfo
// @Param daemonset body string "daemonset of the yaml or josn"
// @Param namespace query string "namespace of the daemonset"
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router /daemonset/:namespace [post]
func (w *WorkBlanceController) CreateDaemonsetByYaml() {

	daemonset := w.Ctx.Input.RequestBody
	namespace := w.Ctx.Input.Param(":namespace")
	var err error
	if strings.Index(string(daemonset), "{") == -1 {
		daemonset, err = yaml.ToJSON(daemonset)
		if err != nil {
			panic(err.Error())
		}

	}
	var daemonsetEntity appsv1.DaemonSet
	//fmt.Println(string(deployment))
	err = json.Unmarshal(daemonset, &daemonsetEntity)
	if err != nil {
		panic(err.Error())
	}

	_, err = clientset.AppsV1().DaemonSets(namespace).Create(&daemonsetEntity)
	if err != nil {
		panic(err.Error())
	} else {
		//fmt.Println("创建陈工")
		w.Data["json"] = map[string]int{"code": 20000}
		w.ServeJSON()
	}

}

// Update daemonset
// @Title update daemonset
// @Description get Userinfo
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router /daemonset [put]
func (w *WorkBlanceController) UpdateDaemonset() {

	namespace := w.Input().Get("namespace")
	daemonset := w.Input().Get("name")

	//fmt.Println(daemonset)

	var daemonset1 appsv1.DaemonSet
	err1 := json.Unmarshal([]byte(daemonset), &daemonset1)
	if err1 != nil {
		panic(err1.Error())
	}

	_, err := clientset.AppsV1().DaemonSets(namespace).Update(&daemonset1)
	if err != nil {
		//w.Data["json"] = map[string]int{"code": 400}
		//w.ServeJSON()
		panic(err.Error())
	} else {

		fmt.Println("Updated daemonset...")
		w.Data["json"] = map[string]int{"code": 20000}
		w.ServeJSON()
	}
}

// GetAll ...
// @Title Get All Daemonsets
// @Description get Daemonsets
// @Param namespace query string false "Namespace of the daemonset"
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router /daemonset [get]
func (w *WorkBlanceController) GetDaemonset() {

	//clientset := getClientset()
	daemosetsClient := clientset.AppsV1().DaemonSets(w.Input().Get("namespace"))

	list, err := daemosetsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	} else {
		w.Data["json"] = map[string]interface{}{"code": 20000, "data": list}
		w.ServeJSON()
	}

}

// @Title Singel Selete
// @Description get Single daemonset By name
// @Param name path string false "name of the daemonset"
// @Param namespace query string false "namespace of the daemonset"
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router /daemonset/:name [get]
func (w *WorkBlanceController) GetSingleDaemonset() {

	name := w.Ctx.Input.Param(":name")
	namespace := w.Input().Get("namespace")
	daemonsetClient := clientset.AppsV1().DaemonSets(namespace)

	daemonset, err := daemonsetClient.Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		w.Data["json"] = map[string]interface{}{"code": 20000}
		w.ServeJSON()
	} else if err != nil {
		panic(err)
	} else {
		w.Data["json"] = map[string]interface{}{"code": 20000, "data": daemonset}
		w.ServeJSON()
	}

}

// Delete daemonset by name
// @Title Get All
// @Description get Userinfo
// @Param name query string true "Name of daemonset"
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router /daemonset [delete]
func (w *WorkBlanceController) DeleteDaemonset() {
	// Delete Deployment

	name := w.Input().Get("name")
	namespace := w.Input().Get("namespace")
	daemontsetsClient := clientset.AppsV1().DaemonSets(namespace)
	deletePolicy := metav1.DeletePropagationForeground
	if err := daemontsetsClient.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	//fmt.Println("Deleted daemontset.")
	w.Data["json"] = map[string]int{"code": 20000}
	w.ServeJSON()

}

// func prompt() {
//  fmt.Printf("-> Press Return key to continue.")
//  scanner := bufio.NewScanner(os.Stdin)
//  for scanner.Scan() {
//   break
//  }
//  if err := scanner.Err(); err != nil {
//   panic(err)
//  }
//  fmt.Println()
// }

func int32Ptr(i int32) *int32 { return &i }

// func getClientset() *kubernetes.Clientset {
//  var kubeconfig *string
//  if home := homedir.HomeDir(); home != "" {
//   kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
//  } else {
//   kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
//  }
//  flag.Parse()

//  config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//  if err != nil {
//   panic(err)
//  }
//  clientset, err := kubernetes.NewForConfig(config)
//  if err != nil {
//   panic(err)
//  }
//  return clientset
// }

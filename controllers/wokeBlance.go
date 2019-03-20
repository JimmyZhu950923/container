package controllers

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	"strconv"

	"github.com/astaxie/beego"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

// Operations about WorkBlance
type WorkBlanceController struct {
	beego.Controller
}

// GetAll ...
// @Title Get All Deployments
// @Description get Deployments
// @Param namespace query string false "Namespace of the deployment"
// @Success 200 {object} models.Userinfo
// @Failure 403
// @router / [get]
func (w *WorkBlanceController) GetDeployments() {
	namespace := w.Input().Get("namespace")

	if namespace == "" {
		namespace = apiv1.NamespaceDefault
	}

	clientset := getClientset()
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
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

	fmt.Println(">>>>>>>>>>>>>")
	fmt.Println(name, num, image, namespace)
	fmt.Println(">>>>>>>>>>>>>")

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
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
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(deployment)

	if err != nil {
		k8Err := err.(*errors.StatusError)
		//fmt.Println(k8Err.ErrStatus.Code)
		w.Data["json"] = map[string]interface{}{"code": k8Err.ErrStatus.Code, "message": "负载" + name + "已存在"}
		w.ServeJSON()
	} else {
		fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
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
	clientset := getClientset()
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get("demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}

		result.Spec.Replicas = int32Ptr(1)                           // reduce replica count
		result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
		_, updateErr := deploymentsClient.Update(result)
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	fmt.Println("Updated deployment...")
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
	clientset := getClientset()
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete("demo-deployment", &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted deployment.")
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

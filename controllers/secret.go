package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"strings"
)

type SecretController struct {
	beego.Controller
}

// @Title getSecrets
// @Description get all Secrets,
// @Success 200 {object} models.User
// @router / [get]
func (s *SecretController) GetSecret() {
	namespace := s.Input().Get("namespace")
	secret, err := clientset.CoreV1().Secrets(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	json := map[string]interface{}{"data": secret, "code": 20000}
	s.Data["json"] = json
	s.ServeJSON()
}

// @Title getSingle
// @Description get one Secret,
// @Param namespace path string false "namespace for secret"
// @Param name query string false "name for secret"
// @Success 200 {object} models.User
// @router /:namespace [get]
func (s *SecretController) GetSingleSecret() {
	namespace := s.Ctx.Input.Param(":namespace")
	name := s.Input().Get("name")
	secret, err := clientset.CoreV1().Secrets(namespace).Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		s.Data["json"] = map[string]interface{}{"code": 20000}
		s.ServeJSON()
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting config %s in namespace %s: %v\n",
			secret, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		s.Data["json"] = map[string]interface{}{"code": 20000, "data": secret}
		s.ServeJSON()
	}
}

// @Title createSecretByYaml
// @Description create new secret
// @Param service1 body string false "secret"
// @Param namespace body string false "namespace for secret"
// @Success 200 {string} 添加成功
// @Failure 403
// @router /:namespace [post]
func (s *SecretController) CreateSecretByYaml() {
	secret := s.Ctx.Input.RequestBody
	namespace := s.Ctx.Input.Param(":namespace")
	fmt.Println("namespace = ", namespace)
	fmt.Println(string(secret))
	var err error
	if strings.Index(string(secret), "{") == -1 {
		secret, err = yaml.ToJSON(secret)
		if err != nil {
			panic(err.Error())
		}
	}

	var secretEntity v1.Secret
	err = json.Unmarshal(secret, &secretEntity)
	if err != nil {
		panic(err.Error())
	}
	_, err = clientset.CoreV1().Secrets(namespace).Create(&secretEntity)
	if err != nil {
		panic(err.Error())
	}
	json := map[string]interface{}{"code": 20000}
	s.Data["json"] = json
	s.ServeJSON()
	return
}

// @Title delSecret
// @Description delete one Secret
// @Param name query string false "name for secret"
// @Param namespace query string false "namespace for secret"
// @Success 200 {string} 删除成功
// @router / [delete]
func (s *ServicesController) DeleteSecret() {
	name := s.Input().Get("name")
	namespace := s.Input().Get("namespace")
	err := clientset.CoreV1().Secrets(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}
	s.Data["json"] = map[string]interface{}{"code": 20000}
	s.ServeJSON()
}

// @Title updSecret
// @Description updeate secret
// @Success 200 {string} 更新成功
// @router / [put]
func (s *SecretController) UpdateSecret() {
	namespace := s.Input().Get("namespace")
	name := s.Input().Get("name")
	var secret v1.Secret
	err1 := json.Unmarshal([]byte(name), &secret)
	if err1 != nil {
		panic(err1.Error())
	}

	_, err := clientset.CoreV1().Secrets(namespace).Update(&secret)

	if err != nil {
		s.Data["json"] = map[string]interface{}{"code": 400, "data": err.Error()}
		s.ServeJSON()
	} else {
		s.Data["json"] = map[string]interface{}{"code": 20000}
		s.ServeJSON()
	}
}

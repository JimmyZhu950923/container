package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"k8s.io/apimachinery/pkg/api/errors"
)
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ConfigController struct {
	beego.Controller
}

// @Title getConfigs
// @Description get all Configs,
// @Success 200 {object} models.User
// @router / [get]
func (c *ConfigController) GetConfig() {
	namespace := c.Input().Get("namespace")
	config, err := clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	json := map[string]interface{}{"data": config, "code": 20000}
	c.Data["json"] = json
	c.ServeJSON()
}

// @Title getSingle
// @Description get one Config,
// @Param namespace path string false "namespace for config"
// @Param name query string false "name for config"
// @Success 200 {object} models.User
// @router /:namespace [get]
func (c *ConfigController) GetSingleConfig() {
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Input().Get("name")
	config, err := clientset.CoreV1().ConfigMaps(namespace).Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		c.Data["json"] = map[string]interface{}{"code": 20000}
		c.ServeJSON()
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting config %s in namespace %s: %v\n",
			config, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		c.Data["json"] = map[string]interface{}{"code": 20000, "data": config}
		c.ServeJSON()
	}
}


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

// @Title createConfigByYaml
// @Description create new config
// @Param service1 body string false "config"
// @Param namespace body string false "namespace for config"
// @Success 200 {string} 添加成功
// @Failure 403
// @router /:namespace [post]
func (c *ConfigController) CreateConfigByYaml() {
	config := c.Ctx.Input.RequestBody
	namespace := c.Ctx.Input.Param(":namespace")
	fmt.Println("namespace = ", namespace)
	fmt.Println(string(config))
	var err error
	if strings.Index(string(config), "{") == -1 {
		config, err = yaml.ToJSON(config)
		if err != nil {
			panic(err.Error())
		}
	}

	var configEntity v1.ConfigMap
	err = json.Unmarshal(config, &configEntity)
	if err != nil {
		panic(err.Error())
	}
	_, err = clientset.CoreV1().ConfigMaps(namespace).Create(&configEntity)
	if err != nil {
		panic(err.Error())
	}
	json := map[string]interface{}{"code": 20000}
	c.Data["json"] = json
	c.ServeJSON()
	return
}

// @Title delConfig
// @Description delete one config
// @Param name query string false "name for config"
// @Param namespace query string false "namespace for config"
// @Success 200 {string} 删除成功
// @router / [delete]
func (c *ConfigController) DeleteConfig() {
	name := c.Input().Get("name")
	namespace := c.Input().Get("namespace")
	err := clientset.CoreV1().ConfigMaps(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}
	c.Data["json"] = map[string]interface{}{"code": 20000}
	c.ServeJSON()
}

// @Title updConfig
// @Description updeate config
// @Success 200 {string} 更新成功
// @router / [put]
func (c *ConfigController) UpdateConfig() {
	namespace := c.Input().Get("namespace")
	name := c.Input().Get("name")
	var config1 v1.ConfigMap
	err1 := json.Unmarshal([]byte(name), &config1)
	if err1 != nil {
		panic(err1.Error())
	}

	_, err := clientset.CoreV1().ConfigMaps(namespace).Update(&config1)

	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 400, "data": err.Error()}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 20000}
		c.ServeJSON()
	}
}

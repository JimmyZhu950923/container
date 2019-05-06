package controllers

import (
	"gt-container-go/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ClusterController struct {
	beego.Controller
}

// @Title Insert Cluster
// @Description Add a cluster
//@Param name query string false "cluster's name"
//@Param displayName query string false "cluster's displayName"
//@Param namespace query string false "cluster's namespace"
//@Param apiService query string false "cluster's api_service"
// @Success 200 {object} models.User
//@router / [post]
func (c *ClusterController) Insert() {
	name := c.Input().Get("name")
	namespace := c.Input().Get("namespace")
	err := models.FindClusterByName(name)
	if err == orm.ErrNoRows {
		var cluster = models.Cluster{Name: name, Namespace: namespace, CreateDate: time.Now()}
		id, _ := models.Insert(&cluster)
		c.Data["json"] = map[string]interface{}{"code": 20000, "data": id}
		c.ServeJSON()
		return
	} else {
		panic(err.Error())
		c.Data["json"] = map[string]interface{}{"code": 20000, "message": name + "已存在"}
		c.ServeJSON()
		return
	}

}

// @Title GetAll
// @Description get all Cluster
// @Success 200 {object} models.User
//@router / [get]
func (c *ClusterController) Select() {
	var clusters []models.Cluster
	row, _ := models.Select(&clusters)
	c.Data["json"] = map[string]interface{}{"code": 20000, "data": clusters, "row": row}
	c.ServeJSON()
	return
}

// @Title Get a cluster
// @Description get cluster by namespace
//@Param namespace query string false "cluster's namespace"
// @Success 200 {object} models.User
//@router /:namespace [get]
func (c *ClusterController) GetByNamespace() {
	var clusters []models.Cluster
	namespace := c.Ctx.Input.Param(":namespace")
	row, _ := models.FindByNamespace(namespace, &clusters)
	c.Data["json"] = map[string]interface{}{"code": 20000, "data": clusters, "row": row}
	c.ServeJSON()
	return
}

// @Title Delete a cluster
// @Description Delete cluster by id
//@Param id query string false "cluster's id"
// @Success 200 {object} models.User
// @router /:id [delete]
func (c *ClusterController) DeleteById() {
	str := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(str, 10, 64)
	num, _ := models.DeleteCluster(id)
	c.Data["json"] = map[string]interface{}{"code": 20000, "num": num}
	c.ServeJSON()
	return
}

//@router /:name [get]
func (c *ClusterController) FindByName() {
	name := c.Ctx.Input.Param(":name")
	err := models.FindClusterByName(name)
	if err == orm.ErrNoRows {
		c.Data["json"] = map[string]interface{}{"code": 20000}
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = map[string]interface{}{"code": 400}
		c.ServeJSON()
		return
	}
}

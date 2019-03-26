package controllers

import (
	"github.com/astaxie/beego"
	corev1 "k8s.io/api/core/v1"
	storageV1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Operations about Users
type StorageController struct {
	beego.Controller
}

// @Title getStorages
// @Description get all Storages,
// @Success 200 {object} models.User
// @router / [get]
func (s *StorageController) GetStorages() {
	storageClassClient := clientset.StorageV1().StorageClasses()
	storage, err := storageClassClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	json := map[string]interface{}{"data": storage, "code": 20000}
	s.Data["json"] = json
	s.ServeJSON()
}

// @Title getSingle
// @Description get one Storage,
// @Param name query string false "name for storage"
// @Success 200 {object} models.User
// @router /:name [get]
func (s *StorageController) GetSingleStorage() {
	name := s.Ctx.Input.Param(":name")
	storageClassClient := clientset.StorageV1().StorageClasses()
	storage, err := storageClassClient.Get(name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	json := map[string]interface{}{"data": storage, "code": 20000}
	s.Data["json"] = json
	s.ServeJSON()
}

// @Title createStorage
// @Description create new storage
// @Success 200 {string} 添加成功
// @router /create [get]
func (s *StorageController) CreateStorage() {
	name := s.Input().Get("name")
	namespace := s.Input().Get("namespace")
	persistentVolumeReclaimPolicy := corev1.PersistentVolumeReclaimPolicy(s.Input().Get("persistentVolumeReclaimPolicy"))
	storageClassClient := clientset.StorageV1().StorageClasses()
	var storage = &storageV1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace:	namespace,
		},
		Provisioner:"",
		Parameters: map[string]string{},
		ReclaimPolicy:&persistentVolumeReclaimPolicy,
	}
	storage1, err := storageClassClient.Create(storage)
	if err != nil {
		panic(err.Error())
	}
	json := map[string]interface{}{"data": storage1, "code": 20000}
	s.Data["json"] = json
	s.ServeJSON()
	return
}

// @Title delStorage
// @Description delete one storage
// @Param name query string false "name for storage"
// @Success 200 {string} 删除成功
// @router / [delete]
func (s *StorageController) DelStorage() {
	name := s.Input().Get("name")
	storageClassClient := clientset.StorageV1().StorageClasses()
	err := storageClassClient.Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}
	json := map[string]interface{}{"code": 20000}
	s.Data["json"] = json
	s.ServeJSON()
}
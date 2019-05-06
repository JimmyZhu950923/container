package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	corev1 "k8s.io/api/core/v1"
	storageV1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"strings"
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
	persistentVolumeReclaimPolicy := corev1.PersistentVolumeReclaimPolicy(s.Input().Get("persistentVolumeReclaimPolicy"))
	storageClassClient := clientset.StorageV1().StorageClasses()
	var storage = &storageV1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
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

// @Title createStorageByYaml
// @Description create new storage
// @Param service1 body string false "storage"
// @Param namespace body string false "namespace for storage"
// @Success 200 {string} 添加成功
// @Failure 403
// @router /:namespace [post]
func (s *StorageController) CreateStorageByYaml() {
	storage := s.Ctx.Input.RequestBody
	var err error
	if strings.Index(string(storage), "{") == -1 {
		storage, err = yaml.ToJSON(storage)
		if err != nil {
			panic(err.Error())
		}
	}

	var storageEntity storageV1.StorageClass
	err = json.Unmarshal(storage, &storageEntity)
	if err != nil {
		panic(err.Error())
	}
	_, err = clientset.StorageV1().StorageClasses().Create(&storageEntity)
	if err != nil {
		panic(err.Error())
	}
	json := map[string]interface{}{"code": 20000}
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

// @Title updStorage
// @Description updeate storage
// @Success 200 {string} 更新成功
// @router / [put]
func (s *StorageController) UpdateStorage() {
	name := s.Input().Get("name")
	var storage1 storageV1.StorageClass
	err1 := json.Unmarshal([]byte(name), &storage1)
	if err1 != nil {
		panic(err1.Error())
	}

	_, err := clientset.StorageV1().StorageClasses().Update(&storage1)

	if err != nil {
		s.Data["json"] = map[string]interface{}{"code": 400, "data": err.Error()}
		s.ServeJSON()
	} else {
		s.Data["json"] = map[string]interface{}{"code": 20000}
		s.ServeJSON()
	}
}

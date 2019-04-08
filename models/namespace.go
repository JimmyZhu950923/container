package models

import (
	"time"
)

type Namespace struct {
	Id         int64
	Name       string
	Status     string
	CreateTime time.Time
}

//查询所有
func GetAllNsp(namespaces *[]Namespace) (row int64, err error) {
	row, err = Db.QueryTable("namespace").All(namespaces)
	return
}

//根据name模糊查询
func GetSingleNsp(name string, namespace *[]Namespace) (row int64, err error) {
	qs := Db.QueryTable("namespace")
	qs = qs.Filter("name__icontains", name)
	row, err = qs.All(namespace)
	return
}

//插入namespace
func InsertNsp(namespace *Namespace) (id int64, err error) {
	Db.Begin()
	id, err = Db.Insert(namespace)
	return
}

//根据name查询namespace
func FindNspByName(name string) (err error) {
	var namespace = Namespace{Name: name}
	err = Db.Read(&namespace, "Name")
	return
}

//根菌name删除namespace
func DeleteNsp(name string) (num int64, err error) {
	Db.Begin()
	num, err = Db.QueryTable("namespace").Filter("name", name).Delete()
	return
}

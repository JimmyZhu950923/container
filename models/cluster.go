package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

type Cluster struct {
	Id          int64
	Name        string
	DisplayName string
	Namespace   string
	CreateDate  time.Time
	UpdateDate  time.Time
	ApiService  string
}

var Db orm.Ormer

func init() {
	
	driver := beego.AppConfig.String("mysql::driver")
	dataSource := beego.AppConfig.String("mysql::dataSource")
	params, _ := beego.AppConfig.Int("mysql::params")
	orm.RegisterDataBase("default", driver, dataSource, params)
	orm.RegisterModel(new(Cluster))
	orm.RegisterModel(new(Namespace))
	orm.DefaultTimeLoc = time.UTC
	Db = orm.NewOrm()
}

//添加集群
func Insert(cluster *Cluster) (id int64, err error) {
	id, err = Db.Insert(cluster)
	return
}

//查询集群
func Select(clusters *[]Cluster) (row int64, err error) {
	row, err = Db.QueryTable("cluster").All(clusters)
	return
}

//按照namespace查询
func FindByNamespace(namespace string, clusters *[]Cluster) (row int64, err error) {
	qs := Db.QueryTable("cluster")
	qs = qs.Filter("namespace__icontains", namespace)
	row, err = qs.All(clusters)
	return
}

//删除集群
func DeleteCluster(id int64) (num int64, err error) {
	num, err = Db.Delete(&Cluster{Id: id})
	return
}

//根据名字查询cluster
func FindClusterByName(name string) (err error) {
	var cluster = Cluster{Name: name}
	err = Db.Read(&cluster, "Name")
	return
}

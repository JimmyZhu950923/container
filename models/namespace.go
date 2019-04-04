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

func GetAllNsp(namespaces *[]Namespace) (row int64, err error) {
	row, err = db.QueryTable("namespace").All(namespaces)
	return
}

package models

import (
	"reflect"
	"time"
)

type Document struct {
	Id     string                 `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string                 `bson:"name" json:"name"`
	Data   map[string]interface{} `bson:"data" json:"data"`
	Number int                    `bson:"number" json:"number"`
	Date   time.Time              `bson:"date" json:"date"`
}

func (d *Document) IsEmpty() bool {
	return reflect.DeepEqual(*d, Document{})
}

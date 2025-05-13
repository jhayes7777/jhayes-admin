
package model
import (
)

// Asdasdas 表d 结构体
type Asdasdas struct {
  Sadasd  *string `json:"sadasd" form:"sadasd" gorm:"primarykey;column:sadasd;"`  //xzczxczx
}


// TableName 表d Asdasdas自定义表名 asdasdas
func (Asdasdas) TableName() string {
    return "asdasdas"
}








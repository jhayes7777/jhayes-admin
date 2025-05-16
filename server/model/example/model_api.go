
// 自动生成模板ModelManage
package example
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 模型接口 结构体  ModelManage
type ModelManage struct {
    global.GVA_MODEL
  CoverImg  *string `json:"coverImg" form:"coverImg" gorm:"column:cover_img;"`  //封面
  ModelUrl  *string `json:"modelUrl" form:"modelUrl" gorm:"column:model_url;"`  //模型地址
}


// TableName 模型接口 ModelManage自定义表名 model_manage
func (ModelManage) TableName() string {
    return "model_manage"
}






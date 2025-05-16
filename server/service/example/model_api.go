
package example

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
    exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
)

type ModelManageService struct {}
// CreateModelManage 创建模型接口记录
// Author [yourname](https://github.com/yourname)
func (MMService *ModelManageService) CreateModelManage(ctx context.Context, MM *example.ModelManage) (err error) {
	err = global.GVA_DB.Create(MM).Error
	return err
}

// DeleteModelManage 删除模型接口记录
// Author [yourname](https://github.com/yourname)
func (MMService *ModelManageService)DeleteModelManage(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&example.ModelManage{},"id = ?",ID).Error
	return err
}

// DeleteModelManageByIds 批量删除模型接口记录
// Author [yourname](https://github.com/yourname)
func (MMService *ModelManageService)DeleteModelManageByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]example.ModelManage{},"id in ?",IDs).Error
	return err
}

// UpdateModelManage 更新模型接口记录
// Author [yourname](https://github.com/yourname)
func (MMService *ModelManageService)UpdateModelManage(ctx context.Context, MM example.ModelManage) (err error) {
	err = global.GVA_DB.Model(&example.ModelManage{}).Where("id = ?",MM.ID).Updates(&MM).Error
	return err
}

// GetModelManage 根据ID获取模型接口记录
// Author [yourname](https://github.com/yourname)
func (MMService *ModelManageService)GetModelManage(ctx context.Context, ID string) (MM example.ModelManage, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&MM).Error
	return
}
// GetModelManageInfoList 分页获取模型接口记录
// Author [yourname](https://github.com/yourname)
func (MMService *ModelManageService)GetModelManageInfoList(ctx context.Context, info exampleReq.ModelManageSearch) (list []example.ModelManage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&example.ModelManage{})
    var MMs []example.ModelManage
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&MMs).Error
	return  MMs, total, err
}
func (MMService *ModelManageService)GetModelManagePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}

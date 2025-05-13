
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model/request"
)

var Asdasdas = new(asdsadsa)

type asdsadsa struct {}
// CreateAsdasdas 创建表d记录
// Author [yourname](https://github.com/yourname)
func (s *asdsadsa) CreateAsdasdas(ctx context.Context, asdsadsa *model.Asdasdas) (err error) {
	err = global.GVA_DB.Create(asdsadsa).Error
	return err
}

// DeleteAsdasdas 删除表d记录
// Author [yourname](https://github.com/yourname)
func (s *asdsadsa) DeleteAsdasdas(ctx context.Context, sadasd string) (err error) {
	err = global.GVA_DB.Delete(&model.Asdasdas{},"sadasd = ?",sadasd).Error
	return err
}

// DeleteAsdasdasByIds 批量删除表d记录
// Author [yourname](https://github.com/yourname)
func (s *asdsadsa) DeleteAsdasdasByIds(ctx context.Context, sadasds []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Asdasdas{},"sadasd in ?",sadasds).Error
	return err
}

// UpdateAsdasdas 更新表d记录
// Author [yourname](https://github.com/yourname)
func (s *asdsadsa) UpdateAsdasdas(ctx context.Context, asdsadsa model.Asdasdas) (err error) {
	err = global.GVA_DB.Model(&model.Asdasdas{}).Where("sadasd = ?",asdsadsa.Sadasd).Updates(&asdsadsa).Error
	return err
}

// GetAsdasdas 根据sadasd获取表d记录
// Author [yourname](https://github.com/yourname)
func (s *asdsadsa) GetAsdasdas(ctx context.Context, sadasd string) (asdsadsa model.Asdasdas, err error) {
	err = global.GVA_DB.Where("sadasd = ?", sadasd).First(&asdsadsa).Error
	return
}
// GetAsdasdasInfoList 分页获取表d记录
// Author [yourname](https://github.com/yourname)
func (s *asdsadsa) GetAsdasdasInfoList(ctx context.Context, info request.AsdasdasSearch) (list []model.Asdasdas, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Asdasdas{})
    var asdsadsas []model.Asdasdas
    // 如果有条件搜索 下方会自动创建搜索语句
  
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&asdsadsas).Error
	return  asdsadsas, total, err
}

func (s *asdsadsa)GetAsdasdasPublic(ctx context.Context) {

}

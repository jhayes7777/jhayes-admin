package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Asdasdas = new(asdsadsa)

type asdsadsa struct {}

// Init 初始化 表d 路由信息
func (r *asdsadsa) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("asdsadsa").Use(middleware.OperationRecord())
		group.POST("createAsdasdas", apiAsdasdas.CreateAsdasdas)   // 新建表d
		group.DELETE("deleteAsdasdas", apiAsdasdas.DeleteAsdasdas) // 删除表d
		group.DELETE("deleteAsdasdasByIds", apiAsdasdas.DeleteAsdasdasByIds) // 批量删除表d
		group.PUT("updateAsdasdas", apiAsdasdas.UpdateAsdasdas)    // 更新表d
	}
	{
	    group := private.Group("asdsadsa")
		group.GET("findAsdasdas", apiAsdasdas.FindAsdasdas)        // 根据ID获取表d
		group.GET("getAsdasdasList", apiAsdasdas.GetAsdasdasList)  // 获取表d列表
	}
	{
	    group := public.Group("asdsadsa")
	    group.GET("getAsdasdasPublic", apiAsdasdas.GetAsdasdasPublic)  // 表d开放接口
	}
}

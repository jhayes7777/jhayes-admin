package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ModelManageRouter struct {}

// InitModelManageRouter 初始化 模型接口 路由信息
func (s *ModelManageRouter) InitModelManageRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	MMRouter := Router.Group("MM").Use(middleware.OperationRecord())
	MMRouterWithoutRecord := Router.Group("MM")
	MMRouterWithoutAuth := PublicRouter.Group("MM")
	{
		MMRouter.POST("createModelManage", MMApi.CreateModelManage)   // 新建模型接口
		MMRouter.DELETE("deleteModelManage", MMApi.DeleteModelManage) // 删除模型接口
		MMRouter.DELETE("deleteModelManageByIds", MMApi.DeleteModelManageByIds) // 批量删除模型接口
		MMRouter.PUT("updateModelManage", MMApi.UpdateModelManage)    // 更新模型接口
	}
	{
		MMRouterWithoutRecord.GET("findModelManage", MMApi.FindModelManage)        // 根据ID获取模型接口
		MMRouterWithoutRecord.GET("getModelManageList", MMApi.GetModelManageList)  // 获取模型接口列表
	}
	{
	    MMRouterWithoutAuth.GET("getModelManagePublic", MMApi.GetModelManagePublic)  // 模型接口开放接口
	}
}

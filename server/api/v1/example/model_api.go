package example

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/example"
    exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ModelManageApi struct {}



// CreateModelManage 创建模型接口
// @Tags ModelManage
// @Summary 创建模型接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.ModelManage true "创建模型接口"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /MM/createModelManage [post]
func (MMApi *ModelManageApi) CreateModelManage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var MM example.ModelManage
	err := c.ShouldBindJSON(&MM)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = MMService.CreateModelManage(ctx,&MM)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteModelManage 删除模型接口
// @Tags ModelManage
// @Summary 删除模型接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.ModelManage true "删除模型接口"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /MM/deleteModelManage [delete]
func (MMApi *ModelManageApi) DeleteModelManage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := MMService.DeleteModelManage(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteModelManageByIds 批量删除模型接口
// @Tags ModelManage
// @Summary 批量删除模型接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /MM/deleteModelManageByIds [delete]
func (MMApi *ModelManageApi) DeleteModelManageByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := MMService.DeleteModelManageByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateModelManage 更新模型接口
// @Tags ModelManage
// @Summary 更新模型接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.ModelManage true "更新模型接口"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /MM/updateModelManage [put]
func (MMApi *ModelManageApi) UpdateModelManage(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var MM example.ModelManage
	err := c.ShouldBindJSON(&MM)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = MMService.UpdateModelManage(ctx,MM)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindModelManage 用id查询模型接口
// @Tags ModelManage
// @Summary 用id查询模型接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询模型接口"
// @Success 200 {object} response.Response{data=example.ModelManage,msg=string} "查询成功"
// @Router /MM/findModelManage [get]
func (MMApi *ModelManageApi) FindModelManage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reMM, err := MMService.GetModelManage(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reMM, c)
}
// GetModelManageList 分页获取模型接口列表
// @Tags ModelManage
// @Summary 分页获取模型接口列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.ModelManageSearch true "分页获取模型接口列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /MM/getModelManageList [get]
func (MMApi *ModelManageApi) GetModelManageList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo exampleReq.ModelManageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := MMService.GetModelManageInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetModelManagePublic 不需要鉴权的模型接口接口
// @Tags ModelManage
// @Summary 不需要鉴权的模型接口接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /MM/getModelManagePublic [get]
func (MMApi *ModelManageApi) GetModelManagePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    MMService.GetModelManagePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的模型接口接口信息",
    }, "获取成功", c)
}

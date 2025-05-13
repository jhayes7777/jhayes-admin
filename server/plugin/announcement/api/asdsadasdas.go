package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var Asdasdas = new(asdsadsa)

type asdsadsa struct {}

// CreateAsdasdas 创建表d
// @Tags Asdasdas
// @Summary 创建表d
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Asdasdas true "创建表d"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /asdsadsa/createAsdasdas [post]
func (a *asdsadsa) CreateAsdasdas(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Asdasdas
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAsdasdas.CreateAsdasdas(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAsdasdas 删除表d
// @Tags Asdasdas
// @Summary 删除表d
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Asdasdas true "删除表d"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /asdsadsa/deleteAsdasdas [delete]
func (a *asdsadsa) DeleteAsdasdas(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	sadasd := c.Query("sadasd")
	err := serviceAsdasdas.DeleteAsdasdas(ctx,sadasd)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteAsdasdasByIds 批量删除表d
// @Tags Asdasdas
// @Summary 批量删除表d
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /asdsadsa/deleteAsdasdasByIds [delete]
func (a *asdsadsa) DeleteAsdasdasByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	sadasds := c.QueryArray("sadasds[]")
	err := serviceAsdasdas.DeleteAsdasdasByIds(ctx,sadasds)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateAsdasdas 更新表d
// @Tags Asdasdas
// @Summary 更新表d
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Asdasdas true "更新表d"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /asdsadsa/updateAsdasdas [put]
func (a *asdsadsa) UpdateAsdasdas(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Asdasdas
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAsdasdas.UpdateAsdasdas(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindAsdasdas 用id查询表d
// @Tags Asdasdas
// @Summary 用id查询表d
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param sadasd query string true "用id查询表d"
// @Success 200 {object} response.Response{data=model.Asdasdas,msg=string} "查询成功"
// @Router /asdsadsa/findAsdasdas [get]
func (a *asdsadsa) FindAsdasdas(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	sadasd := c.Query("sadasd")
	reasdsadsa, err := serviceAsdasdas.GetAsdasdas(ctx,sadasd)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(reasdsadsa, c)
}
// GetAsdasdasList 分页获取表d列表
// @Tags Asdasdas
// @Summary 分页获取表d列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.AsdasdasSearch true "分页获取表d列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /asdsadsa/getAsdasdasList [get]
func (a *asdsadsa) GetAsdasdasList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.AsdasdasSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceAsdasdas.GetAsdasdasInfoList(ctx,pageInfo)
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
// GetAsdasdasPublic 不需要鉴权的表d接口
// @Tags Asdasdas
// @Summary 不需要鉴权的表d接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /asdsadsa/getAsdasdasPublic [get]
func (a *asdsadsa) GetAsdasdasPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceAsdasdas.GetAsdasdasPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的表d接口信息"}, "获取成功", c)
}

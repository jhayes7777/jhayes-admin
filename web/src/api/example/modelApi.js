import service from '@/utils/request'
// @Tags ModelManage
// @Summary 创建模型接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ModelManage true "创建模型接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /MM/createModelManage [post]
export const createModelManage = (data) => {
  return service({
    url: '/MM/createModelManage',
    method: 'post',
    data
  })
}

// @Tags ModelManage
// @Summary 删除模型接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ModelManage true "删除模型接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MM/deleteModelManage [delete]
export const deleteModelManage = (params) => {
  return service({
    url: '/MM/deleteModelManage',
    method: 'delete',
    params
  })
}

// @Tags ModelManage
// @Summary 批量删除模型接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除模型接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MM/deleteModelManage [delete]
export const deleteModelManageByIds = (params) => {
  return service({
    url: '/MM/deleteModelManageByIds',
    method: 'delete',
    params
  })
}

// @Tags ModelManage
// @Summary 更新模型接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ModelManage true "更新模型接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MM/updateModelManage [put]
export const updateModelManage = (data) => {
  return service({
    url: '/MM/updateModelManage',
    method: 'put',
    data
  })
}

// @Tags ModelManage
// @Summary 用id查询模型接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ModelManage true "用id查询模型接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MM/findModelManage [get]
export const findModelManage = (params) => {
  return service({
    url: '/MM/findModelManage',
    method: 'get',
    params
  })
}

// @Tags ModelManage
// @Summary 分页获取模型接口列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取模型接口列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MM/getModelManageList [get]
export const getModelManageList = (params) => {
  return service({
    url: '/MM/getModelManageList',
    method: 'get',
    params
  })
}

// @Tags ModelManage
// @Summary 不需要鉴权的模型接口接口
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.ModelManageSearch true "分页获取模型接口列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /MM/getModelManagePublic [get]
export const getModelManagePublic = () => {
  return service({
    url: '/MM/getModelManagePublic',
    method: 'get',
  })
}

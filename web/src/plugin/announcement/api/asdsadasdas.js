import service from '@/utils/request'
// @Tags Asdasdas
// @Summary 创建表d
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Asdasdas true "创建表d"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /asdsadsa/createAsdasdas [post]
export const createAsdasdas = (data) => {
  return service({
    url: '/asdsadsa/createAsdasdas',
    method: 'post',
    data
  })
}

// @Tags Asdasdas
// @Summary 删除表d
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Asdasdas true "删除表d"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /asdsadsa/deleteAsdasdas [delete]
export const deleteAsdasdas = (params) => {
  return service({
    url: '/asdsadsa/deleteAsdasdas',
    method: 'delete',
    params
  })
}

// @Tags Asdasdas
// @Summary 批量删除表d
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除表d"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /asdsadsa/deleteAsdasdas [delete]
export const deleteAsdasdasByIds = (params) => {
  return service({
    url: '/asdsadsa/deleteAsdasdasByIds',
    method: 'delete',
    params
  })
}

// @Tags Asdasdas
// @Summary 更新表d
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Asdasdas true "更新表d"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /asdsadsa/updateAsdasdas [put]
export const updateAsdasdas = (data) => {
  return service({
    url: '/asdsadsa/updateAsdasdas',
    method: 'put',
    data
  })
}

// @Tags Asdasdas
// @Summary 用id查询表d
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Asdasdas true "用id查询表d"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /asdsadsa/findAsdasdas [get]
export const findAsdasdas = (params) => {
  return service({
    url: '/asdsadsa/findAsdasdas',
    method: 'get',
    params
  })
}

// @Tags Asdasdas
// @Summary 分页获取表d列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取表d列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /asdsadsa/getAsdasdasList [get]
export const getAsdasdasList = (params) => {
  return service({
    url: '/asdsadsa/getAsdasdasList',
    method: 'get',
    params
  })
}
// @Tags Asdasdas
// @Summary 不需要鉴权的表d接口
// @Accept application/json
// @Produce application/json
// @Param data query request.AsdasdasSearch true "分页获取表d列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /asdsadsa/getAsdasdasPublic [get]
export const getAsdasdasPublic = () => {
  return service({
    url: '/asdsadsa/getAsdasdasPublic',
    method: 'get',
  })
}

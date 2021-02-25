package controller

import (
	"Seaman/model"
	"Seaman/service"
	"Seaman/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strconv"
)

/**
 * 群组控制器结构体：用来实现处理群组模块的接口的请求，并返回给客户端
 */
type AppGroupController struct {
	//上下文对象
	Ctx iris.Context

	AppGroupService service.AppGroupService
	//session对象
	Session *sessions.Session
}

func (agc *AppGroupController) BeforeActivation(a mvc.BeforeActivation) {

	//添加群组
	a.Handle("POST", "/", "PostAddAppGroup")

	//删除群组
	a.Handle("DELETE", "/{id}", "DeleteAppGroup")

	//查询群组
	a.Handle("GET", "single/{id}", "GetAppGroup")

	//查询群组数量
	a.Handle("GET", "/count", "GetCount")

	//所有群组列表
	a.Handle("GET", "/list", "GetList")

	//带参数查询群组分页
	a.Handle("GET", "/pageList", "GetPageList")

	//修改群组
	a.Handle("PUT", "/", "UpdateAppGroup")

	//查询群组人员配置
	a.Handle("GET", "groupMembers/{id}", "GetGroupMembers")

	//查询群组人员配置
	a.Handle("PUT", "updateGroupMembers", "PutGroupMembers")
}

/**
 * url: /appGroup/addappGroup
 * type：post
 * descs：添加群组
 */
func (agc *AppGroupController) PostAddAppGroup() mvc.Result {

	var appGroup model.TplAppGroupT
	err := agc.Ctx.ReadJSON(&appGroup)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}
	newAppGroup := &model.TplAppGroupT{
		Name:      appGroup.Name,
		Desp:      appGroup.Desp,
		Status:    appGroup.Status,
		Defaultin: appGroup.Defaultin,
		Code:      appGroup.Code,
	}
	isSuccess := agc.AppGroupService.AddAppGroup(newAppGroup)
	if !isSuccess {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ADD),
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"success": utils.Recode2Text(utils.RESPMSG_SUCCESS_ADD),
		},
	}
}

/**
 * 删除群组
 */
func (agc *AppGroupController) DeleteAppGroup() mvc.Result {

	id := agc.Ctx.Params().Get("id")

	appGroupId, err := strconv.Atoi(id)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}
	delete := agc.AppGroupService.DeleteAppGroup(appGroupId)
	if !delete {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_DELETE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_DELETE),
			},
		}
	} else {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_OK,
				"type":    utils.RESPMSG_SUCCESS_DELETE,
				"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_DELETE),
			},
		}
	}
}

/**
 * 获取群组
 * 请求类型：Get
 */
func (agc *AppGroupController) GetAppGroup() mvc.Result {

	id := agc.Ctx.Params().Get("id")

	appGroupId, err := strconv.Atoi(id)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}

	appGroup, err := agc.AppGroupService.GetAppGroup(appGroupId)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//返回群组
	return mvc.Response{
		Object: appGroup.AppGroupTToRespDesc(),
	}
}

/**
 * 获取群组数量
 * 请求类型：Get
 */
func (agc *AppGroupController) GetCount() mvc.Result {
	//获取页面参数
	name := agc.Ctx.FormValue("name")
	code := agc.Ctx.FormValue("code")
	status := agc.Ctx.FormValue("status")
	iStatus, err := strconv.Atoi(status)
	sType := agc.Ctx.FormValue("type")
	appGroupParam := &model.TplAppGroupT{
		Name:   name,
		Code:   code,
		Status: iStatus,
		Type:   sType,
	}
	//群组总数
	total, err := agc.AppGroupService.GetAppGroupTotalCount(appGroupParam)

	//请求出现错误
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	//正常情况的返回值
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  total,
		},
	}
}

/**
 * 获取群组列表
 * 请求类型：Get
 */
func (agc *AppGroupController) GetGroupMembers() mvc.Result {

	id := agc.Ctx.Params().Get("id")

	appGroupId, err := strconv.Atoi(id)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}
	userList := agc.AppGroupService.GetGroupMembers(appGroupId)
	if len(userList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//将查询到的群组数据进行转换成前端需要的内容
	var respList []interface{}
	for _, user := range userList {
		respList = append(respList, user.UserToRespDesc())
	}

	//返回群组列表
	return mvc.Response{
		Object: &respList,
	}
}

/**
 * 获取群组列表
 * 请求类型：Get
 */
func (agc *AppGroupController) GetList() mvc.Result {

	appGroupList := agc.AppGroupService.GetAppGroupList()
	if len(appGroupList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//将查询到的群组数据进行转换成前端需要的内容
	var respList []interface{}
	for _, appGroup := range appGroupList {
		respList = append(respList, appGroup.AppGroupTToRespDesc())
	}

	//返回群组列表
	return mvc.Response{
		Object: &respList,
	}
}

/**
 * 获取群组带参数分页查询
 * 请求类型：Get
 */
func (agc *AppGroupController) GetPageList() mvc.Result {

	offsetStr := agc.Ctx.FormValue("current")
	limitStr := agc.Ctx.FormValue("pageSize")
	var offset int
	var limit int

	//判断offset和limit两个变量任意一个都不能为""
	if offsetStr == "" || limitStr == "" {

		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}

	offset, err := strconv.Atoi(offsetStr)
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}

	//做页数的限制检查
	if offset <= 0 {
		offset = 0
	} else {
		offset--
	}

	//做最大的限制
	if limit > MaxLimit {
		limit = MaxLimit
	}

	//获取页面参数
	name := agc.Ctx.FormValue("name")
	code := agc.Ctx.FormValue("code")
	status := agc.Ctx.FormValue("status")
	iStatus, err := strconv.Atoi(status)
	sType := agc.Ctx.FormValue("type")
	appGroupParam := &model.TplAppGroupT{
		Name:   name,
		Code:   code,
		Status: iStatus,
		Type:   sType,
	}
	appGroupList := agc.AppGroupService.GetAppGroupPageList(appGroupParam, offset, limit)
	total, _ := agc.AppGroupService.GetAppGroupTotalCount(appGroupParam)
	if len(appGroupList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//将查询到的群组数据进行转换成前端需要的内容
	var respList []interface{}
	for _, appGroup := range appGroupList {
		respList = append(respList, appGroup.AppGroupTToRespDesc())
	}

	//返回群组列表
	return mvc.Response{
		Object: map[string]interface{}{
			"data":     respList,
			"total":    total,
			"success":  true,
			"pageSize": limit,
			"current":  offset,
		},
	}
}

/**
 * type：put
 * descs：修改群组
 */
func (agc *AppGroupController) UpdateAppGroup() mvc.Result {

	var appGroup model.TplAppGroupT
	err := agc.Ctx.ReadJSON(&appGroup)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ADD),
			},
		}
	}
	newAppGroup := &model.TplAppGroupT{
		Id:        appGroup.Id,
		Name:      appGroup.Name,
		Desp:      appGroup.Desp,
		Status:    appGroup.Status,
		Defaultin: appGroup.Defaultin,
		Code:      appGroup.Code,
		Type:      appGroup.Type,
	}
	isSuccess := agc.AppGroupService.UpdateAppGroup(newAppGroup)
	if !isSuccess {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_UPDATE),
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"success": utils.Recode2Text(utils.RESPMSG_SUCCESS_UPDATE),
		},
	}
}

/**
 * type：put
 * descs：修改群组成员
 */
func (agc *AppGroupController) PutGroupMembers() mvc.Result {
	type UserGroupParams struct {
		UserIds string
		GroupId int64
	}
	var userGroupParams UserGroupParams
	err := agc.Ctx.ReadJSON(&userGroupParams)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ADD),
			},
		}
	}
	isSuccess := agc.AppGroupService.UpdateGroupMembers(userGroupParams.GroupId, userGroupParams.UserIds)
	if !isSuccess {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_UPDATE),
			},
		}
	}
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"success": utils.Recode2Text(utils.RESPMSG_SUCCESS_UPDATE),
		},
	}
}

package controller

import (
	"Seaman/model"
	"Seaman/service"
	"Seaman/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strconv"
	"time"
)

/**
 * 项目控制器结构体：用来实现处理项目模块的接口的请求，并返回给客户端
 */
type ProjectController struct {
	//上下文对象
	Ctx iris.Context
	//project service
	ProjectService service.ProjectService
	//session对象
	Session *sessions.Session
}

func (uc *ProjectController) BeforeActivation(a mvc.BeforeActivation) {

	//添加项目
	a.Handle("POST", "/", "PostAddProject")

	//删除项目
	a.Handle("DELETE", "/{id}", "DeleteProject")

	//查询项目
	a.Handle("GET", "single/{id}", "GetProject")

	//查询项目数量
	a.Handle("GET", "/count", "GetCount")

	//所有项目列表
	a.Handle("GET", "/list", "GetList")

	//带参数查询项目分页
	a.Handle("GET", "/pageList", "GetPageList")

	//修改项目
	a.Handle("PUT", "/", "UpdateProject")
}

/**
 * url: /project/addproject
 * type：post
 * descs：添加项目
 */
func (uc *ProjectController) PostAddProject() mvc.Result {

	var project model.SmProjectT
	err := uc.Ctx.ReadJSON(&project)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}

	currentUserId, sessionErr := uc.Session.GetInt64(CURRENTUSERID)

	//解析失败
	if sessionErr != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.EEROR_UNLOGIN,
				"message": utils.Recode2Text(utils.EEROR_UNLOGIN),
			},
		}
	}

	newProject := &model.SmProjectT{
		ProjectCode:      project.ProjectCode,
		ProjectName:      project.ProjectName,
		ProjectDesc:      project.ProjectDesc,
		Status:           project.Status,
		TenantId:         project.TenantId,
		AppName:          project.AppName,
		AppScope:         project.AppScope,
		CreateDate:       time.Now(),
		LastUpdateDate:   time.Now(),
		CreateUserId:     strconv.Itoa(int(currentUserId)),
		LastUpdateUserId: strconv.Itoa(int(currentUserId)),
		Revision:         1,
	}
	isSuccess := uc.ProjectService.AddProject(newProject)
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
 * 删除项目
 */
func (uc *ProjectController) DeleteProject() mvc.Result {

	id := uc.Ctx.Params().Get("id")

	projectId, err := strconv.Atoi(id)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}
	delete := uc.ProjectService.DeleteProject(projectId)
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
 * 获取项目
 * 请求类型：Get
 */
func (uc *ProjectController) GetProject() mvc.Result {

	id := uc.Ctx.Params().Get("id")

	projectId, err := strconv.Atoi(id)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}

	project, err := uc.ProjectService.GetProject(projectId)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//返回项目
	return mvc.Response{
		Object: project.ProjectTToRespDesc(),
	}
}

/**
 * 获取项目数量
 * 请求类型：Get
 */
func (uc *ProjectController) GetCount() mvc.Result {

	//获取页面参数
	name := uc.Ctx.FormValue("name")
	projectParam := &model.SmProjectT{
		ProjectName: name,
	}
	//项目总数
	total, err := uc.ProjectService.GetProjectTotalCount(projectParam)

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
 * 获取项目列表
 * 请求类型：Get
 */
func (uc *ProjectController) GetList() mvc.Result {

	var project model.SmProjectT
	err := uc.Ctx.ReadJSON(&project)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_UPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_UPDATE),
			},
		}
	}
	newProject := &model.SmProjectT{
		Id:             project.Id,
		ProjectName:    project.ProjectName,
		ProjectCode:    project.ProjectCode,
		ProjectDesc:    project.ProjectDesc,
		LastUpdateDate: time.Now(),
	}
	projectList := uc.ProjectService.GetProjectList(newProject)
	if len(projectList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//将查询到的项目数据进行转换成前端需要的内容
	var respList []interface{}
	for _, project := range projectList {
		respList = append(respList, project.ProjectTToRespDesc())
	}

	//返回项目列表
	return mvc.Response{
		Object: &respList,
	}
}

/**
 * 获取项目带参数分页查询
 * 请求类型：Get
 */
func (uc *ProjectController) GetPageList() mvc.Result {

	offsetStr := uc.Ctx.FormValue("current")
	limitStr := uc.Ctx.FormValue("pageSize")
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
	projectName := uc.Ctx.FormValue("projectName")
	projectCode := uc.Ctx.FormValue("projectCode")
	status := uc.Ctx.FormValue("status")
	projectParam := &model.SmProjectT{
		ProjectName: projectName,
		ProjectCode: projectCode,
		Status:      status,
	}
	projectList := uc.ProjectService.GetProjectPageList(projectParam, offset, limit)
	total, _ := uc.ProjectService.GetProjectTotalCount(projectParam)
	if len(projectList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//将查询到的项目数据进行转换成前端需要的内容
	var respList []interface{}
	for _, project := range projectList {
		respList = append(respList, project.ProjectTToRespDesc())
	}

	//返回项目列表
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
 * descs：修改项目
 */
func (uc *ProjectController) UpdateProject() mvc.Result {

	var project model.SmProjectT
	err := uc.Ctx.ReadJSON(&project)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ADD),
			},
		}
	}

	currentUserId, sessionErr := uc.Session.GetInt64(CURRENTUSERID)

	//解析失败
	if sessionErr != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.EEROR_UNLOGIN,
				"message": utils.Recode2Text(utils.EEROR_UNLOGIN),
			},
		}
	}

	newProject := &model.SmProjectT{
		Id:             project.Id,
		ProjectName:    project.ProjectName,
		ProjectCode:    project.ProjectCode,
		ProjectDesc:    project.ProjectDesc,
		Status:         project.Status,
		LastUpdateDate: time.Now(),
		LastUpdateUserId: strconv.Itoa(int(currentUserId)),
	}
	isSuccess := uc.ProjectService.UpdateProject(newProject)
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

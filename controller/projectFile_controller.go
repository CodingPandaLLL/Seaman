package controller

import (
	"Seaman/model"
	"Seaman/service"
	"Seaman/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"path/filepath"
	"strconv"
)

/**
 * 项目文件控制器结构体：用来实现处理项目文件模块的接口的请求，并返回给客户端
 */
type ProjectFileController struct {
	//上下文对象
	Ctx iris.Context
	//projectFile service
	ProjectFileService service.ProjectFileService
	//session对象
	Session *sessions.Session
}

func (uc *ProjectFileController) BeforeActivation(a mvc.BeforeActivation) {

	//添加项目文件
	a.Handle("POST", "/", "PostAddProjectFile")

	//删除项目文件
	a.Handle("DELETE", "/{id}", "DeleteProjectFile")

	//查询项目文件
	a.Handle("GET", "single/{id}", "GetProjectFile")

	//查询项目文件数量
	a.Handle("GET", "/count", "GetCount")

	//所有项目文件列表
	a.Handle("GET", "/list", "GetList")

	//带参数查询项目文件分页
	a.Handle("GET", "/pageList", "GetPageList")

	//修改项目文件
	a.Handle("PUT", "/", "UpdateProjectFile")

	//浏览项目文件图片文件
	a.Handle("GET", "image/{id}", "GetImage")

	//下载项目文件
	a.Handle("GET", "file/{id}", "GetFile")
}

/**
 * url: /projectFile/addprojectFile
 * type：post
 * desc：添加项目文件
 */
func (uc *ProjectFileController) PostAddProjectFile() mvc.Result {

	var projectFile model.SmProjectFileT
	err := uc.Ctx.ReadJSON(&projectFile)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}


	newProjectFile := &model.SmProjectFileT{

	}
	isSuccess := uc.ProjectFileService.AddProjectFile(newProjectFile)
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
 * 删除项目文件
 */
func (uc *ProjectFileController) DeleteProjectFile() mvc.Result {

	id := uc.Ctx.Params().Get("id")

	projectFileId, err := strconv.Atoi(id)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}
	delete := uc.ProjectFileService.DeleteProjectFile(projectFileId)
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
 * 获取项目文件
 * 请求类型：Get
 */
func (uc *ProjectFileController) GetProjectFile() mvc.Result {

	id := uc.Ctx.Params().Get("id")

	projectFileId, err := strconv.Atoi(id)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}

	projectFile, err := uc.ProjectFileService.GetProjectFile(projectFileId)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//返回项目文件
	return mvc.Response{
		Object: projectFile.FileTToRespDesc(),
	}
}

/**
 * 获取项目文件数量
 * 请求类型：Get
 */
func (uc *ProjectFileController) GetCount() mvc.Result {

	//获取页面参数
	//name := uc.Ctx.FormValue("name")
	projectFileParam := &model.SmProjectFileT{
	}
	//项目文件总数
	total, err := uc.ProjectFileService.GetProjectFileTotalCount(projectFileParam)

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
 * 获取项目文件列表
 * 请求类型：Get
 */
func (uc *ProjectFileController) GetList() mvc.Result {

	projectId := uc.Ctx.FormValue("id")
	IProjectId,err := strconv.Atoi(projectId)
	if err!=nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}
	newProjectFile := &model.SmProjectFileT{
		ProjectId:int64(IProjectId),
	}
	projectFileList := uc.ProjectFileService.GetProjectFileList(newProjectFile)
	if len(projectFileList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//将查询到的项目文件数据进行转换成前端需要的内容
	var respList []interface{}
	for _, projectFile := range projectFileList {
		respList = append(respList, projectFile.FileTToRespDesc())
	}

	//返回项目文件列表
	return mvc.Response{
		Object: &respList,
	}
}

/**
 * 获取项目文件带参数分页查询
 * 请求类型：Get
 */
func (uc *ProjectFileController) GetPageList() mvc.Result {

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
	//projectFileName := uc.Ctx.FormValue("projectFileName")
	//projectFileCode := uc.Ctx.FormValue("projectFileCode")
	//status := uc.Ctx.FormValue("status")
	projectFileParam := &model.SmProjectFileT{

	}
	projectFileList := uc.ProjectFileService.GetProjectFilePageList(projectFileParam, offset, limit)
	total, _ := uc.ProjectFileService.GetProjectFileTotalCount(projectFileParam)
	if len(projectFileList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//将查询到的项目文件数据进行转换成前端需要的内容
	var respList []interface{}
	for _, projectFile := range projectFileList {
		respList = append(respList, projectFile.ProjectFileTToRespDesc())
	}

	//返回项目文件列表
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
 * descs：修改项目文件
 */
func (uc *ProjectFileController) UpdateProjectFile() mvc.Result {

	var projectFile model.SmProjectFileT
	err := uc.Ctx.ReadJSON(&projectFile)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ADD),
			},
		}
	}

	newProjectFile := &model.SmProjectFileT{

	}
	isSuccess := uc.ProjectFileService.UpdateProjectFile(newProjectFile)
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
 * 获取项目文件图片
 * 请求类型：Get
 */
func (uc *ProjectFileController) GetImage(ctx context.Context) {

	id := uc.Ctx.Params().Get("id")

	projectFileId, err := strconv.Atoi(id)

	if err != nil {
		iris.New().Logger().Info(err)
	}

	projectFile, err := uc.ProjectFileService.GetProjectFile(projectFileId)
	if err != nil {
		iris.New().Logger().Info(err)
	}
	//拼接路径
	path := filepath.Join(projectFile.FilePath, projectFile.FileName)

	//下载图片
	//ctx.SendFile("C:\\Users\\lllzj\\Desktop\\pic\\2.jpg","2.jpg")

	//浏览图片
	ctx.ServeFile(path,false)
}

/**
 * 获取项目文件图片
 * 请求类型：Get
 */
func (uc *ProjectFileController) GetFile(ctx context.Context) {

	id := uc.Ctx.Params().Get("id")

	projectFileId, err := strconv.Atoi(id)

	if err != nil {
		iris.New().Logger().Info(err)
	}

	projectFile, err := uc.ProjectFileService.GetProjectFile(projectFileId)
	if err != nil {
		iris.New().Logger().Info(err)
	}
	//拼接路径
	path := filepath.Join(projectFile.FilePath, projectFile.FileName)

	//下载图片
	ctx.SendFile(path,projectFile.FileName)


}

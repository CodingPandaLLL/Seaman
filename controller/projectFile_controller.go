package controller

import (
	"Seaman/config"
	"Seaman/model"
	"Seaman/service"
	"Seaman/utils"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"io"
	"os"
	"strconv"
	"strings"
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

func (pc *ProjectFileController) BeforeActivation(a mvc.BeforeActivation) {

	//添加项目文件
	a.Handle("POST", "/{projectId}", "PostAddProjectFile")

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
 * desc：上传项目文件
 */
func (pc *ProjectFileController) PostAddProjectFile() mvc.Result {
	//获取项目ID
	projectId, err := pc.Ctx.Params().GetInt64("projectId")
	if err != nil {
		iris.New().Logger().Info(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			},
		}
	}
	//生成uuid作为batchNo
	batchNo := uuid.Must(uuid.NewV4()).String()
	iris.New().Logger().Info(projectId)
	//获取文件
	file, info, err := pc.Ctx.FormFile("file")
	if err != nil {
		iris.New().Logger().Info(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			},
		}
	}
	defer file.Close()
	realName := info.Filename
	//获取文件的AttachType
	fns := strings.Split(realName, ".")
	attachType := fns[len(fns)-1]

	//拼接新的字符串名称
	saveName := strings.Replace(uuid.Must(uuid.NewV4()).String(), "-", "", -1 )
	saveName = saveName+"."+attachType

	//保存文件至服务器
	out, err := os.OpenFile("./uploads/"+saveName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		iris.New().Logger().Info(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			},
		}
	}
	iris.New().Logger().Info("文件路径：" + out.Name())
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			},
		}
	}
	//从session中获取信息
	initConfig := config.InitConfig()
	currentUserId,errSession:=pc.Session.GetInt64(initConfig.Session.CurrentUserId)
	//session为空
	if errSession!=nil {
		iris.New().Logger().Info("未登录....")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.EEROR_UNLOGIN,
				"message": utils.Recode2Text(utils.EEROR_UNLOGIN),
			},
		}
	}

	//获取项目所在文件夹
	//todo:此处有待优化，更改为服务器地址，将文件地址改为配置化
	appPath := config.GetAppPath()

	//保存文件信息到数据库
	tplFile := &model.TplFileT{
		BatchNo:    batchNo,
		FilePath:   appPath+"/uploads/" + saveName,
		FileName:   info.Filename,
		FileSize:   info.Size,
		FileUid:    uuid.Must(uuid.NewV4()).String(),
		AttachType: attachType,
		CreateUserId: strconv.Itoa(int(currentUserId)),
		LastUpdateUserId: strconv.Itoa(int(currentUserId)),
	}
	isSuccess := pc.ProjectFileService.AddFile(tplFile)
	if !isSuccess {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ADD),
			},
		}
	}

	//保存文件关联信息到项目文件关联表
	newProjectFile := &model.SmProjectFileT{
		ProjectId: projectId,
		FileId:    tplFile.Id,
	}
	isSuccess = pc.ProjectFileService.AddProjectFile(newProjectFile)
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
func (pc *ProjectFileController) DeleteProjectFile() mvc.Result {

	id := pc.Ctx.Params().Get("id")

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
	delete := pc.ProjectFileService.DeleteProjectFile(projectFileId)
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
func (pc *ProjectFileController) GetProjectFile() mvc.Result {

	id := pc.Ctx.Params().Get("id")

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

	projectFile, err := pc.ProjectFileService.GetProjectFile(projectFileId)
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
func (pc *ProjectFileController) GetCount() mvc.Result {

	//获取页面参数
	//name := pc.Ctx.FormValue("name")
	projectFileParam := &model.SmProjectFileT{
	}
	//项目文件总数
	total, err := pc.ProjectFileService.GetProjectFileTotalCount(projectFileParam)

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
func (pc *ProjectFileController) GetList() mvc.Result {

	projectId := pc.Ctx.FormValue("id")
	IProjectId, err := strconv.Atoi(projectId)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}
	newProjectFile := &model.SmProjectFileT{
		ProjectId: int64(IProjectId),
	}
	projectFileList := pc.ProjectFileService.GetProjectFileList(newProjectFile)
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
func (pc *ProjectFileController) GetPageList() mvc.Result {

	offsetStr := pc.Ctx.FormValue("current")
	limitStr := pc.Ctx.FormValue("pageSize")
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
	//projectFileName := pc.Ctx.FormValue("projectFileName")
	//projectFileCode := pc.Ctx.FormValue("projectFileCode")
	//status := pc.Ctx.FormValue("status")
	projectFileParam := &model.SmProjectFileT{

	}
	projectFileList := pc.ProjectFileService.GetProjectFilePageList(projectFileParam, offset, limit)
	total, _ := pc.ProjectFileService.GetProjectFileTotalCount(projectFileParam)
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
func (pc *ProjectFileController) UpdateProjectFile() mvc.Result {

	var projectFile model.SmProjectFileT
	err := pc.Ctx.ReadJSON(&projectFile)

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
	isSuccess := pc.ProjectFileService.UpdateProjectFile(newProjectFile)
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
func (pc *ProjectFileController) GetImage(ctx context.Context) {

	id := pc.Ctx.Params().Get("id")

	projectFileId, err := strconv.Atoi(id)

	if err != nil {
		iris.New().Logger().Info(err)
	}

	projectFile, err := pc.ProjectFileService.GetProjectFile(projectFileId)
	if err != nil {
		iris.New().Logger().Info(err)
	}
	////拼接路径
	//path := filepath.Join(projectFile.FilePath, projectFile.FileName)

	//浏览图片
	ctx.ServeFile(projectFile.FilePath, false)
}

/**
 * 获取项目文件图片
 * 请求类型：Get
 */
func (pc *ProjectFileController) GetFile(ctx context.Context) {

	id := pc.Ctx.Params().Get("id")

	projectFileId, err := strconv.Atoi(id)

	if err != nil {
		iris.New().Logger().Info(err)
	}

	projectFile, err := pc.ProjectFileService.GetProjectFile(projectFileId)
	if err != nil {
		iris.New().Logger().Info(err)
	}
	////拼接路径
	//path := filepath.Join(projectFile.FilePath, projectFile.FileName)

	//下载图片
	ctx.SendFile(projectFile.FilePath, projectFile.FileName)

}

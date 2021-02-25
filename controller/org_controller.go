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
 * 组织机构控制器结构体：用来实现处理组织机构模块的接口的请求，并返回给客户端
 */
type OrgController struct {
	//上下文对象
	Ctx iris.Context

	OrgService service.OrgService
	//session对象
	Session *sessions.Session
}

func (oc *OrgController) BeforeActivation(a mvc.BeforeActivation) {

	//添加组织机构
	a.Handle("POST", "/", "PostAddOrg")

	//删除组织机构
	a.Handle("DELETE", "/{id}", "DeleteOrg")

	//查询组织机构
	a.Handle("GET", "single/{id}", "GetOrg")

	//查询组织机构数量
	a.Handle("GET", "/count", "GetCount")

	//所有组织机构列表
	a.Handle("GET", "/list", "GetList")

	//带参数查询组织机构分页
	a.Handle("GET", "/pageList", "GetPageList")

	//修改组织机构
	a.Handle("PUT", "/", "UpdateOrg")
}

/**
 * url: /org/addorg
 * type：post
 * descs：添加组织机构
 */
func (oc *OrgController) PostAddOrg() mvc.Result {

	var org model.TplOrgT
	err := oc.Ctx.ReadJSON(&org)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}
	newOrg := &model.TplOrgT{
		PId:     org.PId,
		OrgName: org.OrgName,
		OrgCode: org.OrgCode,
		OrgNote: org.OrgNote,
		Filed1:  org.Filed1,
		Filed2:  org.Filed2,
		Filed3:  org.Filed3,
	}
	isSuccess := oc.OrgService.AddOrg(newOrg)
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
 * 删除组织机构
 */
func (oc *OrgController) DeleteOrg() mvc.Result {

	id := oc.Ctx.Params().Get("id")

	orgId, err := strconv.Atoi(id)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}
	delete := oc.OrgService.DeleteOrg(orgId)
	delete = oc.OrgService.DeleteOrgChildren(orgId)
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
 * 获取组织机构
 * 请求类型：Get
 */
func (oc *OrgController) GetOrg() mvc.Result {

	id := oc.Ctx.Params().Get("id")

	orgId, err := strconv.Atoi(id)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PARAM,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PARAM),
			},
		}
	}

	org, err := oc.OrgService.GetOrg(orgId)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//返回组织机构
	return mvc.Response{
		Object: org.OrgTToRespDesc(),
	}
}

/**
 * 获取组织机构数量
 * 请求类型：Get
 */
func (oc *OrgController) GetCount() mvc.Result {

	//组织机构总数
	total, err := oc.OrgService.GetOrgTotalCount()

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
 * 获取组织机构列表
 * 请求类型：Get
 */
func (oc *OrgController) GetList() mvc.Result {

	orgList := oc.OrgService.GetOrgList()
	if len(orgList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}
	//将查询到的组织机构数据进行转换成前端需要的内容
	var respList []interface{}
	for _, org := range orgList {
		respList = append(respList, org.OrgTToRespDesc())
	}

	//返回组织机构列表
	return mvc.Response{
		Object: &respList,
	}
}

/**
 * 获取组织机构带参数分页查询
 * 请求类型：Get
 */
func (oc *OrgController) GetPageList() mvc.Result {

	offsetStr := oc.Ctx.FormValue("current")
	limitStr := oc.Ctx.FormValue("pageSize")
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
	pIdStr := oc.Ctx.FormValue("pId")
	pId, err := strconv.ParseInt(pIdStr, 10, 64)
	orgParam := &model.TplOrgT{
		PId: pId,
	}
	orgList := oc.OrgService.GetOrgPageList(orgParam, offset, limit)
	total, _ := oc.OrgService.GetOrgTotalCount()
	if len(orgList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_QUERY,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_QUERY),
			},
		}
	}

	//将查询到的组织机构数据进行转换成前端需要的内容
	var respList []interface{}
	for _, org := range orgList {
		respList = append(respList, org.OrgTToRespDesc())
	}

	//返回组织机构列表
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
 * descs：修改组织机构
 */
func (oc *OrgController) UpdateOrg() mvc.Result {

	var org model.TplOrgT
	err := oc.Ctx.ReadJSON(&org)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ADD),
			},
		}
	}
	newOrg := &model.TplOrgT{
		Id:      org.Id,
		PId:     org.PId,
		OrgName: org.OrgName,
		OrgCode: org.OrgCode,
		OrgNote: org.OrgNote,
		Filed1:  org.Filed1,
		Filed2:  org.Filed2,
		Filed3:  org.Filed3,
	}
	isSuccess := oc.OrgService.UpdateOrg(newOrg)
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

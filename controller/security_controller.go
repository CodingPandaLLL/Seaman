package controller

import (
	"Seaman/model"
	"Seaman/service"
	"Seaman/utils"
	"crypto/md5"
	"encoding/hex"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strconv"
)

/**
 * 管理员控制器
 */
type SecurityController struct {
	//iris框架自动为每个请求都绑定上下文对象
	Ctx iris.Context

	//安全功能实体
	Service service.SecurityService

	//session对象
	Session *sessions.Session
}

const (
	ADMINTABLENAME = "admin"
	ADMIN          = "adminId"
)

type AdminLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/**
 * 管理员退出功能
 * 请求类型：Get
 * 请求url：admin/singout
 */
func (ac *SecurityController) GetSingout() mvc.Result {

	//删除session，下次需要从新登录
	ac.Session.Delete(ADMIN)
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"success": utils.Recode2Text(utils.RESPMSG_SIGNOUT),
		},
	}
}

/**
 * 处理获取管理员总数的路由请求
 * 请求类型：Get
 * 请求Url：admin/count
 */
func (ac *SecurityController) GetCount() mvc.Result {

	count, err := ac.Service.GetAdminCount()

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"message": utils.Recode2Text(utils.RESPMSG_ERRORADMINCOUNT),
				"count":   0,
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  count,
		},
	}
}

/**
 * 获取管理员信息接口
 * 请求类型：Get
 * 请求url：/admin/info
 */
func (ac *SecurityController) GetInfo() mvc.Result {

	//从session中获取信息
	userByte := ac.Session.Get(ADMIN)

	//session为空
	if userByte == nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.EEROR_UNLOGIN,
				"message": utils.Recode2Text(utils.EEROR_UNLOGIN),
			},
		}
	}

	//解析数据到admin数据结构
	//var admin = new(model.Admin)
	//iris.New().Logger().Info(userByte)
	//jsonStr := "" + userByte.(string) + ""
	//admin = model.Decoder([]byte(jsonStr))

	adminId, err := ac.Session.GetInt64(ADMIN)

	//解析失败
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.EEROR_UNLOGIN,
				"message": utils.Recode2Text(utils.EEROR_UNLOGIN),
			},
		}
	}

	userObject, exit := ac.Service.GetByAdminId(adminId)

	if !exit {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  0,
				"success": "登录失败",
				"message": "用户名或者密码错误,请重新登录",
			},
		}
	}

	//解析成功
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"data":   userObject.UserToRespDesc(),
		},
	}
}

/**
 * 管理员登录功能
 * 接口：/security/login
 */
func (ac *SecurityController) PostLogin(context iris.Context) mvc.Result {

	iris.New().Logger().Info(" security login ")

	var adminLogin AdminLogin
	ac.Ctx.ReadJSON(&adminLogin)

	//数据参数检验
	if adminLogin.Username == "" || adminLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"loginStatus": 0,
				"success":     "登录失败",
				"message":     "用户名或密码为空,请重新填写后尝试登录",
			},
		}
	}

	//对密码进行加密
	pwd := adminLogin.Password
	m := md5.New()
	m.Write([]byte(pwd))
	pwd = hex.EncodeToString(m.Sum(nil))
	//根据用户名、密码到数据库中查询对应的管理信息
	user, exist := ac.Service.GetByAdminNameAndPassword(adminLogin.Username, pwd)

	//管理员不存在
	if !exist {
		return mvc.Response{
			Object: map[string]interface{}{
				"loginStatus": 0,
				"success":     "登录失败",
				"message":     "用户名或者密码错误,请重新登录",
			},
		}
	}

	//管理员存在 设置session
	//userByte := admin.Encoder()
	ac.Session.Set(ADMIN, user.Id)

	resultUser := new(model.LoginUser)
	resultUser.UserId = user.Id
	return mvc.Response{
		Object: map[string]interface{}{
			"loginStatus": 1,
			"success":     "登录成功",
			"message":     "登录成功",
			"user":        resultUser.LoginUserToRespDesc(),
		},
	}
}

/**
 * 获取所有的管理员列表
 * url：
 */
func (ac *SecurityController) GetAll() mvc.Result {
	iris.New().Logger().Info(" 获取所有管理员列表 ")

	offsetStr := ac.Ctx.FormValue("offset")
	limitStr := ac.Ctx.FormValue("limit")
	var offset int
	var limit int

	//判断offset和limit两个变量任意一个都不能为""
	if offsetStr == "" || limitStr == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERLIST,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USERLIST),
			},
		}
	}

	offset, err := strconv.Atoi(offsetStr)
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERLIST,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USERLIST),
			},
		}
	}

	//做页数的限制检查
	if offset <= 0 {
		offset = 0
	}

	//做最大的限制
	if limit > MaxLimit {
		limit = MaxLimit
	}

	adminList := ac.Service.GetAdminList(offset, limit)
	if len(adminList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERLIST,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USERLIST),
			},
		}
	}
	var respList []interface{}
	for _, admin := range adminList {
		respList = append(respList, admin.UserToRespDesc())
	}
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"data":   respList,
		},
	}
}

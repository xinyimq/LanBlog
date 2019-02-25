package controllers

import (
	"github.com/astaxie/beego"
	"MyBlog/models"
	"encoding/json"
)

type AdminController struct {
	beego.Controller
}

func (u *AdminController) URLMapping() {
	u.Mapping("Login", u.Login)
	u.Mapping("Options", u.Options)
}

// @Title AdminLogin
// @Description admin login
// @Param	body		body 	models.Admin	true		"body for Admin content"
// @Success 200
// @Failure 403 body is empty
// @router /login [post]
func (u *AdminController) Login() {
	result := &models.LoginResult{}
	var admin models.Admin
	json.Unmarshal(u.Ctx.Input.RequestBody, &admin)
	beego.BeeLogger.Warning("%+v\n", admin)
	beego.Info(admin.Name)
	beego.Info(admin.Password)
	result.Code = 0
	result.Msg = "OK"
	result.Data.User = admin
	result.Data.Token = models.GenToken()
	result.Data.Name="管理员"
	result.Data.UUId="admin-uuid"
	u.Data["json"] = result
	u.ServeJSON()
}

// @Title BeforeAdminLogin
// @Description cros before admin login
// @router /login [options]
func (u *AdminController) Options() {
	//u.Data["json"] = "test"
	//u.ServeJSON()
}

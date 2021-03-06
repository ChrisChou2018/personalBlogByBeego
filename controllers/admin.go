package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// AdminController 后台管理
type AdminController struct {
	beego.Controller
}

// IsLogin 确认登录结构体
func (c *AdminController) IsLogin() bool {
	islogin := c.GetSession("isLogin")
	fmt.Println("sessionValue", islogin)
	if islogin != nil && islogin == true {
		return true
	}
	return false
}

// Get 后台管理页面访问访问
func (c *AdminController) Get() {
	dataType := c.Input().Get("datatype")
	hasdatatype := false
	if dataType != "" {
		c.Data["datatype"] = dataType
		hasdatatype = true
	}
	c.Data["hasdatatype"] = hasdatatype
	if isLogin := c.IsLogin(); !isLogin {
		signinURL := c.URLFor("AccountController.Get")
		c.Redirect(signinURL, 302)
	}
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "blog/admin.html"
}

package controllers

import (
	"zxing2004/gopub/src/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TestApiController struct {
	beego.Controller
}

func (c *TestApiController) Get() {
	var projects []models.Project
	o := orm.NewOrm()
	o.Raw("SELECT * FROM `project`  WHERE 1=1").QueryRows(&projects)
	c.Data["json"] = projects
	c.ServeJSON()

}

package usercontrollers

import (
	"zxing2004/gopub/src/controllers"

	"github.com/astaxie/beego/orm"
)

type UserController struct {
	controllers.BaseController
}

func (c *UserController) Get() {
	userId, _ := c.GetInt("id")
	if userId == 0 {
		o := orm.NewOrm()
		var users []orm.Params
		o.Raw("SELECT * FROM `user` ").Values(&users)
		c.SetJson(0, users, "")
		return
	} else {
		o := orm.NewOrm()
		var users []orm.Params
		var res orm.Params
		i, err := o.Raw("SELECT * FROM `user` where id = ? ", userId).Values(&users)
		if err == nil && i > 0 {
			res = users[0]
		}
		c.SetJson(0, res, "")
		return
	}

}

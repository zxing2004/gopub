package confcontrollers

import (
	"zxing2004/gopub/src/controllers"

	"encoding/json"
	"zxing2004/gopub/src/models"

	"github.com/astaxie/beego"
)

type ConfController struct {
	controllers.BaseController
}

func (c *ConfController) Get() {
	projectId, _ := c.GetInt("projectId", 0)
	project, _ := models.GetProjectById(projectId)
	c.SetJson(0, project, "")
	return

}
func (c *ConfController) Post() {
	//projectId,_:=c.GetInt("projectId",0)
	var project models.Project
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &project)
	err = models.UpdateProjectById(&project)
	beego.Info(err)
	return
}

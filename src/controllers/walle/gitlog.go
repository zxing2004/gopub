package wallecontrollers

import (
	"zxing2004/gopub/src/controllers"
	"zxing2004/gopub/src/library/components"
	"zxing2004/gopub/src/models"
)

type GitlogController struct {
	controllers.BaseController
}

func (c *GitlogController) Get() {
	if c.Project == nil || c.Project.Id == 0 {
		c.SetJson(1, nil, "Parameter error")
		return
	}
	s := components.BaseComponents{}
	s.SetProject(c.Project)
	s.SetTask(&models.Task{Id: -3})
	err := s.GetGitLog()
	if err != nil {
		c.SetJson(1, nil, "获取GitLog错误—"+err.Error())
		return
	} else {
		c.SetJson(0, nil, "")
		return
	}

}

package p2pcontrollers

import (
	"zxing2004/gopub/src/controllers"
	"zxing2004/gopub/src/library/p2p/init_sever"
)

type TaskController struct {
	controllers.BaseController
}

func (c *TaskController) Get() {
	taskId := c.GetString("taskId")
	ss, _ := init_sever.P2pSvc.QueryTaskNoHttp(taskId)
	c.SetJson(0, ss, "")
	return
}

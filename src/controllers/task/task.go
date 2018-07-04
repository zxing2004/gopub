package taskcontrollers

import (
	"zxing2004/gopub/src/controllers"

	"zxing2004/gopub/src/models"
)

type TaskController struct {
	controllers.BaseController
}

func (c *TaskController) Get() {
	taskId, _ := c.GetInt("taskId", 0)
	task, _ := models.GetTaskById(taskId)
	c.SetJson(0, task, "")
	return

}

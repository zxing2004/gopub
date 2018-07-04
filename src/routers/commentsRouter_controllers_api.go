package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:ProjectController"] = append(beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:ProjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:ProjectController"] = append(beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:ProjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:ProjectController"] = append(beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:ProjectController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:ProjectController"] = append(beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:ProjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:ProjectController"] = append(beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:ProjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TaskController"] = append(beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TaskController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TaskController"] = append(beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TaskController"],
		beego.ControllerComments{
			Method: "GetAllAndProName",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TaskController"] = append(beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TaskController"] = append(beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TaskController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TaskController"] = append(beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TaskController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TaskController"] = append(beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TaskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TokenController"] = append(beego.GlobalControllerRouter["zxing2004/gopub/src/controllers/api:TokenController"],
		beego.ControllerComments{
			Method: "IssueToken",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}

package init_sever

import (
	"os"
	"zxing2004/gopub/src/library/p2p/common"
	"zxing2004/gopub/src/library/p2p/server"

	"github.com/astaxie/beego"
)

var P2pSvc *server.Server

func init() {

}
func Start() {
	cfg := common.ReadJson("agent/server.json")
	_, err := common.ParserConfig(&cfg)
	cfg.Server = true
	P2pSvc, err = server.NewServer(&cfg)
	if err != nil {
		beego.Error("start server error, %s.\n", err.Error())
		if beego.BConfig.RunMode != "docker" {
			os.Exit(4)
		}
	}
	beego.Info("服务端p2p配置检测成功")
	if err := P2pSvc.Start(); err != nil {
		beego.Error("Start service failed, %s.\n", err.Error())
		if beego.BConfig.RunMode != "docker" {
			os.Exit(4)
		}
	}
}

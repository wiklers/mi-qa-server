package main

import (
	"mi-qa-server/protocol/MiQaServer"
	"mi-qa-server/servant"

	"github.com/TarsCloud/TarsGo/tars"
)

//Init servant
func main() {
	miQaServerImp := new(servant.MiQaServerImp)
	miQaServerObj := new(MiQaServer.MiQaServerObj)
	cfg := tars.GetServerConfig()
	miQaServerObj.AddServant(miQaServerImp, cfg.App+"."+cfg.Server+".MiQaServerObj")
	tars.Run()
	greeting := "how r u"
	miQaServerObj.EchoHello("naye", &greeting)
}

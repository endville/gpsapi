package main

import (
	_ "gpsapi/docs"
	_ "gpsapi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"log"

	rpc "gpsapi/rpcClient"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/document"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
	})

	var testReply int
	if err := rpc.Test(rpc.RPC_SIGNAL_TEST, &testReply); err != nil {
		log.Println(err.Error())
	} else {
		if testReply != rpc.RPC_SIGNAL_TEST_REPLY {
			log.Println("RPC service has error.")
		}
	}

	beego.Run()
}

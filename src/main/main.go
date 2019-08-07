//author: richard
// @title Restful API demo
// @version 1.1
// @description 实践Restful API
// @contact.name richard
// @contact.email cugriver@163.com
// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
// @host 192.168.1.101:8080
// @BasePath /v1
// @schemes http https
package main

import (
	"args"
	"fmt"
	"logs"
	"pool"
	"router"
)

func main() {
	var err error
	//init args
	err = args.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//init log
	err = logs.InitLog(args.LogConfigure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer logs.Close()
	//init database
	err = pool.InitDatabase(args.DatabaseConfigure())
	if err != nil {
		logs.Error(err.Error())
		return
	}
	//init route
	err = router.NewRouter()
	if err != nil {
		logs.Error(err.Error())
		return
	}
}

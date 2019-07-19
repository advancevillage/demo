//author: richard
// @title Restful API demo
// @version 1.1
// @description 实践Restful API
// @contact.name richard
// @contact.email cugriver@163.com
// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http https
package main

import (
	"args"
	"router"
)

func main() {
	var err error
	//init args
	args.Init()
	//init route
	err = router.NewRouter()
	if err != nil {
		return
	}
}

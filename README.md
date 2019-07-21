# demo

### Swagger API 注释提取[swag init -g main/main.go]
 * 目录结构  main/main.go 
````golang
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
````
 * API 函数
````go
// @Summary Show a account
// @Description get string by ID
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} api.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 401 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [get]
````
### 构建
 * git submodule init
 * git submodule update
 * cd 3rd
   * git submodule init
   * git submodule update


### 参考文档
 * [10 best practices restful api](https://blog.mwaysolutions.com/2014/06/05/10-best-practices-for-better-restful-api/)
 * [Swagger 注释提取](https://github.com/swaggo/swag)
//author: richard
package args

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"model"
	"os"
)

var (
	host 	string
	port 	string
	configure       model.Configure
	configureFile	string
)

func Init() (err error) {
	var args = os.Args
	var length = len(args)
	host = "100.100.20.36"
	port = "8080"
	configureFile = "config/config.xml"
	for i := 1; i < length; i += 2 {
		switch args[i] {
		case "-h":
			if j := i + 1; j < length {
				host = args[j]
			}
		case "-p":
			if j := i + 1; j < length {
				port = args[j]
			}
		case "-c":
			if j := i + 1; j < length {
				configureFile = args[j]
			}
		default:
			continue
		}
	}
	err = initConfigure(configureFile)
	return
}

//@brief: 初始化配置对象
//@param: file  配置文件
//eg:
//  initConfigure("config/config.xml")
func initConfigure(file string) (err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer func(){
		err = f.Close()
	}()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	err = xml.Unmarshal(buf, &configure)
	if err != nil {
		return
	}
	return
}

func HttpServerAddress() string {
	return fmt.Sprintf("%s:%s", host, port)
}

func DatabaseConfigure() *model.Databases {
	return &configure.Databases
}


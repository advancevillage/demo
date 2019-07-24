//author: richard
package args

import (
	"fmt"
	"os"
)

var (
	host 	string
	port 	string
	config	string
)

func Init() {
	var args = os.Args
	var length = len(args)
	host = "127.0.0.1"
	port = "8080"
	config = "config/config.xml"
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
				config = args[j]
			}
		default:
			continue
		}
	}
}

func HttpServerAddress() string {
	return fmt.Sprintf("%s:%s", host, port)
}

func ConfigureFile() string {
	return config
}


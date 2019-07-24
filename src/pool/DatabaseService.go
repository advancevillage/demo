//author: richard
package pool

import (
	"encoding/xml"
	"io/ioutil"
	"model"
	"os"
)

var (
	databases  model.Databases
)

func Init(file string) (err error) {
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
	err = xml.Unmarshal(buf, &databases)
	if err != nil {
		return
	}
	return nil
}


//author: richard
package pool

import (
	"encoding/xml"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"math/rand"
	"model"
	"os"
	"sync"
)

var (
	databases  model.Databases
	databaseOnce sync.Once
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
	//init  database connection
	databaseOnce.Do(func() {
		// init master-database
		master := &databases.Master
		master.Configure.DB, err = gorm.Open(master.Configure.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",master.Configure.User, master.Configure.Password, master.Configure.Host, master.Configure.Port, master.Configure.Schema, master.Configure.Charset))
		if err != nil {
			return
		}
		// init slave-databases
		slaves := &databases.Slaves
		for i := range slaves.Configure {
			slaves.Configure[i].DB, err = gorm.Open(slaves.Configure[i].Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",slaves.Configure[i].User, slaves.Configure[i].Password, slaves.Configure[i].Host, slaves.Configure[i].Port, slaves.Configure[i].Schema, slaves.Configure[i].Charset))
			if err != nil {
				return
			}
		}
	})
	return
}
// @brief: write or read
// @param: b
// @eg:
//   b = true （write)
//   b = false (read)
func Connection(b bool) (DB *gorm.DB) {
	if b {
		DB = databases.Master.Configure.DB
	} else {
		length := len(databases.Slaves.Configure)
		index := rand.Intn(length)
		DB = databases.Slaves.Configure[index].DB
	}
	return
}


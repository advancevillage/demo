//author: richard
package pool

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"model"
	"sync"
	"util"
)

var (
	databases     *model.Databases
	databaseOnce  sync.Once
	caches 		  *model.Caches
	cacheOnce 	  sync.Once
)

//init database
func InitDatabase(o  *model.Databases) (err error) {
	//init  database connection
	databaseOnce.Do(func() {
		databases = o
		// init master-database
		master := &databases.Master
		master.Member.DB, err = gorm.Open(master.Member.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",master.Member.User, master.Member.Password, master.Member.Host, master.Member.Port, master.Member.Schema, master.Member.Charset))
		if err != nil {
			return
		}
		// init slave-databases
		slaves := &databases.Slaves
		for i := range slaves.Members {
			slaves.Members[i].DB, err = gorm.Open(slaves.Members[i].Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",slaves.Members[i].User, slaves.Members[i].Password, slaves.Members[i].Host, slaves.Members[i].Port, slaves.Members[i].Schema, slaves.Members[i].Charset))
			if err != nil {
				return
			}
		}
	})
	return
}
// @brief: 获取数据库连接
// @param: b boolean
// @eg:
//   b = true （write)  获取写连接
//   b = false (read)   获取读连接
func DatabaseConnection(b bool) (DB *gorm.DB) {
	if b {
		DB = databases.Master.Member.DB
	} else {
		length := len(databases.Slaves.Members)
		index := util.RandomInt(length)
		DB = databases.Slaves.Members[index].DB
	}
	return
}

//init cache
func InitCache(o *model.Caches) (err error) {
	cacheOnce.Do(func() {
		caches = o
		//init master
		master := &caches.Master
		master.Member.Client = redis.NewClient(&redis.Options{Addr:fmt.Sprintf("%s:%s", master.Member.Host, master.Member.Port), Password: master.Member.Password})
		_, err = master.Member.Client.Ping().Result()
		if err != nil {
			return
		}
		//init slaves
		slaves := &caches.Slaves
		for i := range slaves.Members {
			slaves.Members[i].Client = redis.NewClient(&redis.Options{Addr:fmt.Sprintf("%s:%s", slaves.Members[i].Host, slaves.Members[i].Port), Password: slaves.Members[i].Password})
			_, err = slaves.Members[i].Client.Ping().Result()
			if err != nil {
				return
			}
		}
	})
	return
}
// @brief: 获取缓存服务连接
// @param: b boolean
// @eg:
//   b = true （write)  获取写连接
//   b = false (read)   获取读连接
func CacheConnection(b bool) (Client *redis.Client) {
	if b {
		Client = caches.Master.Member.Client
	} else {
		length := len(caches.Slaves.Members)
		index := util.RandomInt(length)
		Client = caches.Slaves.Members[index].Client
	}
	return
}
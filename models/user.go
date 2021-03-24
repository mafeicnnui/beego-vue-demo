package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

// Model Struct
type User struct {
	Id       int
	Name     string `orm:"size(100)"`
	Gender   string `orm:"size(3)"`
	Age      int
	Birthday time.Time `orm:"type(date)"`
}

func init() {
	// set default database
	user, _ := web.AppConfig.String("mysqluser")
	password, _ := web.AppConfig.String("mysqlpass")
	url, _ := web.AppConfig.String("mysqlurls")
	db, _ := web.AppConfig.String("mysqldb")
	charset, _ := web.AppConfig.String("charset")
	timezone, _ := web.AppConfig.String("timezone")
	ds := user + ":" + password + "@tcp(" + url + ")/" + db + "?charset=" + charset + "&loc=" + timezone
	orm.RegisterDataBase("default", "mysql", ds)

	// register model
	orm.RegisterModel(new(User))
	// create table
	orm.RunSyncdb("default", false, true)
}

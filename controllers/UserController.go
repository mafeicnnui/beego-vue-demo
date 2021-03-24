package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"student/models"
	"time"
)

type UserController struct {
	BaseController
}

type UserControllerAdd struct {
	BaseController
}

type UserControllerSearch struct {
	BaseController
}
type UserControllerSearch_ struct {
	BaseController
}

//query user
func (c *UserController) Get() {
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	u := models.User{Id: id}
	err := o.Read(&u)
	if err != nil {
		c.ErrorJson(500, err.Error(), nil)

	} else {
		c.SuccessJson(&u)
	}
}

//insert user
func (c *UserController) Put() {
	name := c.GetString("name")
	gender := c.GetString("gender")
	age, _ := c.GetInt("age")
	template := "2006-01-02"
	birthday, err := time.Parse(template, c.GetString("birthday"))
	if err != nil {
		c.ErrorJson(500, err.Error(), nil)
	}
	fmt.Printf("User info:%s,%s,%d,%s", name, gender, age, birthday)
	o := orm.NewOrm()
	user := models.User{Name: name, Gender: gender, Age: age, Birthday: birthday}
	id, err := o.Insert(&user)
	if err != nil {
		c.ErrorJson(500, err.Error(), nil)
	} else {
		c.SuccessJson("id=" + strconv.FormatInt(id, 10))
	}

}

//update user
func (c *UserController) Post() {
	o := orm.NewOrm()
	id, _ := c.GetInt("id")
	user := models.User{Id: id}
	err := o.Read(&user)
	user.Name = c.GetString("name")
	user.Age, _ = c.GetInt("age")
	user.Gender = c.GetString("gender")
	template := "2006-01-02"
	user.Birthday, err = time.Parse(template, c.GetString("birthday"))
	if err != nil {
		c.ErrorJson(500, err.Error(), nil)
	}
	fmt.Println("Update:", &user)
	_, err = o.Update(&user)
	if err != nil {
		c.ErrorJson(500, err.Error(), nil)
	} else {
		c.SuccessJson(nil)
	}

}

//delete user
func (c *UserController) Delete() {
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	user := models.User{Id: id}
	_, err := o.Delete(&user)
	if err != nil {
		c.ErrorJson(500, err.Error(), nil)
	} else {
		c.SuccessJson(nil)
	}
}

//add user page
func (c *UserControllerAdd) Get() {
	c.TplName = "add_user.html"
}

//query user page
func (c *UserControllerSearch) Get() {
	c.TplName = "query_user.html"
}

func (c *UserControllerSearch_) Get() {
	name := c.GetString("name")
	var users []*models.User
	orm := orm.NewOrm()
	qs := orm.QueryTable("user")
	n, err := qs.Filter("name__contains", name).All(&users) // 过滤器
	if err == nil && n > 0 {
		c.SuccessJson(&users)
	} else {
		c.ErrorJson(500, err.Error(), nil)
	}
}

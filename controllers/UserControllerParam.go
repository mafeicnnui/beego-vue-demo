package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"student/models"
)

type UserControllerPar struct {
	BaseController
}

//query user
func (c *UserControllerPar) Get() {
	id := c.Ctx.Input.Param(":id")
	int32, err := strconv.Atoi(id)
	if err != nil {
		c.ErrorJson(500, err.Error(), nil)
	}

	o := orm.NewOrm()
	u := models.User{Id: int32}
	err = o.Read(&u)
	if err != nil {
		c.ErrorJson(500, err.Error(), nil)

	} else {
		c.SuccessJson(&u)
	}
}

//delete user
func (c *UserControllerPar) Delete() {
	id := c.Ctx.Input.Param(":id")
	int32, err := strconv.Atoi(id)
	if err != nil {
		c.ErrorJson(500, err.Error(), nil)
	}

	o := orm.NewOrm()
	user := models.User{Id: int32}
	_, err = o.Delete(&user)
	if err != nil {
		c.ErrorJson(500, err.Error(), nil)
	} else {
		c.SuccessJson(nil)
	}
}

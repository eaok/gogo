package controllers

import (
	"beego-api/models"

	"github.com/astaxie/beego"
)

type LokcolController struct {
	beego.Controller
}

// TestApi ...
// @Title Get
// @Description create IntAuthTokenCache
// @Success 201 {object} LokcolController.Data["json"]
// @Failure 403 body is empty
// @router / [get]
func (c *LokcolController) TestApi() {
	c.Data["json"] = map[string]interface{}{
		"ret": 1000,
	}
	c.ServeJSON()
}

func (c *LokcolController) ArithmeticAction() {
	var a, b int
	c.Ctx.Input.Bind(&a, "a")
	c.Ctx.Input.Bind(&b, "b")
	result := a + b

	version := c.Ctx.Input.Param(":version")
	action := c.Ctx.Input.Param(":action")

	c.Data["json"] = map[string]interface{}{
		"ret":     1000,
		"version": version,
		"action":  action,
		"result":  result,
	}
	c.ServeJSON()
}

func (c *LokcolController) StudentList() {
	l, err := models.GetAllStudent(map[string]string{}, []string{}, []string{}, []string{}, 0, 100)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Data["student_list"] = l
		c.TplName = "templates/student.html"
	}
}

func (c *LokcolController) AuthToken() {
	var intAuthToken string
	c.Ctx.Input.Bind(&intAuthToken, "intAuthToken")
	v, err := models.GetIntAuthTokenCacheById(intAuthToken)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"ret": 1000,
		}
		c.ServeJSON()
	} else {
		c.Ctx.WriteString("Welcome " + v.Device + " user from " + v.Ip + "!")
	}
}

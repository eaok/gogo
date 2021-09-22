package api

import (
	"api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//index
func IndexUsers(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

//增加一条记录
func AddUsers(c *gin.Context) {
	name := c.Request.FormValue("name")
	telephone := c.Request.FormValue("telephone")
	person := models.Person{
		Name:      name,
		Telephone: telephone,
	}
	id := person.Create()
	msg := fmt.Sprintf("insert successful %d", id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

//获得一条记录
func GetOne(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	p := models.Person{
		Id: id,
	}
	rs, _ := p.GetRow()
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//获得所有记录
func GetAll(c *gin.Context) {
	p := models.Person{}
	rs, _ := p.GetRows()
	c.JSON(http.StatusOK, gin.H{
		"list": rs,
	})
}

//更新telephone
func UpdateUser(c *gin.Context) {
	ids := c.Request.FormValue("id")
	id, _ := strconv.Atoi(ids)
	telephone := c.Request.FormValue("telephone")
	person := models.Person{
		Id:        id,
		Telephone: telephone,
	}
	row := person.Update()
	msg := fmt.Sprintf("updated successful %d", row)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

//删除一条记录
func DelUser(c *gin.Context) {
	ids := c.Request.FormValue("id")
	id, _ := strconv.Atoi(ids)
	row := models.Delete(id)
	msg := fmt.Sprintf("delete successful %d", row)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_b"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_a"`
	}
	FieldD string `form:"field_b"`
}

type Login struct {
	User     string `form:"username" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

// 1.绑定 JSON
// {"user": "pan", "password": "123"}
func bindJSON(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.User != "pan" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

// 2.绑定 XML
//<?xml version="1.0" encoding="UTF-8"?>
//<root>
//	<user>user</user>
//	<password>123</password>
//</root>
func bindXML(c *gin.Context) {
	var xml Login
	if err := c.ShouldBindXML(&xml); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if xml.User != "pan" || xml.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

// 3.绑定 form
// username=pan&password=123
func bindForm(c *gin.Context) {
	var form Login

	// 根据 Content-Type Header 推断使用哪个绑定器
	//if err := c.Bind(&form); err != nil {
	//if err := c.BindWith(&form, binding.Form); err != nil {
	//if err := c.ShouldBindWith(&form, binding.Form); err != nil {
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.User != "pan" || form.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

// 4.绑定 URI
// curl -v -X POST localhost:8080/login/pan/987fbc97-4bed-5078-9f07-9141ba07c9f3
func bindUri(c *gin.Context) {
	var person Person
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
}

//绑定到嵌套结构体
func GetData(c *gin.Context) {
	var b StructB
	var cc StructC
	var d StructD

	if err := c.Bind(&b); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	if err := c.Bind(&cc); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	if err := c.Bind(&d); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	c.JSON(200, gin.H{"a": b.NestedStruct, "b": b.FieldB})
	c.JSON(200, gin.H{"a": cc.NestedStructPointer, "c": cc.FieldC})
	c.JSON(200, gin.H{"x": d.NestedAnonyStruct, "d": d.FieldD})
}

func main() {
	router := gin.Default()

	router.POST("/loginJSON", bindJSON)
	router.POST("/loginXML", bindXML)
	router.POST("/loginForm", bindForm)
	router.POST("/login/:name/:id", bindUri)

	router.GET("/get_data", GetData)

	router.Run()
}

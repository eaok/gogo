package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DqmUserRole struct {
	ID        int64     `gorm:"column:id;primary_key" json:"id"`
	UserId    string    `gorm:"column:user_id" json:"user_id"`
	RoleId    string    `gorm:"column:role_id" json:"role_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func affair(db *gorm.DB) {
	var err error
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err = tx.Create(&DqmUserRole{UserId: "8", RoleId: "8"}).Error; err != nil {
		//tx.Rollback()
		//return
	}

	if err = tx.Create(&DqmUserRole{UserId: "9", RoleId: "9"}).Error; err != nil {
		//tx.Rollback()
		//return
	}
}

func main() {
	db, connErr := gorm.Open("mysql", "root:root@/gormdb?charset=utf8mb4&parseTime=True&loc=Local")
	if connErr != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// 如果不设置则gorm默认表名会在后面加s（注意struct的名字要和表名一致，否则会找不到）
	db.SingularTable(true)

	//如果表不存在就创建表
	if db.HasTable(DqmUserRole{}) == false {
		db.CreateTable(&DqmUserRole{})
		userRole1 := &DqmUserRole{
			UserId: "1",
			RoleId: "1,2",
		}
		db.Create(userRole1)
		userRole2 := &DqmUserRole{
			UserId: "2",
			RoleId: "2",
		}
		db.Create(userRole2)
	}

	/**
	初级查询
	*/
	var dqmUserRole DqmUserRole
	db.First(&dqmUserRole) // 按照主键顺序的第一条记录
	fmt.Println("roleId: ", dqmUserRole.RoleId)

	var dqmUserRole1 DqmUserRole
	db.Last(&dqmUserRole1) // 按照主键顺序的最后一条记录
	fmt.Println("roleId: ", dqmUserRole1.RoleId)

	var dqmUserRoels []DqmUserRole
	db.Find(&dqmUserRoels) // 所有记录
	fmt.Println("dqmUserRoles: ", dqmUserRoels)

	var dqmUserRole2 DqmUserRole
	db.First(&dqmUserRole2, 2) // 按照给定主键查找第一条记录
	fmt.Println("roleId: ", dqmUserRole2.RoleId)

	/**
	where查询
	*/
	var dqmUserRole3 DqmUserRole
	// 根据条件查询得到满足条件的第一条记录
	db.Where("role_id = ?", "2").First(&dqmUserRole3)
	fmt.Println("where roleId: ", dqmUserRole3.RoleId)

	var dqmUserRoles4 []DqmUserRole
	// 根据条件查询得到满足条件的所有记录
	db.Where("user_id = ?", "1").Find(&dqmUserRoles4)
	fmt.Println("where dqmUserRoles: ", dqmUserRoles4)

	var dqmUserRole5 []DqmUserRole
	// like模糊查询
	db.Where("role_id like ?", "%2").Find(&dqmUserRole5)
	fmt.Println("where dqmUserRoles: ", dqmUserRole5)

	var dqmUserRole6 []DqmUserRole
	db.Where("updated_at > ?", "2019-02-08 18:08:27").Find(&dqmUserRole6)
	fmt.Println("where dqmUserRoles: ", dqmUserRole6)

	var dqmUserRole7 DqmUserRole
	// struct结构查询条件
	db.Where(&DqmUserRole{RoleId: "1,2", UserId: "1"}).First(&dqmUserRole7)
	fmt.Println("where dqmUserRole: ", dqmUserRole7)

	var dqmUserRole8 DqmUserRole
	// map结构查询条件
	db.Where(map[string]interface{}{"role_id": "1,2", "user_id": "1"}).Find(&dqmUserRole8)
	fmt.Println("where dqmUserRole: ", dqmUserRole8)

	/**
	Not查询
	*/
	var dqmUserRole9 DqmUserRole
	db.Not([]int64{1}).First(&dqmUserRole9)
	fmt.Println("not dqmUserRole: ", dqmUserRole9)

	/**
	Or查询
	*/
	var dqmUserRole10 []DqmUserRole
	db.Where(&DqmUserRole{RoleId: "1,2"}).Or(map[string]interface{}{"user_id": "2"}).Find(&dqmUserRole10)
	fmt.Println("or dqmUserRoles: ", dqmUserRole10)

	/**
	FirstOrInit Attrs
	*/
	var dqmUserRole11 DqmUserRole
	// 查不到该条记录，则使用attrs值替换
	db.Where("user_id = ?", "0").Attrs("role_id", "12").FirstOrInit(&dqmUserRole11)
	fmt.Println("after FirstOrInit: ", dqmUserRole11)

	var dqmUserRole12 DqmUserRole
	// 查到记录，则使用数据库中的值
	db.Where("user_id = ?", "1").Attrs("role_id", "2").FirstOrInit(&dqmUserRole12)
	fmt.Println("after FirstOrInit: ", dqmUserRole12)

	/**
	Assign
	*/
	var dqmUserRole13 DqmUserRole
	// 不管是否找到对应记录，使用Assign值替代查询到的值
	db.Where("role_id = ?", "1,2").Assign(DqmUserRole{UserId: "15"}).FirstOrInit(&dqmUserRole13)
	fmt.Println("assign dqmUserRole: ", dqmUserRole13)

	/**
	FirstOrCreate
	*/
	var dqmUserRole14 DqmUserRole
	// 如果记录存在则返回结果，如果不存在则创建
	db.Where(&DqmUserRole{UserId: "3", RoleId: "3"}).FirstOrCreate(&dqmUserRole14)
	fmt.Println("firstOrCreate dqmUserRole: ", dqmUserRole14)

	/**
	select
	*/
	var dqmUserRole15 []DqmUserRole
	db.Select("id").Find(&dqmUserRole15)
	fmt.Println("select dqmUserRoles: ", dqmUserRole15)

	/**
	order
	*/
	var dqmUserRole16 []DqmUserRole
	db.Order("user_id desc").Find(&dqmUserRole16) // 注意这里的order要在find前面，否则不生效
	fmt.Println("order dqmUserRoles: ", dqmUserRole16)

	/**
	limit
	*/
	var dqmUserRole17 []DqmUserRole
	db.Limit(2).Find(&dqmUserRole17)
	fmt.Println("limit dqmUserRoles: ", dqmUserRole17)

	/**
	offset
	*/
	var dqmUserRole18 []DqmUserRole
	db.Limit(10).Offset(2).Find(&dqmUserRole18) // 如果只有offset没有limit则不会生效
	fmt.Println("offset dqmUserRoles: ", dqmUserRole18)

	/**
	count
	*/
	var dqmUserRole19 []DqmUserRole
	var count int64
	db.Find(&dqmUserRole19).Count(&count) // 如果只有offset没有limit则不会生效
	fmt.Println("count count: ", count)

	/**
	子查询
	*/
	var dqmUserRole20 []DqmUserRole
	rows, _ := db.Where("user_id in (?)", []string{"1", "2"}).Select("id").Find(&dqmUserRole20).Rows()
	var str []string
	var id int64
	for rows.Next() {
		rows.Scan(&id)
		str = append(str, strconv.Itoa(int(id)))
	}

	var dqmUserRole21 []DqmUserRole
	db.Where("id in (?)", str).Find(&dqmUserRole21)
	fmt.Println("subSelect dqmUserRoles: ", dqmUserRole21)

	/**
	将查询结果存入指定结构
	*/
	type Result struct {
		Id int64
	}
	var results []Result
	db.Select("id").Where("user_id in (?)", []string{"1", "2"}).Find(&dqmUserRole20).Scan(&results)
	fmt.Println("ids: ", results)

	/**
	update
	*/
	var dqmUserRole22 DqmUserRole
	db.Model(&dqmUserRole22).Where("user_id = ?", "4").Update("role_id", "4,8")
	fmt.Println("update dqmUserRole: ", dqmUserRole22)

	var dqmUserRole23 DqmUserRole
	db.Model(&dqmUserRole22).Select("role_id").Where("id = ?", 1004).Update(map[string]interface{}{"user_id": "5", "role_id": "5"})
	fmt.Println("update dqmUserRole: ", dqmUserRole23)

	/**
	执行原生sql语句
	*/
	var dqmUserRole24 []DqmUserRole
	db.Exec("select * from dqm_user_role").Find(&dqmUserRole24)
	fmt.Println("sql dqmUserRole: ", dqmUserRole24)

	/**
	事务
	*/
	affair(db)

	/**
	错误处理
	*/
	var dqmUserRole25 DqmUserRole
	err := db.Where("role_id = ?", 54321).First(&dqmUserRole25).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("ErrRecordNotFound, record not found")
	} else {
		fmt.Println("err: ", err)
	}
	fmt.Println("err dqmUserRole: ", dqmUserRole25)
}

package model

import (
	"fmt"
	"main/utils"
)

//User 结构体
type User struct {
	ID       int
	UserName string
	Password string
	age      int
	sex      string
}

//AddUser 添加User 方法一
func (user *User) addUser() error {
	//1.sql语句
	sqlStr := "insert into user_tb(username,password,age,sex) values(?,?,?,?)"
	//2.预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常：", err)
		return err
	}
	//3.执行
	_, err2 := inStmt.Exec("admin6", "123456", 25, "男")

	if err2 != nil {
		fmt.Println("执行出现异常：", err2)
		return err2
	}
	return nil
}

//AddUser1 添加User 方法二
func (user *User) addUser1() error {
	//1.sql语句
	sqlStr := "insert into user_tb(username,password,age,sex) values(?,?,?,?)"

	//3.执行
	_, err := utils.Db.Exec(sqlStr, "admin9", "666666", 25, "男")
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//DeleteUser 删除User
func (user *User) delUser() error {
	sqlStr := "delete from user_tb where id=?"
	_, err := utils.Db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Println("删除失败： ", err)
		return err
	}
	return nil
}

//updateUser 修改User
func (user *User) updateUser() error {
	sqlStr := "update user_tb set username=?,password=?,age=?,sex=? where id=?"
	_, err := utils.Db.Exec(sqlStr, "小明", "999999", 15, "男", 2)
	if err != nil {
		fmt.Println("修改失败：", err)
		return err
	}
	return nil

}

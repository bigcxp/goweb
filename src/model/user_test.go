package model

import (
	"fmt"
	"testing"
)

//在测试之前做点其他的事情
func TestMain(m *testing.M) {
	fmt.Println("测试开始")
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("主测试函数执行：")
	//通过t.Run(String,func)执行
	t.Run("添加用户函数开始执行：", testAddUser)
	t.Run("删除用户函数开始执行：", testDelUser)
	t.Run("修改用户函数开始执行：", testUpdateUser)

}

//如果Test不是大写的，函数不执行TestXxxx ,可以将它设置成一个子测试程序
func testAddUser(t *testing.T) {
	fmt.Println("开始测试添加用户：")
	user := &User{}
	//调用添加用户的方法
	user.addUser()
	user.addUser1()
}

//测试删除子函数
func estDelUser(t *testing.T) {
	fmt.Println("开始删除用户：")
	user := &User{}
	//调用删除用户的方法
	user.delUser()
}

//测试删除子函数
func testUpdateUser(t *testing.T) {
	fmt.Println("开始修改用户：")
	user := &User{}
	//调用删除用户的方法
	user.updateUser()
}

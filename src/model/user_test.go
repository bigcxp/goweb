package model

import (
	"fmt"
	"testing"
)

// func TestDelUser(t *testing.T) {
// 	user := &User{}
// 	user.delUser()
// 	fmt.Println("删除用户成功！")
// }

func TestUpdateUser(t *testing.T) {
	user := &User{}
	user.updateUser()
	fmt.Println("修改用户成功!")
}

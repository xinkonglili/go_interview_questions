/*
提供如下可用函数
// CreateIndex 创建所需索引
func CreateIndex(coll string) error
// OnceDo 全网执行一次
func OnceDo(call func() error) error
// 您需要实现在如下两个包： user, department 中完成 user, dprt两个coll的索引初始化工作。
并且不能在main中显式调用
*/
package main

import "fmt"

type userM interface {
	CreateIndex(coll string) error
}
type user struct {
}

type departmentM interface {
	CreateIndex(coll string) error
}
type department struct {
}

func (u *user) CreateIndex(coll string) error {
	return nil
}

func (depart *department) CreateIndex(coll string) error {
	return nil
}
func OnceDo(call func() error) error {
	return nil
}

// 初始化索引
func initUserIndex() error {
	u := &user{}
	err := u.CreateIndex("user")
	if err != nil {
		return fmt.Errorf("为用户coll创建索引失败: %v", err)
	}
	return nil
}
func initDepartmentIndex() error {
	depart := &department{}
	err := depart.CreateIndex("department")
	if err != nil {
		return fmt.Errorf("为用户coll创建索引失败: %v", err)
	}
	return nil
}
func main() {
	err1 := OnceDo(initUserIndex)
	if err1 != nil {
		fmt.Errorf("为用户coll创建索引失败: %v", err1)
	}
	err2 := OnceDo(initDepartmentIndex)
	if err2 != nil {
		fmt.Errorf("为用户coll创建索引失败: %v", err2)
	}
}

package main

import (
	"fmt"
)

type AliPay struct {
}

type Cash struct {
}

type ContainCanUseFaceID interface {
	CanUseFaceID()
}

type ContainStolen interface {
	Stolen()
}

// 添加具备各自特点的方法
func (receiver AliPay) CanUseFaceID() {

}

func (receiver Cash) Stolen() {

}

func Print(payMethod interface{}) {
	// 用类型分支进行支付方法的特性判断
	switch payMethod.(type) {
	case ContainCanUseFaceID:
		fmt.Printf("ContainCanUseFaceID %T\n", payMethod)
		break
	case ContainStolen:
		fmt.Printf("ContainStolen %T\n", payMethod)
		break
	default:
		fmt.Println("没有匹配上")
	}
}
func main() {
	Print(new(AliPay))
	Print(new(Cash))
}

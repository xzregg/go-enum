package enum

import (
	"fmt"
	"testing"
)

func TestEnum_GetLabel(t *testing.T) {

	ColorEnum := GenerateEnum(&struct {
		Enum   `key:"color" label:"颜色"`
		Red    string `key:"red" label:"红色" choise:"order_status" `
		Yellow string `key:"yellow" label:"黄色"`
		Black  int    `key:"1" label:"黑色"`
		White  int    `key:"2" label:"白色"`
	}{})

	var OrderStatus = GenerateEnum(&struct {
		Enum         `key:"order_status" label:"支付状态"`
		OrderSuccess string `key:"Order_Success" label:"支付成功"`
		OrderPaying  string `key:"Order_Paying" label:"支付中"`
	}{})

	var MajorEnum = GenerateEnum(&struct {
		Enum     `key:"major" label:"识别方式"`
		Fruit    string `key:"fruit" label:"水果识别"`
		Drink    string `key:"drink" label:"饮料识别"`
		Quantity string `key:"quantity" label:"数数量"`
	}{})

	var AttachEnum = GenerateEnum(&struct {
		Enum           `key:"attach" label:"识别类型"`
		Dishes         string `key:"dishes" label:"菜品识别"`
		Bowls          string `key:"bowls" label:"碗碟识别"`
		FruitVegetable string `key:"label" lable:"果蔬识别"`
	}{})

	var RoleType = GenerateEnum(&struct {
		Enum             `key:"role_type" label:"角色类型"`
		UserAdminType    int `key:"1" label:"超管"`
		UserMerchantType int `key:"2" label:"高级用户"`
		UserNormalType   int `key:"3" label:"普通用户"`
	}{})

	fmt.Printf("%v %v\n", ColorEnum.GetEnumName(), ColorEnum.GetLabel("red"))
	fmt.Printf("AllEnumMap[\"color\"] i=%p  ColorEnum.EnumMap i=%p\n", AllEnumMap["color"], ColorEnum.GetEnumMap())
	fmt.Printf("%+v\n", AllEnumMap["color"])
	fmt.Printf("%+v\n", AllEnumMap["color"]["yellow"])
	fmt.Printf("%+v\n", AttachEnum)
	fmt.Printf("%+v\n", MajorEnum)
	fmt.Printf("%+v\n", OrderStatus)
	fmt.Printf("%+v\n", RoleType)
	if ColorEnum.Red != "red" {
		t.Errorf("ColorEnum.Red != red")
	}
	if ColorEnum.White != 2 {
		t.Errorf("ColorEnum.White != 2")
	}

}

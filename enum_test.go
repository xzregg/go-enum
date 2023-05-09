package enum

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnum_GetLabel(t *testing.T) {

	ColorEnum := GenerateEnum(&struct {
		Enum   `key:"color" label:"颜色"`
		Red    string `key:"red" label:"红色" choise:"order_status" `
		Yellow string `key:"yellow" label:"黄色"`
		Black  int    `key:"1" label:"黑色"`
		White  int    `2 白色`
		Blue   int    `3   蓝色`
	}{})

	var RoleType = GenerateEnum(&struct {
		Enum             `key:"role_type" label:"角色类型"`
		UserAdminType    int `key:"1" label:"超管"`
		UserMerchantType int `2  高级用户`
		UserNormalType   int `key:"3" label:"普通用户"`
	}{})

	fmt.Printf("%v %v\n", ColorEnum.GetEnumName(), ColorEnum.GetLabel("red"))
	fmt.Printf("AllEnumMap[\"color\"] i=%p  ColorEnum.EnumMap i=%p\n", AllEnumMap["color"], ColorEnum.GetEnumMap())
	fmt.Printf("%+v\n", AllEnumMap["color"])
	fmt.Printf("%+v\n", AllEnumMap["color"]["yellow"])
	fmt.Printf("%+v\n", RoleType)

	assert.Equal(t, ColorEnum.Red, "red", "ColorEnum.Red != red")
	assert.Equal(t, ColorEnum.White, 2, "ColorEnum.White != 2")

	fmt.Printf("%p\n", &AllEnumMap)

	d, err := json.Marshal(AllEnumMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(d))

}

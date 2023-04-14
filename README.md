go-enum
==============================
go 枚举


Installation
------------

```
go get github.com/xzregg/go-enum
```

Examples
--------

A basic example:

```go
import (
"fmt"
"github.com/xzregg/go-enum"
"testing"
)

func TestEnum_GetLabel(t *testing.T) {

	ColorEnum := enum.GenerateEnum(&struct {
		enum.Enum `key:"color" label:"颜色"`
		Red       string `key:"red" label:"红色" choise:"order_status" `
		Yellow    string `key:"yellow" label:"黄色"`
		Black     int    `key:"1" label:"黑色"`
		White     int    `key:"2" label:"白色"`
	}{})

	fmt.Printf("%v %v\n", ColorEnum.GetEnumName(), ColorEnum.GetLabel("red")) // color 红色
	fmt.Printf("AllEnumMap[\"color\"] i=%p  ColorEnum.EnumMap i=%p\n", enum.AllEnumMap["color"], ColorEnum.GetEnumMap())
	fmt.Printf("%+v\n", enum.AllEnumMap["color"]) // map[1:黑色 2:白色 red:红色 yellow:黄色]
	fmt.Printf("%+v\n", enum.AllEnumMap["color"]["yellow"]) // 黄色

	if ColorEnum.Red != "red" {
		t.Errorf("ColorEnum.Red != red")
	}
	if ColorEnum.White != 2 {
		t.Errorf("ColorEnum.White != 2")
	}

}
```

License
-------

MIT, see [LICENSE](LICENSE)

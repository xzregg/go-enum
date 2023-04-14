go-enum
==============================



Installation
------------



```
go get github.com/xzregg/go-enum
```

Examples
--------

A basic example:

```go
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

	fmt.Printf("%v %v\n", ColorEnum.GetEnumName(), ColorEnum.GetLabel("red"))
	fmt.Printf("AllEnumMap[\"color\"] i=%p  ColorEnum.EnumMap i=%p\n", AllEnumMap["color"], ColorEnum.GetEnumMap())
	fmt.Printf("%+v\n", AllEnumMap["color"])
	fmt.Printf("%+v\n", AllEnumMap["color"]["yellow"])

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

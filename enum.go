package enum

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ItemEnumMap map[any]string

func (i ItemEnumMap) MarshalJSON() ([]byte, error) {
	tmp := map[string]string{}
	for k, v := range i {
		str := fmt.Sprintf("%v", k)
		tmp[str] = v
	}
	return json.Marshal(tmp)
}

type InterFaceEnum interface {
	InitMap(enumName string)
	GetLabel(key any) string
	SetLabel(key any, label string)
	GetEnumMap() ItemEnumMap
	GetEnumName() string
}

type Enum struct {
	EnumMap  ItemEnumMap
	enumName string
}

var AllEnumMap = map[string]ItemEnumMap{}

func (e *Enum) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.EnumMap)
}

func (e *Enum) InitMap(enumName string) {
	if e.EnumMap == nil {
		e.enumName = enumName
		e.EnumMap = ItemEnumMap{}
		AllEnumMap[enumName] = e.EnumMap
		//		fmt.Printf("%s %p i=%p\n", enumName, AllEnumMap[enumName], e.EnumMap)
	}
}

func (e *Enum) GetLabel(key any) string {
	if label, ok := e.EnumMap[key]; ok {
		return label
	}
	return ""
}

func (e *Enum) SetLabel(key any, label string) {
	e.EnumMap[key] = label
}

func (e *Enum) GetEnumMap() ItemEnumMap {
	return e.EnumMap
}
func (e *Enum) GetEnumName() string {
	return e.enumName
}

var valueFunc = make(map[reflect.Kind]func(fieldValue *reflect.Value, tagValue string), 0)

func init() {
	valueFunc[reflect.Bool] = func(fieldValue *reflect.Value, tagValue string) {
		value, _ := strconv.ParseBool(tagValue)
		fieldValue.SetBool(value)
	}

	valueFunc[reflect.Int] = func(fieldValue *reflect.Value, tagValue string) {
		value, _ := strconv.ParseInt(tagValue, 10, 64)
		fieldValue.SetInt(value)
	}
	valueFunc[reflect.Int8] = valueFunc[reflect.Int]
	valueFunc[reflect.Int16] = valueFunc[reflect.Int]
	valueFunc[reflect.Int32] = valueFunc[reflect.Int]
	valueFunc[reflect.Int64] = func(fieldValue *reflect.Value, tagValue string) {

		if fieldValue.Type() == reflect.TypeOf(time.Second) {
			value, _ := time.ParseDuration(tagValue)
			fieldValue.Set(reflect.ValueOf(value))
		} else {
			value, _ := strconv.ParseInt(tagValue, 10, 64)
			fieldValue.SetInt(value)
		}
	}
	valueFunc[reflect.Float32] = func(fieldValue *reflect.Value, tagValue string) {
		value, _ := strconv.ParseFloat(tagValue, 64)
		fieldValue.SetFloat(value)
	}

	valueFunc[reflect.Float64] = valueFunc[reflect.Float32]
	valueFunc[reflect.Uint] = func(fieldValue *reflect.Value, tagValue string) {
		value, _ := strconv.ParseUint(tagValue, 10, 64)
		fieldValue.SetUint(value)
	}

	valueFunc[reflect.Uint8] = valueFunc[reflect.Uint]
	valueFunc[reflect.Uint16] = valueFunc[reflect.Uint]
	valueFunc[reflect.Uint32] = valueFunc[reflect.Uint]
	valueFunc[reflect.Uint64] = valueFunc[reflect.Uint]

	valueFunc[reflect.String] = func(fieldValue *reflect.Value, tagValue string) {
		fieldValue.SetString(tagValue)
	}
}

// 设置枚举值
func setEnumValue(reflectValue *reflect.Value, tagValue string) {
	if f, ok := valueFunc[reflectValue.Kind()]; ok {
		f(reflectValue, tagValue)
	}
}

// GenerateEnum
// 如果一个函数的参数是接口类型，传进去的参数可以是指针，也可以不是指针，这得看你传的对象是怎么实现这个接口类型的
// 如果实现接口的方法的接收器是指针类型，则传参给接口类型时必需是指针，如果不是，则随便传
func GenerateEnum[T InterFaceEnum](enumStruct T) T {
	// 使用 reflect 包获取结构体的类型信息
	t := reflect.TypeOf(enumStruct)
	v := reflect.ValueOf(enumStruct)
	//修改开始
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	// 循环遍历结构体的所有字段
	for i := 0; i < t.NumField(); i++ {
		// 获取字段的名称、类型和注释
		field := t.Field(i)
		value := v.Field(i)

		name := field.Name
		key := field.Tag.Get("key")
		label := field.Tag.Get("label")
		if key == "" && label == "" {
			splits := strings.SplitN(string(field.Tag), " ", 2)
			for i := range splits {
				splits[i] = strings.TrimSpace(splits[i])
			}
			key, label = splits[0], splits[1]
		}

		if name == "Enum" && key != "" {
			if key != "" {
				enumStruct.InitMap(key)
				continue
			} else {
				fmt.Println("Enum 未指定 key")
				break
			}
		}
		if label != "" {
			if key == "" {
				fmt.Printf("field: %s key : %s label is empty\n", name, key)
				continue
			}
			setEnumValue(&value, key)
			enumStruct.SetLabel(value.Interface(), label)

		}
	}

	return enumStruct
}

/*
var ColorEnum = GenerateEnum(&struct {
	Enum   `key:"color" label:"颜色"`
	Red    string `key:"red" label:"红色"`
	Yellow string `key:"yellow" label:"黄色"`
	Black  int    `1 黑色`
	White  string    `white 白色`
}{})
*/

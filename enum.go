package enum

import (
	"fmt"
	"github.com/mcuadros/go-defaults"
	"reflect"
	"strconv"
	"time"
)

type ItemEnumMap map[any]string

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

var defaultFillerMap = map[string]*defaults.Filler{}

func GetDefaultFiller(tagLookup string) *defaults.Filler {
	defaultFiller, ok := defaultFillerMap[tagLookup]
	if !ok {
		defaultFiller = newDefaultFiller(tagLookup)
		defaultFillerMap[tagLookup] = defaultFiller
	}
	return defaultFiller
}

func newDefaultFiller(tagLookup string) *defaults.Filler {
	funcs := make(map[reflect.Kind]defaults.FillerFunc, 0)
	funcs[reflect.Bool] = func(field *defaults.FieldData) {
		value, _ := strconv.ParseBool(field.TagValue)
		field.Value.SetBool(value)
	}

	funcs[reflect.Int] = func(field *defaults.FieldData) {
		value, _ := strconv.ParseInt(field.TagValue, 10, 64)
		field.Value.SetInt(value)
	}

	funcs[reflect.Int8] = funcs[reflect.Int]
	funcs[reflect.Int16] = funcs[reflect.Int]
	funcs[reflect.Int32] = funcs[reflect.Int]
	funcs[reflect.Int64] = func(field *defaults.FieldData) {
		if field.Field.Type == reflect.TypeOf(time.Second) {
			value, _ := time.ParseDuration(field.TagValue)
			field.Value.Set(reflect.ValueOf(value))
		} else {
			value, _ := strconv.ParseInt(field.TagValue, 10, 64)
			field.Value.SetInt(value)
		}
	}

	funcs[reflect.Float32] = func(field *defaults.FieldData) {
		value, _ := strconv.ParseFloat(field.TagValue, 64)
		field.Value.SetFloat(value)
	}

	funcs[reflect.Float64] = funcs[reflect.Float32]

	funcs[reflect.Uint] = func(field *defaults.FieldData) {
		value, _ := strconv.ParseUint(field.TagValue, 10, 64)
		field.Value.SetUint(value)
	}

	funcs[reflect.Uint8] = funcs[reflect.Uint]
	funcs[reflect.Uint16] = funcs[reflect.Uint]
	funcs[reflect.Uint32] = funcs[reflect.Uint]
	funcs[reflect.Uint64] = funcs[reflect.Uint]

	funcs[reflect.String] = func(field *defaults.FieldData) {
		field.Value.SetString(field.TagValue)
	}

	types := make(map[defaults.TypeHash]defaults.FillerFunc, 1)
	types["time.Duration"] = func(field *defaults.FieldData) {
		d, _ := time.ParseDuration(field.TagValue)
		field.Value.Set(reflect.ValueOf(d))
	}
	return &defaults.Filler{FuncByKind: funcs, FuncByType: types, Tag: tagLookup}
}

// GenerateEnum
// 如果一个函数的参数是接口类型，传进去的参数可以是指针，也可以不是指针，这得看你传的对象是怎么实现这个接口类型的
// 如果实现接口的方法的接收器是指针类型，则传参给接口类型时必需是指针，如果不是，则随便传
func GenerateEnum[T InterFaceEnum](enumStruct T) T {
	// 使用 reflect 包获取结构体的类型信息
	t := reflect.TypeOf(enumStruct)

	//修改开始
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// 循环遍历结构体的所有字段
	for i := 0; i < t.NumField(); i++ {
		// 获取字段的名称、类型和注释
		field := t.Field(i)
		name := field.Name

		key := field.Tag.Get("key")
		label := field.Tag.Get("label")

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
			enumStruct.SetLabel(key, label)
		}
	}
	GetDefaultFiller("key").Fill(enumStruct)
	return enumStruct
}

/*
var ColorEnum = GenerateEnum(&struct {
	Enum   `key:"color" label:"颜色"`
	Red    string `key:"red" label:"红色"`
	Yellow string `key:"yellow" label:"黄色"`
	Black  int    `key:"black" label:"黑色"`
	White  int    `key:"white" label:"白色"`
}{})
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func main() {
	//reflect1()
	//reflect2()
	//reflect3()
	//reflect4()
	//reflect5()
	//reflect6()
	//reflect7()
	//reflect8()
	//reflect9()
	reflect10()
}

type person struct {
	Name string `json:"name" bson:"b_name"` // 必须是大写，外部可访问
	Age  int    `json:"age" bson:"a_name"`
}

func reflect10() {
	p := person{"yufei", 28}
	pt := reflect.TypeOf(p)
	pv := reflect.ValueOf(p)

	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")

	for i := 0; i < pt.NumField(); i++ {
		sf := pt.Field(i)
		key := sf.Tag.Get("json")
		jsonBuilder.WriteString("\"" + key + "\"")
		jsonBuilder.WriteString(":")

		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", pv.Field(i)))

		if i < pt.NumField()-1 {
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")
	fmt.Println(jsonBuilder.String())
}

func reflect9() {
	p := person{"yufei", 28}
	pt := reflect.TypeOf(p)
	for i := 0; i < pt.NumField(); i++ {
		sf := pt.Field(i)
		fmt.Println(sf.Name)
		fmt.Println(sf.Tag.Get("json"))
		fmt.Println(sf.Tag.Get("bson"))
	}
}

func reflect8() {
	p := person{"yufei", 28}
	pt := reflect.TypeOf(p)
	for i := 0; i < pt.NumField(); i++ {
		fs := pt.Field(i)
		fs.Tag.Get("json")
		fmt.Printf("field's name is %s, tag is %s \n", fs.Name, fs.Tag.Get("json"))
	}
}

func reflect7() {
	p := person{"yufei", 28}

	// 结构体 -》 json
	s, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(s))
	}

	// json str ->结构体
	str := "{\"name\":\"李四\",\"Age\":40}"
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p)
}

func reflect6() {
	p := person{"yufei", 28}

	pt := reflect.TypeOf(p)
	a := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	b := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fmt.Println(pt.Implements(a))
	fmt.Println(pt.Implements(b))
}

func reflect5() {
	p := person{"yufei", 28}
	tp := reflect.TypeOf(p)

	for i := 0; i < tp.NumField(); i++ {
		fmt.Println("字段", tp.Field(i).Name)
	}
	for i := 0; i < tp.NumMethod(); i++ {
		fmt.Println("方法", tp.Method(i).Name)
	}
}

func reflect4() {
	p := person{Name: "yufeii", Age: 28}
	pp := reflect.ValueOf(&p)
	pv := reflect.ValueOf(p)
	fmt.Println(pp.Kind())
	fmt.Println(pv.Kind())
}

func reflect3() {
	i := 3
	a := reflect.ValueOf(&i)
	a.Elem().SetInt(4)
	fmt.Println(i)

	p := person{Name: "yufei", Age: 18}
	pp := reflect.ValueOf(&p)
	pp.Elem().Field(0).SetString("yufeifei")
	pp.Elem().Field(1).SetInt(28)
	fmt.Println(p)
}

func (p person) String() string {
	return fmt.Sprintf("%s, %d", p.Name, p.Age)
}

func reflect2() {
	i := 2
	iv := reflect.ValueOf(i)
	i1 := iv.Interface().(int)
	fmt.Println(i1)
}

func reflect1() {
	i := 3
	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)
	fmt.Println(iv, it)
}

package main

import (
	"fmt"
	"reflect"
	"strings"
)

type order struct {
	orderId    int
	customerId int
}

func createQuery1(q interface{}) {
	t := reflect.TypeOf(q)
	k := t.Kind()
	v := reflect.ValueOf(q)
	k1 := v.Kind()
	fmt.Println(t)
	fmt.Println(k)
	fmt.Println(v)
	fmt.Println(k1)
}

func createQuery2(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("field num is:", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Type is :%T,Field is :%d,value is %v,i is %v\n", i, v.Field(i), v.Field(i), reflect.ValueOf(i))
		}
	}
}

func createQuery3(q interface{}) {
	valueOf := reflect.ValueOf(q)

	if valueOf.Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name() // 结构体名字
		query := fmt.Sprintf("insert into %s value (", t)

		for i := 0; i < valueOf.NumField(); i++ {
			kind := valueOf.Field(i).Kind()

			fmt.Println(reflect.TypeOf(valueOf.Field(i)).Name())
			switch kind {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, valueOf.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s,%d", query, valueOf.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, valueOf.Field(i).String())
				} else {
					query = fmt.Sprintf("%s,\"%s\"", query, valueOf.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}

// func createQuery4(q interface{}) {
// 	valueOf := reflect.ValueOf(q)
// 	m := map[int]interface{}{}
// 	if valueOf.Kind() == reflect.Struct {
// 		typeOf := reflect.TypeOf(q)
// 		query := fmt.Sprintf("insert into %s(", typeOf.Name())
// 		for i := 0; i < valueOf.NumField(); i++ {
// 			kind := valueOf.Field(i).Kind()
// 			name := typeOf.Field(i).Name
// 			switch kind {
// 			case reflect.Int:
//
// 			}
// 		}
// 	}
// }

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func main() {
	o := order{
		orderId:    2,
		customerId: 3,
	}
	createQuery1(o)
	createQuery2(o)
	createQuery3(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery3(e)
	var s string = ""
	fmt.Println(strings.Split(s, ":"))
}

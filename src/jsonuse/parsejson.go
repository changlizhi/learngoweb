package jsonuse

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

var jsonstr = `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`

/*
首先查找tag含有Foo的可导出的struct字段(首字母大写)
其次查找字段名是Foo的导出字段
最后查找类似FOO或者FoO这样的除了首字母之外其他大小写不敏感的导出字段
*/

type Server struct {
	ServerName string
	ServerIp   string
}
type Serverslice struct {
	Servers []Server
}

func ParseJson() {
	var s Serverslice
	json.Unmarshal([]byte(jsonstr), &s)
	fmt.Println(s)
}

/*
Go类型和JSON类型的对应关系如下：

bool 代表 JSON booleans,
float64 代表 JSON numbers,
string 代表 JSON strings,
nil 代表 JSON null.
*/

var b = []byte(`{"name":"wednesday","Age":6,"Parents":["Gmoez","Morticia"]}`)

func ParseJsonInterface() {
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		panic(err)
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type i dont' know")
		}
	}
	//以上为官方推荐方法
	//以下未bitly开源的方法
	js, err := simplejson.NewJson([]byte(
		`{
			"test":{
			"array":[1,"2",3],
			"int":10,
			"float":5.150,
			"bignum":12222222222222222333,
			"string":"simplejson",
			"boolean":true
			}
		}`))
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()
	fmt.Println(arr)
	fmt.Println(i)
	fmt.Println(ms)

}

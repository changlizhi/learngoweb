package jsonuse

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServerJson struct {
	ServerName string
	ServerIp   string
}
type ServerJsonSlice struct {
	Servers []ServerJson
}

type ServerJsonTag struct {
	ServerName string `json:"serverName"`
	ServerIp   string `json:"serverIp"`
}

type ServerJsonSliceTag struct {
	Servers []ServerJsonTag `json:"servers"`
}

func UseJsonExport() {
	var s ServerJsonSlice
	s.Servers = append(s.Servers, ServerJson{ServerIp: "127.0.0.1", ServerName: "ccc_VPN"})
	s.Servers = append(s.Servers, ServerJson{ServerIp: "127.0.0.2", ServerName: "lll_VPN"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

func UseJsonExportTag() {
	var s ServerJsonSliceTag
	s.Servers = append(s.Servers, ServerJsonTag{ServerIp: "127.0.0.1", ServerName: "ccc_VPN"})
	s.Servers = append(s.Servers, ServerJsonTag{ServerIp: "127.0.0.2", ServerName: "lll_VPN"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

/*
针对JSON的输出，我们在定义struct tag的时候需要注意的几点是:

字段的tag是"-"，那么这个字段不会输出到JSON
tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName
tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串
*/

type ServerShiLi struct {
	// ID 不会导出到JSON中
	ID int `json:"-"`

	// ServerName 的值会进行二次JSON编码
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`

	// 如果 ServerIP 为空，则不输出到JSON串中
	ServerIP string `json:"serverIP,omitempty"`
}

func UseShiLiJson() {
	s := ServerShiLi{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:    ``,
	}
	b, _ := json.Marshal(s)
	os.Stdout.Write(b)
}

/*
Marshal函数只有在转换成功的时候才会返回数据，在转换的过程中我们需要注意几点：

JSON对象只支持string作为key，所以要编码一个map，那么必须是map[string]T这种类型(T是Go语言中任意的类型)
Channel, complex和function是不能被编码成JSON的
嵌套的数据是不能编码的，不然会让JSON编码进入死循环
指针在编码的时候会输出指针指向的内容，而空指针会输出null
*/

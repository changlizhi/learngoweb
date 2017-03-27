package uselocale

import (
	"fmt"
	"strings"
	"time"
)

var locales map[string]map[string]string

func Locale_map() {
	locales = make(map[string]map[string]string, 2)
	en := make(map[string]string, 10)
	en["pea"] = "pea"
	en["bean"] = "bean"
	locales["en"] = en
	cn := make(map[string]string, 10)
	cn["pea"] = "豌豆"
	cn["bean"] = "毛豆"
	locales["zh-CN"] = cn

	en["time_zone"] = "America/Chicago"
	cn["time_zone"] = "Asia/Shanghai"
	lang := "zh-CN"
	loc, _ := time.LoadLocation(msg(lang, "time_zone"))
	t := time.Now()
	t = t.In(loc)
	fmt.Println(t.Format(time.RFC3339))

	en["date_format"] = "%Y-%m-%d %H:%M:%S"
	cn["date_format"] = "%Y年-%m月-%d日 %H时:%M分:%S秒"
	fmt.Println(date(msg(lang, "date_format"), t))

}
func date(format string, t time.Time) string {
	_, month, day := t.Date()
	hour, min, sec := t.Clock()

	fmt.Println(sec)
	format = strings.Replace(format, "%Y", "2017", -1)
	format = strings.Replace(format, "%m", string(month), -1)
	format = strings.Replace(format, "%d", string(day), -1)
	format = strings.Replace(format, "%H", string(hour), -1)
	format = strings.Replace(format, "%M", string(min), -1)
	format = strings.Replace(format, "%S", string(sec), -1)
	return format
}

func msg(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok := v[key]; ok {
			return v2
		}
	}
	return ""
}

//views
//  |--en // 英文模板
//      |--images // 存储图片信息
//      |--js // 存储 JS 文件
//      |--css // 存储 css 文件
//      index.tpl // 用户首页
//      login.tpl // 登陆首页
//  |--zh-CN // 中文模板
//      |--images
//      |--js
//      |--css
//      index.tpl
//      login.tpl

// s1, _ := template.ParseFiles("views"+lang+"index.tpl")
// VV.Lang=lang
// s1.Execute(os.Stdout, VV)

//<script type="text/javascript" src="views/{{.VV.Lang}}/js/jquery/jquery-1.8.0.min.js">

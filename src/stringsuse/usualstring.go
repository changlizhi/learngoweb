package stringsuse

import (
	"fmt"
	"strconv"
	"strings"
)

func Usual_string_use() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("", ""))

	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ","))

	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))

	fmt.Println("ba" + strings.Repeat("na", 2))

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panda", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo sod's"))

	fmt.Printf("[%q]", strings.Trim(" !!! sd!!   ", " ")) //去除指定字符串

}

func Usual_conv() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 234, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}

func Usual_format() {
	a := strconv.FormatBool(false) //转换为字符串类型
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)
}
func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func Usual_parse() {
	a, err := strconv.ParseBool("false")
	checkErr(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkErr(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkErr(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkErr(err)
	e, err := strconv.Atoi("1023")
	checkErr(err)
	fmt.Println(a, b, c, d, e)

}

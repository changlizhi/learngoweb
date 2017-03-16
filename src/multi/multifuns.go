package multi

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "efefe"

// g, h := 222,"fff"这种不带var 的声明只能在函数中使用

const (
	x = iota
	y = iota
)

func jichu() {
	//xz := 333
	//fz := "3332"
	//fmt.Println("ddd")
	//fz, xz = xz,fz
	//fmt.Println(xz, fz, x,&fz)
	arr_test := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_test)
	slice_teset := arr_test[0:3]
	fmt.Println(slice_teset)
	_ = append(slice_teset, 8)
	fmt.Println(slice_teset)
	fmt.Println(arr_test)
}

func use_go_to() {
	i := 0
Here:
	println(i)
	i++
	if i < 10 {
		goto Here
	}
	println(i)
}

func use_for() {
	sum := 0
	for index := 0; index < 11; index++ {
		sum += index
	}
	fmt.Println("简单for循环")
	fmt.Println(sum)
	for index := 10; index > 0; index-- {
		if index == 5 {
			continue
		}
		fmt.Println(index)
		if index == 5 {
			break
		}
		fmt.Println(index)
	}
	var in_map map[string]string
	in_map["one"] = "test_one"
	in_map["one"] = "test_double_one"
	in_map["two"] = "test_two"
	in_map["three"] = "test_three"
	in_map["four"] = "test_four"

}

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

type Human struct {
	name   string
	age    int
	weight int
}
type Student struct {
	//Student 拥有了Human的所有字段，匿名字段的意义
	Human
	speciality string
}

//go hello(a,b,c)启动一个线程管理器goroutine

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func personuse() {

	var P person
	P.name = "clz"
	P.age = 28
	fmt.Printf("the person's name= %s\n", P.name)

	P2 := person{"clz2", 30}
	fmt.Printf("clz2 's age = %d", P2.age)

	P3 := person{age: 33, name: "clz3"}
	fmt.Printf("\n clz3's age = %d", P3.age)

	var tom person
	tom.name, tom.age = "tom", 13
	bob := person{age: 23, name: "Bob"}
	paul := person{"pual", 43}
	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)
	fmt.Printf("\n Of %s and %s,%s is Older by %d year! ", tom.name, bob.name, tb_Older, tb_diff)
	fmt.Printf("\n Of %s and %s,%s is Older by %d year! ", tom.name, paul.name, tp_Older, tp_diff)
	fmt.Printf("\n Of %s and %s,%s is Older by %d year! ", bob.name, paul.name, bp_Older, bp_diff)

}
func niming() {

	mark := Student{Human{"mark", 25, 120}, "computer gift"}
	fmt.Printf("mark's name is %s", mark.name)
	mark.speciality = "AI"
	fmt.Printf("mark's new speciality is %s\n", mark.speciality)
	mark.age = 33
	fmt.Println("mark's new age is %d", mark.age)

}

func routineUse() {
	go say("word")
	say("hello")

}

type Hu struct {
	phone string
	age   int
}
type Tao struct {
	phone string
	age   int
}

type HuTao struct {
	Hu
	Tao
}

func hutaofunc() {
	hutao := HuTao{Hu{phone: "hu phone", age: 12}, Tao{phone: "tao phone", age: 13}}
	fmt.Printf("hutao's age is %d ", hutao.Hu.age)
}

func sumchannel(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func usechannel() {
	a := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sumchannel(a[:len(a)/2], c)
	go sumchannel(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func bufferChannel() {
	c := make(chan int, 2)
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func feibo(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func rangeChannel() {
	c := make(chan int, 10)
	go feibo(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func select_feibo(c, quit chan int) {
	x, y := 1, 1
	o := make(chan bool)
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		case <-time.After(5 * time.Second):
			println("timeout")
			o <- true
			break

		default:
			fmt.Println("当c阻塞时执行")
		}
	}
}

func select_channel() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	select_feibo(c, quit)
}

func SayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的}
}
func Listenhttp() {
	http.HandleFunc("/", SayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloNameMy(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func sayhelloNameMy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello new route!</h1>")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("templates/login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//yanzheng token hefaxing
		} else {
			//bucunzai baocuo
		}

		fmt.Println("username:", r.Form["username"])
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("password:", r.Form["username"])
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
	//
	//if len(r.Form["username"][0]) == 0 {
	//	//为空的处理
	//}
	//getint, err := strconv.Atoi(r.Form.Get("age"))
	//if err != nil {
	//	//数字转化错误
	//}
	//if getint > 100 {
	//	//太大了
	//}
	//if m, _ := regexp.MatchString("^[0-9]+$]", r.Form.Get("age")); !m {
	//	//正则表达式匹配数字
	//	return
	//}
	//if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("realName")); !m {
	//	//正则表达式匹配中文
	//	return
	//}
	//if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engName")); !m {
	//	return
	//}
	//if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
	//	return
	//}
	//if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("phone")); !m {
	//	return
	//}
	//
	////验证下拉框
	//slice := []string{"apple", "pear", "banana"}
	//for _, v := range slice {
	//	if v == r.Form.Get("fruit") {
	//		return
	//	}
	//}
	//slice2 := []string{"1","2"}
	//for _, v := range slice2 {
	//	if v == r.Form.Get("gender") {
	//		return
	//	}
	//}
	//slice3 := []string{"football", "basketball", "tennis"}
	//a := Slice_diff(r.Form["interest"], slice3)
	//if a == nil {
	//	return
	//}
	////shijian zhuanhuan
	//t := time.Date(2017, time.January, 3, 12, 1, 43, 0, time.UTC)
	//fmt.Printf("go launced at %s \n", t.Local())
	//
	////18为身份证
	//if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
	//	return
	//}

}
func Slice_diff(i []string, i2 []string) interface{} {
	return nil
}

type FileHeader struct {
	Filname string
	Header  textproto.MIMEHeader
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		curTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("templates/upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func useLogin() {
	http.HandleFunc("/", SayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error open file")
		return err
	}
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func usePost() {
	//模拟客户端请求
	target_url := "http://localhost:9090/upload"
	fileName := "texts/foraes.txt"
	postFile(fileName, target_url)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func useDb() {

	//"mysql",
	db, err := sql.Open("mysql", "root:root@/test")
	checkErr(err)
	stmt, err := db.Prepare("INSERT userinfo set username=?,departname=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("ast", "yanfa", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("updatename", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	rows, err := db.Query("select * from userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}

type Person struct {
	NAME  string
	PHONE string
}

type Men struct {
	Persons []Person
}

const (
	URL = "127.0.0.1:27017"
)

func useMongo() {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("mydb")
	collection := db.C("person")
	countNum, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("things obj counts:", countNum)
	temp := &Person{
		PHONE: "12221",
		NAME:  "clsznem",
	}
	err = collection.Insert(&Person{"ale", "028-99322221"}, temp)

	if err != nil {
		panic(err)
	}
	result := Person{}
	err = collection.Find(bson.M{"phone": "222"}).One(&result)
	fmt.Println("Phone:", result.NAME, result.PHONE)

	var personAll Men
	iter := collection.Find(nil).Iter()
	for iter.Next(&result) {
		fmt.Printf("Result:%v\n", result.NAME)
		personAll.Persons = append(personAll.Persons, result)
	}
	err = collection.Update(bson.M{"name": "ccc"}, bson.M{"$set": bson.M{"name": "ddd"}})
	err = collection.Update(bson.M{"name": "ddd"}, bson.M{"$set": bson.M{"phone": "1222211"}})
	err = collection.Update(bson.M{"name": "aaa"}, bson.M{"phone": "123", "name": "bbb"})
	_, err = collection.RemoveAll(bson.M{"name": "Ale"})
	_, err = collection.RemoveAll(bson.M{"name": "clsznem"})
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionId() string
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(max_life_time int64)
}

var provides = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session:Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session:Register called twice for provide " + name)
	}
	provides[name] = provider
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.max_life_time)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

type Manager struct {
	cookieName    string
	lock          sync.Mutex
	provider      Provider
	max_life_time int64
}

func newManager(providerName, cookieName string, max_life_time int64) (*Manager, error) {
	provider, ok := provides[providerName]
	if !ok {
		return nil, fmt.Errorf("session:unknown provide %q (forgotten import?)", providerName)
	}
	return &Manager{provider: provider, cookieName: cookieName, max_life_time: max_life_time}, nil
}

var global_sessions *Manager

func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.max_life_time)
	time.AfterFunc(time.Duration(manager.max_life_time), func() {
		manager.GC()
	})
}
func init() {
	global_sessions, _ = newManager("memory", "gosessionid", 3600)
	//go global_sessions.GC()
}

func sessionLogin(w http.ResponseWriter, r *http.Request) {
	sess := global_sessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
		http.Redirect(w, r, "/", 302)
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := global_sessions.SessionStart(w, r)
	create_time := sess.Get("create_time")
	if create_time == nil {
		sess.Set("create_time", time.Now().Unix())
	} else if (create_time.(int64) + 360) < (time.Now().Unix()) {
		global_sessions.SessionDestroy(w, r)
		sess = global_sessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}

func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}

}

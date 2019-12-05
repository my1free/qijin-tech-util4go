package main

import (
		"fmt"
	"io/ioutil"
	"time"
	"github.com/kataras/iris/core/errors"
		"net/http"
	"github.com/google/uuid"
	"encoding/base64"
	"crypto/md5"
	"encoding/json"
	"strings"
)

const (
	cacheSize = 100 * 1024 * 1024
	expire    = 30 // expire in 30 seconds
)

type redisReqTypeVal struct {
	Priority int
	Resource map[string]interface{} `json:"resource"`
	Extra    string
}

var expressions = []string{
	"foo > 0",
	"bar.Value in ['a', 'b', 'c']",
	"name matches '^hello.+$'",
	"version > '40004'",
	"now().Sub(startedAt).String()",
	"all(tweets, {len(.Message) <= 280}) ",
}

var environment = map[string]interface{}{
	"foo":       1,
	"bar":       struct{ Value string }{"c"},
	"name":      "hello world",
	"startedAt": time.Now(),
	"version":   "4.0.1",
	"now":       func() time.Time { return time.Now() },
	"tweets":    []tweet{{"first tweet"}},
}

type tweet struct {
	Message string
}

func main() {
	s := time.Now().UTC().Format(http.TimeFormat)
	fmt.Println(s)

	fmt.Println(uuid.New().String())

	body := make(map[string]interface{})
	body["asdf"] = "asdf"
	bodyB, _ := json.Marshal(body)
	fmt.Println(base64.StdEncoding.EncodeToString(md5.New().Sum(bodyB)))

	fmt.Println(IsBlank(" s "))
	fmt.Println(IsBlank(""))
	fmt.Println(IsBlank("  "))

}

func IsBlank(s string) bool {
	return len(strings.Trim(s, " ")) == 0
}


func Tt() (string, error) {
	return "s", errors.New("asdfasdf")
}

func Retry(f func() (error), maxTimes int, interval time.Duration) error {
	// 如果 maxTimes 参数非法，则只执行一次
	if maxTimes <= 0 {
		f()
		return nil
	}
	var err error
	for t := maxTimes; t > 0; t-- {
		if err = f(); err != nil {
			time.Sleep(interval)
			continue
		} else {
			break
		}
	}
	return err
}

func WriteFile() {
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Print("asdfasd" + "\n")
	d1 := []byte("hello\ngo\n")
	ioutil.WriteFile("test.txt", d1, 0644)
}

func gor(chs chan int) {
	time.Sleep(time.Duration(1) * time.Second)
	chs <- 10
	fmt.Print("go routine\n")
}

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}

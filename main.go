package main

import (
		"fmt"
	"io/ioutil"
	"time"
	"github.com/kataras/iris/core/errors"
							"strings"
	"encoding/json"
	"git.inke.cn/BackendPlatform/golang/logging"
	"git.inke.cn/gmu/server/common/toolbox/utils"
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

type OriginText struct {
	Timestamp int64  `json:"timestamp"`
	Host      string `json:"host"`
	Appkey    string `json:"appkey"`
	From      string `json:"from"`
	To        string `json:"to"`
	MsgID     string `json:"msg_id"`
	ChatType  string `json:"chat_type"`
	Payload   struct {
		Ext    map[string]interface{} `json:"ext"`
		Bodies []struct {
			URL    string  `json:"url"`
			Type   string  `json:"type"`
			Msg    string  `json:"msg"`
			Addr   string  `json:"addr"`
			Lng    float64 `json:"lng"`
			Lat    float64 `json:"lat"`
			Length int     `json:"length"`
			Size   struct {
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"size"`
			Secret     string `json:"secret"`
			Filename   string `json:"filename"`
			FileLength int    `json:"file_length"`
		} `json:"bodies"`
	} `json:"payload"`
	CallID    string `json:"callId"`
	EventType string `json:"eventType"`
	Security  string `json:"security"`
}

func main() {
	currTime := "2020-02-02 02:02:10"
	utils.StrToTime(currTime, utils.DefaultLayout).Unix()
	originText:="{\"timestamp\":1580785500187,\"host\":\"msync@vip6-ali-beijing-msync-55\",\"appkey\":\"gmu#jimu\",\"from\":\"new_jimu_gmu_5351659\",\"to\":\"new_jimu_gmu_1\",\"msg_id\":\"701027569301981696\",\"chat_type\":\"chat\",\"payload\":{\"ext\":{\"msgtype\":{\"choice\":{\"menuid\":\"TransferToKf\"}},\"TrackId\":\"535165911580785500100\",\"HXFromUserVipStatus\":\"0\",\"HXFromUserSenderTime\":\"1580785500\",\"HXFromUserName\":\"Hey\",\"HXFromUserIntent\":\"103\",\"HXFromUserID\":\"5351659\",\"HXFromUserAvatar\":\"http:\\/\\/avatar.cdn.gmugmu.com\\/5351659_917537\",\"HXConversationId\":\"new_jimu_gmu_5351659new_jimu_gmu_1\"},\"bodies\":[{\"type\":\"txt\",\"msg\":\"\\u8f6c\\u4eba\\u5de5\\u5ba2\\u670d\"}]},\"callId\":\"gmu#jimu_701027569301981696\",\"eventType\":\"chat\",\"security\":\"8aa8bf0db2b9b26117a4e96862d547e9\"}"

	originTextS := new(OriginText)
	if err := json.Unmarshal([]byte(originText), originTextS); err != nil {
		logging.Errorf("unmarshal originText error | originText=%s | err=%+v", originText, err)
	}
	payload := originTextS.Payload
	ext := payload.Ext
	trackId := ext["TrackId"].(string)
	fmt.Println(originTextS)
	fmt.Println(payload)
	fmt.Println(ext)
	fmt.Println(trackId)

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

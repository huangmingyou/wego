package wxapi

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = fromUserName
	textResponseBody.ToUserName = toUserName
	textResponseBody.MsgType = "text"
	textResponseBody.Content = content
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}

func parseTextRequestBody(r *http.Request) *TextRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	requestBody := &TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

func msgrecv(r *http.Request, w http.ResponseWriter) {
	uinfo := new(Userinfo)
	textRequestBody := parseTextRequestBody(r)
	if textRequestBody != nil {
		fmt.Printf("recv %s msg: %s\n", textRequestBody.FromUserName, textRequestBody.Content)
	}
	//access token 需要做成一个服务
	getuserinfo(textRequestBody.FromUserName,atoken,uinfo)
	fmt.Println(uinfo.Nickname)
	responseTextBody, _ := makeTextResponseBody(textRequestBody.ToUserName,textRequestBody.FromUserName,uinfo.Nickname+",你好!")
	//回用户信息
	fmt.Fprintf(w,"%s\n",string(responseTextBody))
}

package wxapi
import (
	"encoding/xml"
	"time"
)

// msg

type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}

type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

//user

type Userinfo struct {
        Subscribe int `json:"subscribe"`
        Openid    string  `json:"openid"`
        Nickname  string  `json:"nickname"`
        Sex       int `json:"sex"`
        City    string  `json:"city"`
        Country string  `json:"country"`
        Province string  `json:"province"`
        Language string  `json:"language"`
        Headimgurl string  `json:"headimgurl"`
        Subscribe_time int64 `json:"subscribe_time"`
        Unionid string  `json:"unionid"`
        Remark string  `json:"remark"`
        Groupid int64 `json:"groupid"`
}



// access token

type Accesstoken struct {
        AccessToken string  `json:"access_token"`
        ExpiresIn   float64 `json:"expires_in"`
}

// weixin error
type Wxerr struct {
        Errcode float64
        Errmsg  string
}



package wxapi

// 用户管理

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)


//todo: actoken 不应该作为参数传递
func getuserinfo(openid string, actoken string,useri *Userinfo) (err error){
	url := "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + actoken + "&openid=" + openid + "&lang=zh_CN"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body,&useri)

	return
}




package wxapi

import (
	"fmt"
	"net/http"
	"sort"
)

const token = "debian1980"

func Wxinit(w http.ResponseWriter, r *http.Request) {
	//http://mp.weixin.qq.com/wiki/17/2d4265491f12608cd170a95559800f2d.html
	//微信验证服务器地址有效性
	/*
	   加密/校验流程如下：
	   1. 将token、timestamp、nonce三个参数进行字典序排序
	   2. 将三个参数字符串拼接成一个字符串进行sha1加密
	   3. 开发者获得加密后的字符串可与signature对比，标识该请求来源于微信
	*/

	//判断请求的种类,可以用nginx实现

	switch r.Method {

	case "GET":
		{

			r.ParseForm()

			sig := r.Form.Get("signature")
			es := r.Form.Get("echostr")
			ts := r.Form.Get("timestamp")
			non := r.Form.Get("nonce")

			tmps := []string{token, ts, non}
			sort.Strings(tmps)
			tmpStr := tmps[0] + tmps[1] + tmps[2]
			sha1st := Str2sha1(tmpStr)

			//debug
			fmt.Println("none", non)
			fmt.Println("echostr", es)
			fmt.Println("timestamp", ts)
			fmt.Println("sig", sig)
			fmt.Println("sha1", sha1st)

			if sha1st == sig {
				fmt.Fprintf(w, "%s", es)
			} else {
				fmt.Fprintf(w, "%s", "Wrong token!")
				fmt.Println("403")
			}
		}

	case "POST":
		{
			msgrecv(r, w)
		}
	}
}

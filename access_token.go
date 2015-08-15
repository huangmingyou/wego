package wxapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)


const dbfile="/home/hmy/wego.db"

func getat(id, sec string) (string, float64, error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + id + "&secret=" + sec

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	if bytes.Contains(body, []byte("access_token")) {
		atr := Accesstoken{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			return "", 0.0, err
		}
		return atr.AccessToken, atr.ExpiresIn, nil
	} else {
		fmt.Println("return err")
		ater := Wxerr{}
		err = json.Unmarshal(body, &ater)
		if err != nil {
			return "", 0.0, err
		}
		return "", 0.0, fmt.Errorf("%s", ater.Errmsg)
	}

}
//获取 微信 access_token 保存到数据库，有一个过期时间戳

func Fetchtoken(id,sec string) (string,error) {
	db,err := sql.Open("sqlite3",dbfile)
	if err != nil { 
		log.Fatal(err)
	}

	defer db.Close()

	rows,err := db.Query("select tk from access_token where ts <"+"1438765601"+" limit 1")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var token string
		rows.Scan(&token)
		if token != "" {
			fmt.Println(token)
			return token,err
	}
	}
	return "",err

}

package main

/* huangmingyou@gmail.com
2015-07-29
*/
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"wego/wxapi"
)

func main() {

	wxapi.Fetchtoken("ok","ok")
	http.HandleFunc("/", wxapi.Wxinit)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("wego: ", err)
		fmt.Println("err\n")
		os.Exit(1)
	}
}

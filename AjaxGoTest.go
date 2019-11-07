package main

import (
	"fmt"
	"github.com/Ericwyn/AjaxGo/Ajax"
)

func main() {
	Ajax.Send(Ajax.Request{
		Url:    "http://ip-api.com/json/?lang=zh-CN",
		Method: Ajax.GET,
		Success: func(response *Ajax.Response) {
			fmt.Println(response.Body)
		},
		Fail: func(status int, errMsg string) {
			fmt.Println("request Error")
		},
		Always: func() {
			fmt.Println("always run this func before request")
		},
	})
}

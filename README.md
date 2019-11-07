# AjaxGo
send http request with Golang, just like javascript Ajax

## Introduction

    func main() {
    	Ajax.Send(Ajax.Request{
    		Url:"http://ip-api.com/json/?lang=zh-CN",
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

## Support
 - GET Request
     - send URL Data
 - POST Request
     - send JSON data
     - send Form Data
     - send URL Data
 - Callback
     - Success Callback
     - Fail Callback
     - Always Run Callback
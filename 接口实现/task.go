
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)
func main(){
	//第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	http.HandleFunc("/shopee/entry_task", test)
	//服务器要监听的主机地址和端口号
	err:=http.ListenAndServe(":8099", nil)
	if err != nil {
		fmt.Println("ListenAnerve error: ", err.Error())
	}
}
//定义接口返回struct对象
type Resp struct {
	Code    int `json:"err_code"`
	Msg     string `json:"err_msg"`
	Refenence string `json:"refenence"`
}



func test(writer http.ResponseWriter,  request *http.Request)  {

	var result  Resp

	//获取客户端通过GET/POST方式传递的参数
	request.ParseForm()
	param_number,found1:=request.Form["a"]
	param_name,found2:=request.Form["b"]
	//判断是否传递参数a,b ,参数未传返回system error

	//time.Sleep(time.Duration(10)*time.Millisecond)
	if!(found1 && found2){
		result.Code = 11
		result.Msg = "system error"
		result.Refenence="Parameter a,b must be passed"
		if err := json.NewEncoder(writer).Encode(result); err != nil {
			log.Fatal(err)

		}
		return
	}
	//判断参数a,b是否为空
	if (param_number[0]==""||param_name[0]==""){
		result.Code = 21
		result.Msg = "empry or wrong params"
		result.Refenence="Parameter cannot be empty"
		if err := json.NewEncoder(writer).Encode(result); err != nil {
			log.Fatal(err)

		}
		return
	}
	name:=param_name[0]
	//判断参数a是否为int类型，或数字组成的字符串
	number,err :=strconv.Atoi(param_number[0])
	if err != nil{
		result.Code = 21
		result.Msg = "empry or wrong params"
		result.Refenence="Incorrect parameter type"
		if err := json.NewEncoder(writer).Encode(result); err != nil {
			log.Fatal(err)

		}
		return
	}


	result.Code = 0
	result.Msg = "success"
	result.Refenence=fmt.Sprintf("No.%d is %s",number,name)
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)

	}
}
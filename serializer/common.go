/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-16 17:25:51
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-21 16:37:24
 * @FilePath: /go-mall/serializer/common.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

func BuildListResponse(item interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data:   DataList{Item: item, Total: total},
		Msg:    "ok",
	}
}

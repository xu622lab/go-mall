/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-21 20:12:01
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-21 20:14:28
 * @FilePath: /go-mall/api/v1/common.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"encoding/json"
	"go-mall/serializer"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 400,
			Msg:    "JSON类型不匹配",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: 400,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}

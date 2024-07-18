/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-16 16:00:03
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-18 10:29:57
 * @FilePath: /go-mall/pkg/e/code.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	ErrorExistUser             = 30001
	ErrorFailEncryption        = 30002
	ErrorUserNotFound          = 30003
	ErrorNotCompare            = 30004
	ErrorAuthToken             = 30005
	ErrorAuthCheckTokenTimeout = 30006
)

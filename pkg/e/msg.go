/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-16 16:00:23
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-25 10:46:55
 * @FilePath: /go-mall/pkg/e/msg.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",

	ErrorExistUser:             "用户名已存在",
	ErrorFailEncryption:        "密码加密失败",
	ErrorUserNotFound:          "用户不存在",
	ErrorNotCompare:            "密码错误",
	ErrorAuthToken:             "token签发失败",
	ErrorAuthCheckTokenTimeout: "token过期",
	ErrorUploadFail:            "头像上传失败",
	ErrorSendEmail:             "邮件发送失败",
	ErrorProductImgUpload:      "商品图片上传错误",
}

// GetMsg获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}

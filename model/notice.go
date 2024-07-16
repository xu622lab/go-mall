/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-16 11:31:49
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-16 11:32:02
 * @FilePath: /go-mall/model/notice.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// Notice 公告模型 存放公告和邮件模板
type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}

/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-25 15:55:51
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-25 16:01:49
 * @FilePath: /go-mall/cache/key.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey = "rank"
)

// 标识一个商品点击量的key
func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}

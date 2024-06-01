/**
 * Jwt工具类
 * @author Xinze
 * @since 2021/1/12
 * @File : jwt
 */
package utility

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// 获取客户端IP
func GetClientIp(r *ghttp.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.GetClientIp()
	}
	return ip
}

func InStringArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// 数组反转
func Reverse(arr *[]string) {
	length := len(*arr)
	var temp string
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
}

// GetPlantformUid 获取平台用户Id
func GetPlantformUid(serviceUserId string, userId string) int {
	puid := fmt.Sprintf("%s-%s", serviceUserId, userId)
	return BkdrHash(puid)
}

// BkdrHash 计算BKDR哈希值
func BkdrHash(str string) int {
	seed := 131
	var hash int

	for _, char := range str {
		hash = seed*hash + int(char)
	}

	return hash & 0x7FFFFFFF
}

// 生成4位随机数字字符串
func GenerateRandomNumbers(n int) string {
	rand.Seed(time.Now().UnixNano())
	min := int64(1 * int(math.Pow10(n-1))) // 1000 for 4 digits
	max := int64(1*int(math.Pow10(n)) - 1) // 9999 for 4 digits
	return strconv.FormatInt(rand.Int63n(max-min+1)+min, 10)
}

func IsEmpty(value interface{}, traceSource ...bool) bool {
	return g.IsEmpty(value, traceSource...)
}

func IsNotEmpty(value interface{}, traceSource ...bool) bool {
	return !g.IsEmpty(value, traceSource...)
}

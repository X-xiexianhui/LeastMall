package common

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gomarkdown/markdown"
	_ "github.com/jinzhu/gorm"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"regexp"
	"strconv"
	"time"
)

// TimestampToDate 时间戳转换为日期格式
func TimestampToDate(timestamp int) string {

	t := time.Unix(int64(timestamp), 0)

	return t.Format("2006-01-02 15:04:05")
}

// GetUnix 获取当前时间戳
func GetUnix() int64 {
	fmt.Println(time.Now().Unix())
	return time.Now().Unix()
}

// GetUnixNano 获取时间戳Nano时间
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// Md5 Md5加密
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

// VerifyEmail 验证邮箱
func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// FormatDay 获取日期
func FormatDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// GenerateOrderId 生成订单号
func GenerateOrderId() string {
	template := "200601021504"
	return time.Now().Format(template) + GetRandomNum()
}

// SendMsg 发送验证码
func SendMsg(str string) {
	// 短信验证码需要到相关网站申请
	// 目前先固定一个值
	err := ioutil.WriteFile("test_send.txt", []byte(str), 06666)
	if err != nil {
		return
	}
}

// FormatAttribute 格式化级标题
func FormatAttribute(str string) string {
	md := []byte(str)
	htmlByte := markdown.ToHTML(md, nil, nil)
	return string(htmlByte)
}

// Mul 乘法的函数
func Mul(price float64, num int) float64 {
	return price * float64(num)
}

// GetRandomNum 封装一个生产随机数的方法
func GetRandomNum() string {
	var str string
	for i := 0; i < 4; i++ {
		current := rand.Intn(10) //0-9   "math/rand"
		str += strconv.Itoa(current)
	}
	return str
}

func FormatBase64(img *multipart.FileHeader) string {
	picture, _ := img.Open()
	data, _ := ioutil.ReadAll(picture)
	base64Str := base64.StdEncoding.EncodeToString(data)
	return base64Str
}

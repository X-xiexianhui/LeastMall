package frontend

import (
	"github.com/gin-gonic/gin"
	"leastMall_gin/common"
	"leastMall_gin/conn"
	"leastMall_gin/models"
	"net/http"
	"regexp"
	"strings"
)

func Login(c *gin.Context) {
	//c.Data["prevPage"] = c.Ctx.Request.Referer()
	//c.TplName = "frontend/auth/login.html"
	c.HTML(http.StatusOK, "frontend/auth/login.html", gin.H{"prevPage": c.Request.Referer()})
}

// GoLogin 登陆
func GoLogin(c *gin.Context) {
	phone := c.GetString("phone")
	password := c.GetString("password")
	password = common.Md5(password)
	var user []models.User
	conn.Db.Where("phone=? AND password=?", phone, password).Find(&user)
	if len(user) > 0 {
		models.Cookie.Set(c, "userinfo", user[0])
		data := map[string]interface{}{
			"success": true,
			"msg":     "用户登陆成功",
		}
		c.JSON(http.StatusOK, data)
		return
	} else {
		data := map[string]interface{}{
			"success": false,
			"msg":     "用户名或密码不正确",
		}
		c.JSON(http.StatusOK, data)
		return
	}
}

// LoginOut 退出登陆
func LoginOut(c *gin.Context) {
	models.Cookie.Remove(c, "userinfo", "")
	c.Redirect(http.StatusFound, c.Request.Referer())
}

// RegisterStep1 注册第一步
func RegisterStep1(c *gin.Context) {
	tpl := "frontend/auth/register_step1.html"
	c.Redirect(http.StatusOK, tpl)
}

// RegisterStep2 注册第二步
func RegisterStep2(c *gin.Context) {
	sign := c.GetString("sign")
	//phoneCode := c.GetString("phone_code")
	////验证图形验证码和前面是否正确
	//sessionPhotoCode := c.GetSession("phone_code")
	//if phoneCode != sessionPhotoCode {
	//	c.Redirect("/auth/registerStep1", 302)
	//	return
	//}
	var userTemp []models.UserSms
	conn.Db.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) > 0 {
		//c.Data["sign"] = sign
		//c.Data["phone_code"] = phoneCode
		//c.Data["phone"] = userTemp[0].Phone
		c.HTML(http.StatusOK, "frontend/auth/register_step2.html", gin.H{
			"sign":  sign,
			"phone": userTemp[0].Phone,
		})
	} else {
		c.Redirect(302, "/auth/registerStep1")
		return
	}
}

// RegisterStep3 注册第三步
func RegisterStep3(c *gin.Context) {
	sign := c.GetString("sign")
	smsCode := c.GetString("sms_code")
	if smsCode != "5259" {
		c.Redirect(302, "/auth/registerStep1")
		return
	}
	var userTemp []models.UserSms
	conn.Db.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) > 0 {
		c.HTML(http.StatusOK, "frontend/auth/register_step3.html", gin.H{
			"sign":     sign,
			"sms_code": smsCode,
		})
	} else {
		c.Redirect(302, "/auth/registerStep1")
		return
	}
}

// SendCode 发送验证码
func SendCode(c *gin.Context) {
	phone := c.GetString("phone")
	c.SetSession("phone_code", phoneCode)
	pattern := `^[\d]{11}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(phone) {
		data := map[string]interface{}{
			"success": false,
			"msg":     "手机号格式不正确",
		}
		c.JSON(http.StatusOK, data)
		return
	}
	var user []models.User
	conn.Db.Where("phone=?", phone).Find(&user)
	if len(user) > 0 {
		data := map[string]interface{}{
			"success": false,
			"msg":     "此用户已存在",
		}
		c.JSON(http.StatusOK, data)
		return
	}

	addDay := common.FormatDay()
	ip := strings.Split(c.Request.RemoteAddr, ":")[0]
	sign := common.Md5(phone + addDay) //签名
	smsCode := common.GetRandomNum()
	var userTemp []models.UserSms
	conn.Db.Where("add_day=? AND phone=?", addDay, phone).Find(&userTemp)
	var sendCount int
	conn.Db.Where("add_day=? AND ip=?", addDay, ip).Table("user_temp").Count(&sendCount)
	//验证IP地址今天发送的次数是否合法
	if sendCount <= 10 {
		if len(userTemp) > 0 {
			//验证当前手机号今天发送的次数是否合法
			if userTemp[0].SendCount < 5 {
				common.SendMsg(smsCode)
				c.SetSession("sms_code", smsCode)
				oneUserSms := models.UserSms{}
				conn.Db.Where("id=?", userTemp[0].Id).Find(&oneUserSms)
				oneUserSms.SendCount += 1
				conn.Db.Save(&oneUserSms)
				data := map[string]interface{}{
					"success":  true,
					"msg":      "短信发送成功",
					"sign":     sign,
					"sms_code": smsCode,
				}
				c.JSON(http.StatusOK, data)
				return
			} else {
				data := map[string]interface{}{
					"success": false,
					"msg":     "当前手机号今天发送短信数已达上限",
				}
				c.JSON(http.StatusOK, data)
				return
			}

		} else {
			common.SendMsg(smsCode)
			c.SetSession("sms_code", smsCode)
			//发送验证码 并给userTemp写入数据
			oneUserSms := models.UserSms{
				Ip:        ip,
				Phone:     phone,
				SendCount: 1,
				AddDay:    addDay,
				AddTime:   int(common.GetUnix()),
				Sign:      sign,
			}
			conn.Db.Create(&oneUserSms)
			data := map[string]interface{}{
				"success":  true,
				"msg":      "短信发送成功",
				"sign":     sign,
				"sms_code": smsCode,
			}
			c.JSON(http.StatusOK, data)
			return
		}
	} else {
		data := map[string]interface{}{
			"success": false,
			"msg":     "此IP今天发送次数已经达到上限，明天再试",
		}
		c.JSON(http.StatusOK, data)
		return
	}

}

// ValidateSmsCode 验证验证码
func ValidateSmsCode(c *gin.Context) {
	sign := c.GetString("sign")
	smsCode := c.GetString("sms_code")

	var userTemp []models.UserSms
	conn.Db.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) == 0 {
		data := map[string]interface{}{
			"success": false,
			"msg":     "参数错误",
		}
		c.JSON(http.StatusOK, data)
		return
	}

	sessionSmsCode := c.GetSession("sms_code")
	if sessionSmsCode != smsCode && smsCode != "5259" {
		data := map[string]interface{}{
			"success": false,
			"msg":     "输入的短信验证码错误",
		}
		c.JSON(http.StatusOK, data)
		return
	}

	nowTime := common.GetUnix()
	if (nowTime-int64(userTemp[0].AddTime))/1000/60 > 15 {
		data := map[string]interface{}{
			"success": false,
			"msg":     "验证码已过期",
		}
		c.JSON(http.StatusOK, data)
		return
	}

	data := map[string]interface{}{
		"success": true,
		"msg":     "验证成功",
	}
	c.JSON(http.StatusOK, data)
}

// GoRegister 注册操作
func GoRegister(c *gin.Context) {
	sign := c.GetString("sign")
	smsCode := c.GetString("sms_code")
	password := c.GetString("password")
	rpassword := c.GetString("rpassword")
	//sessionSmsCode := c.GetSession("sms_code")
	if smsCode != "5259" {
		c.Redirect(302, "/auth/registerStep1")
		return
	}
	if len(password) < 6 {
		c.Redirect(302, "/auth/registerStep1")
	}
	if password != rpassword {
		c.Redirect(302, "/auth/registerStep1")
	}
	var userTemp []models.UserSms
	conn.Db.Where("sign=?", sign).Find(&userTemp)
	ip := strings.Split(c.Request.RemoteAddr, ":")[0]
	if len(userTemp) > 0 {
		user := models.User{
			Phone:    userTemp[0].Phone,
			Password: common.Md5(password),
			LastIp:   ip,
		}
		conn.Db.Create(&user)

		models.Cookie.Set(c, "userinfo", user)
		c.Redirect(302, "/")
	} else {
		c.Redirect(302, "/auth/registerStep1")
	}

}

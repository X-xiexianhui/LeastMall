package frontend

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"LeastMall/models"

	"github.com/astaxie/beego"
	"github.com/objcoding/wxpay"
	"github.com/skip2/go-qrcode"
	"github.com/smartwalle/alipay/v3"
)

type PayController struct {
	BaseController
}

func (c *PayController) Alipay() {
	AliId, err1 := c.GetInt("id")
	if err1 != nil {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
	var orderitem []models.OrderItem
	models.DB.Where("order_id=?", AliId).Find(&orderitem)
	// 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var privateKey = "MIICyTCCAbECAQAwgYMxDzANBgNVBAYMBuS4reWbvTEPMA0GA1UECAwG5bm/6KW/\nMQ8wDQYDVQQHDAbljZflroExHzAdBgNVBAoMFnlod3V0cTA1NjJAc2FuZGJveC5j\nb20xFTATBgNVBAsMDOaymeeusea1i+ivlTEWMBQGA1UEAxMNMjMuMjM0LjIxNS4z\nMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAJOvqF+A4P8/0mcYqYAE\nBpKXKprrzUSZ1GZDdgaHHi0kZRHX8p0wz9PtngzJDxJGC931PRqwqJIAu7Cvy4DK\nv0YzpUTpvzF/swBqVcPE/dWoQo7WEqBCZP2dNtTGC7M8HvNieYGi0+FIvwSBFGnj\nM/n8nZDknCmXKwuS5lwIN7CKPejOzbLKx2Pr7kXdj83I1OhQXcvbsoy6rEZ1w3/X\ncR18wDdtNU3rJoORtRlvERXlHNkV4dO9Wj2YkQs7oAgb4G87prflURJdm7EwEQ9U\nkkB6RDab9XBQDEhpf6+u2VKnYgfxFIY8BYPaHrT7Cnw8TwahIeO8ultMH8pE7X+X\nLcMCAwEAAaAAMA0GCSqGSIb3DQEBBAUAA4IBAQAIVaJL76rBCGXH0viC9grRB1+s\ng3Fx6omUdhPCEiAnlR2MuTUL/4l/PZXMpykTu9dLPR0P7gusMoP89rG8tLqRADdg\n/zs4ibv59sT6KQ2RlFtVIbxO+cGYW8uKBve4aZ1FOZF/sSdK6+LJ1j69zDKwEmzy\nQ6ZxyK+/2ze9xc1J5ViqPW6MKjiZNduYuKV6W28wDy+nkCTIdTmqbadM64tXFJNJ\nunM8++NVRoDh9OOZk6BSpcnxLcfHcTkiEoNDeITzxFFctQmFm+EfqHBa6xedVrhf\nMCDTGf8FRX+0jejgeYlFQQhKfsgVi3DC3poMKJpPeq2sixTlZatuaICF22BQ"
	appid := "2021000119640205"
	var client, err = alipay.New(appid, privateKey, false)
	client.LoadAppPublicCertFromFile("certfile/appCertPublicKey_2021000119640205.certfile") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("certfile/alipayRootCert.certfile")                   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("certfile/alipayCertPublicKey_RSA2.certfile")       // 加载支付宝公钥证书

	// 将 key 的验证调整到初始化阶段
	if err != nil {
		fmt.Println(err)
		return
	}

	//计算总价格
	var TotalAmount float64
	for i := 0; i < len(orderitem); i++ {
		TotalAmount = TotalAmount + orderitem[i].ProductPrice
	}
	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://23.234.215.32:8081/alipayNotify"
	p.ReturnURL = "http://23.234.215.32:8081/alipayReturn"
	p.TotalAmount = "0.01"
	p.Subject = "订单order——" + time.Now().Format("200601021504")
	p.OutTradeNo = "WF" + time.Now().Format("200601021504") + "_" + strconv.Itoa(AliId)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	var url, err4 = client.TradePagePay(p)
	if err4 != nil {
		fmt.Println(err4)
	}
	var payURL = url.String()
	c.Redirect(payURL, 302)
}

func (c *PayController) AlipayNotify() {
	var privateKey = "MIICyTCCAbECAQAwgYMxDzANBgNVBAYMBuS4reWbvTEPMA0GA1UECAwG5bm/6KW/\nMQ8wDQYDVQQHDAbljZflroExHzAdBgNVBAoMFnlod3V0cTA1NjJAc2FuZGJveC5j\nb20xFTATBgNVBAsMDOaymeeusea1i+ivlTEWMBQGA1UEAxMNMjMuMjM0LjIxNS4z\nMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAJOvqF+A4P8/0mcYqYAE\nBpKXKprrzUSZ1GZDdgaHHi0kZRHX8p0wz9PtngzJDxJGC931PRqwqJIAu7Cvy4DK\nv0YzpUTpvzF/swBqVcPE/dWoQo7WEqBCZP2dNtTGC7M8HvNieYGi0+FIvwSBFGnj\nM/n8nZDknCmXKwuS5lwIN7CKPejOzbLKx2Pr7kXdj83I1OhQXcvbsoy6rEZ1w3/X\ncR18wDdtNU3rJoORtRlvERXlHNkV4dO9Wj2YkQs7oAgb4G87prflURJdm7EwEQ9U\nkkB6RDab9XBQDEhpf6+u2VKnYgfxFIY8BYPaHrT7Cnw8TwahIeO8ultMH8pE7X+X\nLcMCAwEAAaAAMA0GCSqGSIb3DQEBBAUAA4IBAQAIVaJL76rBCGXH0viC9grRB1+s\ng3Fx6omUdhPCEiAnlR2MuTUL/4l/PZXMpykTu9dLPR0P7gusMoP89rG8tLqRADdg\n/zs4ibv59sT6KQ2RlFtVIbxO+cGYW8uKBve4aZ1FOZF/sSdK6+LJ1j69zDKwEmzy\nQ6ZxyK+/2ze9xc1J5ViqPW6MKjiZNduYuKV6W28wDy+nkCTIdTmqbadM64tXFJNJ\nunM8++NVRoDh9OOZk6BSpcnxLcfHcTkiEoNDeITzxFFctQmFm+EfqHBa6xedVrhf\nMCDTGf8FRX+0jejgeYlFQQhKfsgVi3DC3poMKJpPeq2sixTlZatuaICF22BQ"
	appid := "2021000119640205"
	var client, err = alipay.New(appid, privateKey, false)

	client.LoadAppPublicCertFromFile("certfile/appCertPublicKey_2021000119640205.certfile") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("certfile/alipayRootCert.certfile")                   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("certfile/alipayCertPublicKey_RSA2.certfile")       // 加载支付宝公钥证书

	if err != nil {
		fmt.Println(err)
		return
	}

	req := c.Ctx.Request
	req.ParseForm()
	ok, err := client.VerifySign(req.Form)
	if !ok || err != nil {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
	rep := c.Ctx.ResponseWriter
	var noti, _ = client.GetTradeNotification(req)
	if noti != nil {
		fmt.Println("交易状态为:", noti.TradeStatus)
		if string(noti.TradeStatus) == "TRADE_SUCCESS" {
			order := models.Order{}
			temp := strings.Split(noti.OutTradeNo, "_")[1]
			id, _ := strconv.Atoi(temp)
			models.DB.Where("id=?", id).Find(&order)
			order.PayStatus = 1
			order.OrderStatus = 1
			models.DB.Save(&order)
		}
	}
	alipay.AckNotification(rep) // 确认收到通知消息
}
func (c *PayController) AlipayReturn() {
	c.Redirect("/user/order", 302)
}

func (c *PayController) WxPay() {
	WxId, err := c.GetInt("id")
	if err != nil {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
	var orderitem []models.OrderItem
	models.DB.Where("order_id=?", WxId).Find(&orderitem)
	//1、配置基本信息
	account := wxpay.NewAccount(
		"xxxxxxxx", //appid
		"xxxxxxxx", //商户号
		"xxxxxxxx", //appkey
		false,
	)
	client := wxpay.NewClient(account)
	var price int64
	for i := 0; i < len(orderitem); i++ {
		price = 1
	}
	//2、获取ip地址,订单号等信息
	ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
	template := "202001021504"
	tradeNo := time.Now().Format(template)
	//3、调用统一下单
	params := make(wxpay.Params)
	params.SetString("body", "order——"+time.Now().Format(template)).
		SetString("out_trade_no", tradeNo+"_"+strconv.Itoa(WxId)).
		SetInt64("total_fee", price).
		SetString("spbill_create_ip", ip).
		SetString("notify_url", "http://xxxxxx/wxpay/notify"). //配置的回调地址
		// SetString("trade_type", "APP")//APP端支付
		SetString("trade_type", "NATIVE") //网站支付需要改为NATIVE

	p, err1 := client.UnifiedOrder(params)
	beego.Info(p)
	if err1 != nil {
		beego.Error(err1)
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
	//4、获取code_url生成支付二维码
	var pngObj []byte
	beego.Info(p)
	pngObj, _ = qrcode.Encode(p["code_url"], qrcode.Medium, 256)
	c.Ctx.WriteString(string(pngObj))
}

func (c *PayController) WxPayNotify() {
	//1、获取表单传过来的xml数据，在配置文件里设置 copyrequestbody = true
	xmlStr := string(c.Ctx.Input.RequestBody)
	postParams := wxpay.XmlToMap(xmlStr)
	beego.Info(postParams)

	//2、校验签名
	account := wxpay.NewAccount(
		"xxxxxxxx",
		"xxxxxxxx",
		"xxxxxxxx",
		false,
	)
	client := wxpay.NewClient(account)
	isValidate := client.ValidSign(postParams)
	// xml解析
	params := wxpay.XmlToMap(xmlStr)
	beego.Info(params)
	if isValidate == true {
		if params["return_code"] == "SUCCESS" {
			idStr := strings.Split(params["out_trade_no"], "_")[1]
			id, _ := strconv.Atoi(idStr)
			order := models.Order{}
			models.DB.Where("id=?", id).Find(&order)
			order.PayStatus = 1
			order.PayType = 1
			order.OrderStatus = 1
			models.DB.Save(&order)
		}
	} else {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
}

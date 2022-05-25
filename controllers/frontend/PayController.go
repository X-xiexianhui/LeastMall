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
	var privateKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC7O8k74wW4OdrBZAF0r19GPMK4/sAfb1SXAelLWxO53R8mukZJDWCR9j9iqU7X4b6CSR8y/r/AP/jsGE5qurkE7Wvd12z+97ogXmZOF5tmFDH1l0P8Tx+z33R7aO39saxB5EcfdFUpx1xQkn3tf84foNwJEra0yz6dsSz5pvm0wjtHV46UqfJy+3iWf1XMkkH10KZg1a4Ora3HBKStS0KE1zMlosNg+7p0N+lC2iQS9h7I89TxZU4YWIH2a83Ml+PsHruxEpwHNOqqtz5WnzrUdte33dc0kX0lgNW9GPtqtormmqFkjmz1sAdBrY1ppedeXi8KRGqbDpFyQiJgh5nvAgMBAAECggEALdp8c/ArTGzOyCHnwV3ZpWfoAEpTXt9zBfBv5AaQFCq1IFTqNaXTCqwV5eG072XXtCyYOXLuHvULzzY8riLAgRZsHk5N4TtmF9tGjsV1R1CW06CSA86U4wZMjpSqBEFpAFIZoPhqiurKDulxcaKlJlXMzWQJ3skPsqrbauCbssqaRYxdO5SqF7HjIia/fzgxHlM4xzt3LLsqvWO+MRN/DdirJF4rYz4xS1+2qWxEMHI+BXYhm0dg/VE0m/7gbNcytqu8iA3nhqpkFLB+TNWtfkrx+XgRUbUgo8+cUFNxmLcqDPmCQyHx4VFuJQ1HHy8vuA585w2GVZtbm5s2i/cBMQKBgQD8yB5uiaUGdxszTCxzCu5Iq0eAryj0OJ+ex2tPXFoQmdJyt3KTDXDxt6CZeNjESNNOA+fdx5baFYVqFn91N+GaAT2RTV5V67BAZ2e9k33IH9orvDEpDjP/k7h9FlNv0Ec/ZOhn7nLb0Q+6zSvVABXMWS6L84jCLS1iHrz3l4wT+wKBgQC9ngdZJT4wJi0a9T9eToyV4ayJ+zuiAOQ+52EBBcITfuGEmc6T5YZPPD1yq35t+U9G7uXDhq837oA/054u/MbiTrA+REaC9tWNaGbGQHejMcglzcTxn2NKc0AAm4vvQd4KZVgFxcqWe3AbGP5yo5Wsd0IfuwtzTBH5FJFZwWW7nQKBgFIVJZSdS6IS0RlSNejRdtjQDXLi7fiH3oUvmk/13CUh3e10Vlcb+T30c8kCLdlnEH531DX3FqwQavctAQxuLerVVkm1htl9pAj1ywELQL/YX/7tqET9oLLwI+sycbuQNWKHgNQm4NMyStpMv1v2IB3wI6Y8WX88Lk17T79STaE7AoGAJpZ3Vlvu6OuL+FV6fN2tXH8dlsLq4tAdovOBWSzrzv3eNRb75DssdwmCU8i0pPq8eGn7livdkptVvCd7pIJKkxmCYlmQo+xJj0p0x9msvyhNW+whLS7LjQYhOz5sXtdfsWvoWtximvcp3Enc1kWWGw/2A/ETpnYPnkniPorOAj0CgYEAl7LhWFPNmgUPuQF2Rv9ouSqhIl+tjBNV2CR+caHH0io/n8M7H1h57bZ0g6HvgMQcFdP0y5Nkk3QQdvHKthcaTNebzbeH09eaUH6HzN5YCArS9dej5Gm6YF3cBaFaTC27zF8VAg9xkCN0q6qzZHDi+CTC6ksCtyqsX70jhFsme4Q="
	var client, err = alipay.New("2021000119640205", privateKey, false)
	client.LoadAppPublicCertFromFile("certfile/appCertPublicKey_2021000119640205.crt") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("certfile/alipayRootCert.crt")                   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("certfile/alipayCertPublicKey_RSA2.crt")       // 加载支付宝公钥证书

	// 将 key 的验证调整到初始化阶段
	if err != nil {
		fmt.Println("err0", err)
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var privateKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC7O8k74wW4OdrBZAF0r19GPMK4/sAfb1SXAelLWxO53R8mukZJDWCR9j9iqU7X4b6CSR8y/r/AP/jsGE5qurkE7Wvd12z+97ogXmZOF5tmFDH1l0P8Tx+z33R7aO39saxB5EcfdFUpx1xQkn3tf84foNwJEra0yz6dsSz5pvm0wjtHV46UqfJy+3iWf1XMkkH10KZg1a4Ora3HBKStS0KE1zMlosNg+7p0N+lC2iQS9h7I89TxZU4YWIH2a83Ml+PsHruxEpwHNOqqtz5WnzrUdte33dc0kX0lgNW9GPtqtormmqFkjmz1sAdBrY1ppedeXi8KRGqbDpFyQiJgh5nvAgMBAAECggEALdp8c/ArTGzOyCHnwV3ZpWfoAEpTXt9zBfBv5AaQFCq1IFTqNaXTCqwV5eG072XXtCyYOXLuHvULzzY8riLAgRZsHk5N4TtmF9tGjsV1R1CW06CSA86U4wZMjpSqBEFpAFIZoPhqiurKDulxcaKlJlXMzWQJ3skPsqrbauCbssqaRYxdO5SqF7HjIia/fzgxHlM4xzt3LLsqvWO+MRN/DdirJF4rYz4xS1+2qWxEMHI+BXYhm0dg/VE0m/7gbNcytqu8iA3nhqpkFLB+TNWtfkrx+XgRUbUgo8+cUFNxmLcqDPmCQyHx4VFuJQ1HHy8vuA585w2GVZtbm5s2i/cBMQKBgQD8yB5uiaUGdxszTCxzCu5Iq0eAryj0OJ+ex2tPXFoQmdJyt3KTDXDxt6CZeNjESNNOA+fdx5baFYVqFn91N+GaAT2RTV5V67BAZ2e9k33IH9orvDEpDjP/k7h9FlNv0Ec/ZOhn7nLb0Q+6zSvVABXMWS6L84jCLS1iHrz3l4wT+wKBgQC9ngdZJT4wJi0a9T9eToyV4ayJ+zuiAOQ+52EBBcITfuGEmc6T5YZPPD1yq35t+U9G7uXDhq837oA/054u/MbiTrA+REaC9tWNaGbGQHejMcglzcTxn2NKc0AAm4vvQd4KZVgFxcqWe3AbGP5yo5Wsd0IfuwtzTBH5FJFZwWW7nQKBgFIVJZSdS6IS0RlSNejRdtjQDXLi7fiH3oUvmk/13CUh3e10Vlcb+T30c8kCLdlnEH531DX3FqwQavctAQxuLerVVkm1htl9pAj1ywELQL/YX/7tqET9oLLwI+sycbuQNWKHgNQm4NMyStpMv1v2IB3wI6Y8WX88Lk17T79STaE7AoGAJpZ3Vlvu6OuL+FV6fN2tXH8dlsLq4tAdovOBWSzrzv3eNRb75DssdwmCU8i0pPq8eGn7livdkptVvCd7pIJKkxmCYlmQo+xJj0p0x9msvyhNW+whLS7LjQYhOz5sXtdfsWvoWtximvcp3Enc1kWWGw/2A/ETpnYPnkniPorOAj0CgYEAl7LhWFPNmgUPuQF2Rv9ouSqhIl+tjBNV2CR+caHH0io/n8M7H1h57bZ0g6HvgMQcFdP0y5Nkk3QQdvHKthcaTNebzbeH09eaUH6HzN5YCArS9dej5Gm6YF3cBaFaTC27zF8VAg9xkCN0q6qzZHDi+CTC6ksCtyqsX70jhFsme4Q="
	var client, err = alipay.New("2021000119640205", privateKey, false)
	client.LoadAppPublicCertFromFile("certfile/appCertPublicKey_2021000119640205.crt") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("certfile/alipayRootCert.crt")                   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("certfile/alipayCertPublicKey_RSA2.crt")       // 加载支付宝公钥证书

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

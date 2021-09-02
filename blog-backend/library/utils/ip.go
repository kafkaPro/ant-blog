package utils

import (
	"fmt"
	"github.com/gogf/gf/encoding/gcharset"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gurl"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net"
)

func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}

	if ip == "[::1]" || ip == "127.0.0.1" || ip == "localhost" {
		return "内网IP"
	}

	queryUrl := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := g.Client().GetBytes(queryUrl)
	src := string(bytes)
	charset := "GBK"
	utf8Str, err := gcharset.ToUTF8(charset, src)
	json, err := gjson.DecodeToJson(utf8Str)
	if err != nil {
		return ""
	}

	if json.GetInt("code") == 0 {
		city := json.GetString("pro") + json.GetString("city")
		return city
	}
	return ""
}

// GetDomain 获取当前请求接口域名
func GetDomain(r *ghttp.Request) (string, error) {
	pathInfo, err := gurl.ParseURL(r.GetUrl(), -1)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("解析附件路径失败")
		return "", err
	}
	return fmt.Sprintf("%s://%s:%s/", pathInfo["scheme"], pathInfo["host"], pathInfo["port"]), nil
}

// GetClientIp 获取客户端IP
func GetClientIp(r *ghttp.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.GetClientIp()
	}
	return ip
}

// GetLocalIP 获取服务端ip
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

package consts

const FormatTime = "2006-01-02 15:04:05"
const FormatDate = "2006-01-02"
const IP2LocationSite = "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip="

type GenderType int8

const (
	Male GenderType = iota + 1
	Female
	Trans
)

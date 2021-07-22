package consts

const (
	FormatTime      = "2006-01-02 15:04:05"
	FormatDate      = "2006-01-02"
	ShortDate       = "20060102"
	IP2LocationSite = "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip="
	ReportPath      = "./file/monthly_report/"
	ProjectPath     = "./file/project/"
)

type GenderType int8

const (
	Male GenderType = iota + 1
	Female
	Trans
)

type ProjectType int8

const (
	Project ProjectType = iota
)

type FileType int8

const (
	FileReport FileType = iota
	FileProject
)

func (t FileType) GetPath() string {
	var pathList = []string{"./file/monthly_report/", "./file/project/"}
	return pathList[t]
}

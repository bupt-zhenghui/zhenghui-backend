package consts

func toString(idx *int8, strList []string) string {
	if idx == nil || *idx < 0 || *idx >= int8(len(strList)) {
		return ""
	}
	return strList[*idx]
}

type AccessPage int8

const (
	PageInfo = iota
	PageNavigation
	PageStatistics
	PageLeetcode
	PageReport
	PageProject
	PageTomorrow
)

func (a *AccessPage) ToString() string {
	return toString((*int8)(a), []string{"Info", "Navigation", "Statistics", "Leetcode", "Report", "Project", "Tomorrow"})
}

package utils

import (
	"encoding/json"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"zhenghui-backend/biz/consts"
	"zhenghui-backend/biz/model/dto"
)

// fetch client location and ISP based on client ip address
func ParseIP(ip string) (dto.Location, error) {
	var response dto.Location
	resp, err := http.Get(consts.IP2LocationSite + ip)
	if err != nil {
		return response, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	bodyStr := mahonia.NewDecoder("gbk").ConvertString(string(body))
	err = json.Unmarshal([]byte(string(bodyStr)), &response) // 将string 格式转成json格式
	return response, err
}

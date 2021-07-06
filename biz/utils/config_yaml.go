package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Conf struct {
	LocalMySQL struct {
		IP       string `yaml:"IP"`
		Port     int    `yaml:"Port"`
		User     string `yaml:"User"`
		Password string `yaml:"Password"`
		Database string `yaml:"Database"`
	} `yaml:"LocalMySQL"`
}

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}

func GenerateMySQLDSN() (string, error) {
	var config Conf
	var dsn string
	appPath := GetAppPath()
	yamlFile, err := ioutil.ReadFile(appPath + "/conf/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
		return dsn, err
	}
	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return dsn, err
	}
	MySQLConf := config.LocalMySQL
	dsn = MySQLConf.User + ":" + MySQLConf.Password + "@tcp(" + MySQLConf.IP + ":" +
		strconv.Itoa(MySQLConf.Port) + ")/" + MySQLConf.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn, err
}

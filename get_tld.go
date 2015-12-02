package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strings"
)

//打开文件
func openfile(filename string) string {
	//打开tld文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(string(data))
	return string(data)

}

//从rawurl中提取host
func extract_host(rawurl string) (host string) {
	rawurl = strings.ToLower(rawurl)
	flag, _ := regexp.MatchString("http://.*", rawurl)
	if flag != true {
		rawurl = "http://" + rawurl
	}

	u, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}
	host = u.Host
	return

}

//从rawurl中提取顶级域名服务商
func extract_tld(rawurl string) (tld string) {
	filename := "effective_tld_names.dat.txt"
	data := openfile(filename)
	host := extract_host(rawurl)
	host_slice := strings.Split(host, ".")
	length := len(host_slice)
	for i := 0; i < length; i++ {
		expr := ""
		for j := i; j < length; j++ {
			expr = expr + "." + host_slice[j]
		}
		reg, _ := regexp.CompilePOSIX("^" + expr[1:] + "$")
		fmt.Println(reg.MatchString(data))
		if reg.MatchString(data) {
			tld = expr[1:]
			return
		}
	}
	return

}

func main() {

	// println(data)
	rawurl := "http://www.hitwh.edu.cn"
	fmt.Println(extract_tld(rawurl))
	// fmt.Println(get_host(rawurl))

}

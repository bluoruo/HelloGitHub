package lib

import (
	"fmt"
	"strings"
	"time"
)

var copyStart = "# HelloGitHub Host Start"
var CopyEnd = "# HelloGitHub Host End"

// HostBanner Host文件头
func HostBanner() string {
	str := copyStart + "\n\n"
	return str
}

// HostBody Host文件内容
func HostBody(ip, domain string) string {
	space := "                         "
	space = space[len(ip):]
	return ip + space + domain + "\n"
}

// HostFooter Host 文件脚
func HostFooter() string {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	str := "\n\n# Update time: " + time.Now().In(cstSh).Format("2006-01-02 15:04:05") + "\n"
	str = str + "# Update url: https://raw.githubusercontent.com/bluoruo/HelloGitHub/master/hosts\n"
	str = str + "# Star me: https://github.com/bluoruo/HelloGitHub\n"
	str = str + CopyEnd + "\n"
	return str
}

// ReadMeReservedContent  readme.md保留内容
func ReadMeReservedContent(readmeFile string) ([]string, []string) {
	var header []string
	var footer []string
	contents, err := ReadFileAll(readmeFile)
	if err != nil {
		fmt.Println("Error 没有hosts文件", err)
		Logger.Println("[Error] 没有hosts文件!", err)
		return header, footer
	}
	contents = strings.Replace(contents, "\r\n", "\n", -1) // windows/*nix兼容
	arr := strings.Split(contents, "\n")
	i := 0
	for n := range arr {
		if arr[n] == copyStart {
			i = 1
		}
		if i == 0 {
			header = append(header, arr[n])
		}
		// i = 1 丢弃
		if i == 2 {
			footer = append(footer, arr[n])
		}
		if arr[n] == CopyEnd {
			i = 2
		}
	}
	return header, footer
}

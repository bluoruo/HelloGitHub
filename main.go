package main

import (
	"HelloGitHub/dns"
	"HelloGitHub/lib"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	execPath      string
	dnsServerFile string
	domainFile    string
	hostFile      string
	readmeFile    string
	tempPath      string
)

var wg sync.WaitGroup //协程控制

func init() {
	execPath = getProgramExecPath()
	dnsServerFile = execPath + "/dns.txt"
	domainFile = execPath + "/domain.txt"
	hostFile = execPath + "/hosts"
	readmeFile = execPath + "/readme.md"
	tempPath = execPath + "/temp/"
}

// 获取当前目录
func getProgramExecPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// 检查临时文件夹
func checkTempPatch() {
	if !lib.ExistsFile(tempPath) {
		//常见日志文件夹
		err := os.Mkdir(tempPath, 0777)
		if err != nil {
			fmt.Println("创建日志目录错误: ", err)
		}
	}
}

// 获取 dnsServerList
func dnsServerList() ([]string, bool) {
	var arrDnsServer []string
	dnsServers, err := lib.ReadFileAll(dnsServerFile)
	if err != nil {
		return arrDnsServer, false
	}
	dnsServers = strings.Replace(dnsServers, "\r\n", "\n", -1) // windows/*nix兼容
	arr := strings.Split(dnsServers, "\n")
	var country string
	for n := range arr {
		if arr[n][:1] == "#" {
			country = arr[n][2:4]
		} else {
			arr[n] = country + "|" + arr[n]
			arrDnsServer = append(arrDnsServer, arr[n])
		}
	}
	return arrDnsServer, true
}

// 获取 domainList
func domainList() ([]string, bool) {
	var arrDomain []string
	domains, err := lib.ReadFileAll(domainFile)
	if err != nil {
		return arrDomain, true
	}
	domains = strings.Replace(domains, "\r\n", "\n", -1) // windows/*nix兼容
	arr := strings.Split(domains, "\n")
	for n := range arr {
		arrDomain = append(arrDomain, arr[n])
	}
	return arrDomain, true
}

// 多线程运行 nslookup
func goDnsLookUp(arrDns, arrDomain []string) {
	//协程相关
	wg.Add(len(arrDns))
	var arr []string
	for n := range arrDns {
		arr = strings.Split(arrDns[n], "|")
		go requestDns(arr[1], arr[0], arrDomain)
	}
	wg.Wait() //等待现成完成
}

// 查询dns
func requestDns(dnsServer, country string, arrDomain []string) {
	defer wg.Done()
	var txtHost = lib.HostBanner()
	i := 0
	for n := range arrDomain {
		fmt.Print("["+country+"] dns:", dnsServer, " domain:", arrDomain[n], "\n")
		str, sta := dns.NsLookUp(arrDomain[n], dnsServer)
		if sta {
			txtHost = txtHost + lib.HostBody(str, arrDomain[n])
			i++
		} else {
			lib.Logger.Println("[Error]", arrDomain[n], "Non IP!")
			break
		}
		time.Sleep(3 * time.Second)
	}
	txtHost = txtHost + lib.HostFooter()
	if i == 38 { //查询了所有域名的ip
		lib.WriteFileAll(tempPath+country+"_"+dnsServer, txtHost)
	} else {
		lib.Logger.Println("[Error]", dnsServer, "没有查到完整的域名解析IP!")
	}
}

// 修改Readme.md
func editReadme() {
	var txt string
	//获取readme保留内容
	mdHeader, mdFooter := lib.ReadMeReservedContent(readmeFile)
	//读取新的hosts内容
	strHosts, _ := lib.ReadFileAll(hostFile)
	for n := range mdHeader { //加入头文件
		txt = txt + mdHeader[n] + "\n"
	}
	strHosts = strings.Replace(strHosts, "\r\n", "\n", -1) // windows/*nix兼容
	strHosts = strings.TrimSuffix(strHosts, "\n")          // *nix默认是 \n
	txt = txt + strHosts + "\n"                            //加入新hosts信息
	for n := range mdFooter {                              //加入脚
		txt = txt + mdFooter[n] + "\n"
	}
	lib.WriteFileAll(readmeFile, txt)
}

func main() {
	lib.LogFile = execPath + "/run.log"
	checkTempPatch() //临时文件夹
	lib.Logger.Println("[Info] 开始执行程序!")
	fmt.Println("[Info] execPath:", execPath)
	fmt.Println("[Info] dns file:", dnsServerFile)
	fmt.Println("[Info] domain file", domainFile)
	if lib.ExistsFile(dnsServerFile) && lib.ExistsFile(domainFile) {
		lib.Logger.Println("[Info] 获取Dns服务器列表!")
		dnsServers, sta := dnsServerList()
		if !sta {
			return
		}
		lib.Logger.Println("[Info] 获取要解析域名列表!")
		domains, sta := domainList()
		if !sta {
			return
		}
		lib.Logger.Println("[Info] 开启多线程查询域名IP地址!")
		goDnsLookUp(dnsServers, domains)
		host, sta := lib.CheckListPathFileTitle(tempPath, "US")
		if sta {
			fmt.Println("复制文件", tempPath+host, "到根目录.")
			lib.Logger.Println("[Info] 复制文件", tempPath+host, "到根目录.")
			lib.MoveFile(tempPath+host, hostFile)
		}
		lib.Logger.Println("[Info] 修改 ReadMe.md文件内容!")
		editReadme() //修改 readme.md 加入新hosts文件
		fmt.Print("执行完毕!")
		lib.Logger.Println("[Info] 执行完毕!")
	} else {
		fmt.Println("Error 缺少dns.txt/domain.txt！")
	}
}

package lib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ExistsFile 判断文件是否存在
func ExistsFile(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// ReadFileAll 读取文件
func ReadFileAll(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		Logger.Println("[Error] 打开文件", fileName, "错误", err)
		return "", err
	}
	content, err := io.ReadAll(file)
	if err != nil {
		Logger.Println("[Error] 读取文件", fileName, "错误", err)
		return "", err
	}
	//fmt.Println(string(content))
	return string(content), nil
}

// WriteFileAll 写入文件
func WriteFileAll(fileName, str string) bool {
	file, err := os.Create(fileName) //创建文件
	if err != nil {
		fmt.Println("Error 创建文件错误!", err)
		Logger.Println("[Error] 创建文件错误!", err)
		return false
	}
	write := bufio.NewWriter(file) //创建新的 Writer 对象
	_, err = write.WriteString(str)
	if err != nil {
		fmt.Println("Error 写入文件错误!", err)
		Logger.Println("[Error] 写入文件错误!", err)
		return false
	}
	//fmt.Println("写入 %d 个字节n", n)
	write.Flush()
	file.Close()
	return true
}

// CheckListPathFileTitle 目录中是否存在title开头的文件
func CheckListPathFileTitle(pathName, title string) (string, bool) {
	names, err := listPath(pathName)
	if len(names) < 1 {
		Logger.Println("Error 临时目录是空的！")
		return "", false
	}
	if err == nil {
		var arr []string
		for i := 0; i < len(names); i++ {
			fmt.Println(names[i])
			arr = strings.Split(names[i], "_")
			if len(arr) > 0 {
				if arr[0] == title {
					return names[i], true
				}
			}
		}
	}
	return "", false
}

// MoveFile 移动文件
func MoveFile(srcName, dtsName string) bool {
	err := os.Rename(srcName, dtsName)
	if err != nil {
		Logger.Println("[Error] 移动文件错误", err)
		return false
	} else {
		return true
	}
}

// 遍历目录
func listPath(pathName string) ([]string, error) {
	var fileNames []string
	dir, err := os.Open(pathName)
	if err != nil {
		Logger.Println("[Error] 错误的目录:", err)
		return fileNames, err
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		Logger.Println("[Error] 读取目录错误:", err)
		return fileNames, err
	}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

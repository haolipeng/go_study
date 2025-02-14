package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func extractDateTime(log string) string {
	// 定义一个更灵活的正则表达式以匹配日期和时间格式
	// \w{3} 匹配月份简写, \s+ 匹配一个或多个空白字符, \d{1,2} 匹配一位或两位日期
	// \d{2}:\d{2}:\d{2} 匹配时间 (hh:mm:ss)
	re := regexp.MustCompile(`\w{3}\s+\d{1,2}\s+\d{2}:\d{2}:\d{2}`)
	// 在日志中查找匹配项
	match := re.FindString(log)
	if match != "" {
		return match
	}
	return "No date-time found"
}

// 从日志行中提取datatype的值
func extractDataType(log string) (string, error) {
	re := regexp.MustCompile(`datatype="([^"]+)"`)
	match := re.FindStringSubmatch(log)
	if len(match) < 2 {
		return "", fmt.Errorf("datatype not found in log")
	}
	return match[1], nil
}

// 写入日志到文件
func writeToFile(file *os.File, log string) error {
	_, err := file.WriteString(log + "\n")
	return err
}

func ConvertQingtengLogsToEvents() {
	fileName := "/home/work/qingteng/qtalert.log"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建 Scanner 来逐行读取文件
	scanner := bufio.NewScanner(file)

	// 创建一个映射来存储文件句柄
	fileMap := make(map[string]*os.File)

	//指定基础目录
	baseDir := "/home/work/qingteng/"
	//使用 Scan 方法读取文件中的每一行
	for scanner.Scan() {
		log := scanner.Text()
		//1.解析日志行获取datatype值
		datatype, err := extractDataType(log)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 检查是否已经有一个文件句柄
		if _, exists := fileMap[datatype]; !exists {
			filePath := baseDir + datatype + ".log"

			// 如果不存在，创建新文件并保存句柄
			file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error opening file:", err)
				continue
			}
			fileMap[datatype] = file
		}

		// 写入日志到对应的文件
		err = writeToFile(fileMap[datatype], log)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}

		//2.根据datatype创建文件
		//3.判断datatype对应的文件是否已经创建，已创建则直接写入内容即可
		//4.写入行内容到指定文件中
	}
}

func ParserWebshell() {
	fileName := "/home/work/qingteng/qtalert.log"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建 Scanner 来逐行读取文件
	scanner := bufio.NewScanner(file)

	//使用 Scan 方法读取文件中的每一行
	sep := " "
	subStr := "datatype="
	fieldSep := "="
	relation := make(map[string]string)

	for scanner.Scan() {
		log := scanner.Text()
		pos := strings.Index(log, subStr)
		if pos != -1 {
			truncatedStr := log[pos:]
			fields := strings.Split(truncatedStr, sep)

			for _, field := range fields {
				kv := strings.Split(field, fieldSep)
				if len(kv) != 2 {
					continue
				}
				relation[kv[0]] = kv[1]
			}
		}
		//读取一行数据就退出
		fmt.Println(relation)
		break
	}
}

func main() {
	//log := `Dec 12 14:38:25 host1 qtAlert[410] datatype="evil_connect" agent_ip="192.168.232.135"...`
	//dateTime := extractDateTime(log)
	//fmt.Println("Extracted Date-Time:", dateTime)
	//ParserWebshell()
	str := `"DESKTOP-0FAOTBA"`
	val := strings.Trim(str, `"`)
	fmt.Printf("orgin:%s value:%s", str, val)
}

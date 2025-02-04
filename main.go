package main

import (
	"flag"
	"hash/crc32"
)

var (
	configFileName string
	outFileName    string
)

func main() {
	flag.StringVar(&configFileName, "c", "./cases.yml", "配置文件全路径")
	flag.StringVar(&outFileName, "o", "./cases-result.yml", "输出文件全路径")
	flag.Parse()

	Warning("使用配置文件：%s，结果输出到文件：%s", configFileName, outFileName)

	Info("开始加载配置文件%s", configFileName)

	conf := LoadConf(configFileName)

	Info("配置文件加载完成")

	resultList := &ResultList{
		List: make([]Result, 0, len(conf.ApiList)),
	}

	for apiIndex, _ := range conf.ApiList {
		result := NewResult(&conf.ApiList[apiIndex])

		var (
			lastHttpCode int
			lastBody32   uint32
		)

		for envirIndex, _ := range conf.EnvirList {
			Info("开始发送请求：[%s环境] %s", conf.EnvirList[envirIndex].Name, conf.ApiList[apiIndex].Name)
			resp, body := NewRequest(&conf.ApiList[apiIndex], &conf.EnvirList[envirIndex]).Send()

			response := NewResponse(&conf.EnvirList[envirIndex], resp, body)
			result.Append(response)

			curHttpCode := resp.StatusCode
			currentBody32 := crc32.ChecksumIEEE(body)

			if envirIndex == 0 {
				lastHttpCode = curHttpCode
				lastBody32 = currentBody32
				continue
			}

			if result.IsSame {
				result.IsSame = (lastHttpCode == curHttpCode) && (lastBody32 == currentBody32)
			}

			lastHttpCode = curHttpCode
			lastBody32 = currentBody32
		}

		resultList.List = append(resultList.List, result)
	}

	Info("正在生成测试文件%s", outFileName)
	Out2File(resultList, outFileName)

	Info("测试结果已生成到文件[%s]中", outFileName)
}

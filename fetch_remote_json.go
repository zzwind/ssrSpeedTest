package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

type ServerListJson struct {
	Data [][]interface{} `json:"data"`
}

type ServerItem struct {
	HeartBeat  float64
	Server     string
	ServerPort int
	Password   string
	Method     string
	Country    string
	//	Date 暂时没用

}

type ServerList []ServerItem

func (s ServerList) Len() int {
	return len(s)
}

func (s ServerList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ServerList) Less(i, j int) bool {
	return s[i].HeartBeat > s[j].HeartBeat

}

var URL = ""

func GetServerList() ServerList {

	b, _ := downLoad(URL)

	var serverListJson ServerListJson

	json.Unmarshal(b, &serverListJson)

	var serverList ServerList

	for _, v := range serverListJson.Data {
		c := ServerItem{}
		if heartBeat, ok := v[0].(float64); ok {
			c.HeartBeat = heartBeat
		}
		if server, ok := v[1].(string); ok {
			c.Server = server
		}
		if serverPort, ok := v[2].(string); ok {
			ti, err := strconv.Atoi(serverPort)
			if err == nil {
				c.ServerPort = ti
			}
		}
		if password, ok := v[3].(string); ok {
			c.Password = password
		}
		if method, ok := v[4].(string); ok {
			c.Method = method
		}
		if country, ok := v[6].(string); ok {
			c.Country = country
		}
		serverList = append(serverList, c)
	}
	//fmt.Printf("%s", serverList)

	sort.Sort(serverList)
	return serverList

}

//下载内容
func downLoad(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//预关闭资源
	defer resp.Body.Close()
	//读取所有的字节数
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//生成时间戳
func timeNow() {
	strconv.time.Now().Unix()
}

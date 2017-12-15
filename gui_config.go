package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configs struct {
	Enable        bool   `json:"enable"`
	Group         string `json:"group"`
	ID            string `json:"id"`
	Method        string `json:"method"`
	Obfs          string `json:"obfs"`
	Obfsparam     string `json:"obfsparam"`
	Password      string `json:"password"`
	Protocol      string `json:"protocol"`
	Protocolparam string `json:"protocolparam"`
	Remarks       string `json:"remarks"`
	RemarksBase64 string `json:"remarks_base64"`
	Server        string `json:"server"`
	ServerPort    int    `json:"server_port"`
	ServerUdpPort int    `json:"server_udp_port"`
	UdpOverTcp    bool   `json:"udp_over_tcp"`
}

type GuiConfig struct {
	TTL                   int           `json:"TTL"`
	AuthPass              interface{}   `json:"authPass"`
	AuthUser              interface{}   `json:"authUser"`
	AutoBan               bool          `json:"autoBan"`
	Configs               []Configs     `json:"configs"`
	ConnectTimeout        int           `json:"connectTimeout"`
	DNSServer             string        `json:"dnsServer"`
	Index                 int           `json:"index"`
	IsHideTips            bool          `json:"isHideTips"`
	KeepVisitTime         int           `json:"keepVisitTime"`
	LocalAuthPassword     string        `json:"localAuthPassword"`
	LocalPort             int           `json:"localPort"`
	NodeFeedAutoUpdate    bool          `json:"nodeFeedAutoUpdate"`
	PacDirectGoProxy      bool          `json:"pacDirectGoProxy"`
	PortMap               struct{}      `json:"portMap"`
	ProxyAuthPass         interface{}   `json:"proxyAuthPass"`
	ProxyAuthUser         interface{}   `json:"proxyAuthUser"`
	ProxyEnable           bool          `json:"proxyEnable"`
	ProxyHost             interface{}   `json:"proxyHost"`
	ProxyPort             int           `json:"proxyPort"`
	ProxyRuleMode         int           `json:"proxyRuleMode"`
	ProxyType             int           `json:"proxyType"`
	ProxyUserAgent        interface{}   `json:"proxyUserAgent"`
	Random                bool          `json:"random"`
	RandomAlgorithm       int           `json:"randomAlgorithm"`
	RandomInGroup         bool          `json:"randomInGroup"`
	ReconnectTimes        int           `json:"reconnectTimes"`
	SameHostForSameTarget bool          `json:"sameHostForSameTarget"`
	ServerSubscribes      []interface{} `json:"serverSubscribes"`
	ShareOverLan          bool          `json:"shareOverLan"`
	SysProxyMode          int           `json:"sysProxyMode"`
	Token                 struct{}      `json:"token"`
}

var ConfigFilePath = "./gui-config.json.bak"
var Group = "free-ss.site"
var WriteFilePath = "./gui-config.json"

func getConfig() GuiConfig {

	b, _ := ioutil.ReadFile(ConfigFilePath)

	var guiConfig GuiConfig
	err := json.Unmarshal(b, &guiConfig)

	if err != nil {
		fmt.Println(err)
	}
	return guiConfig

}

func setConfigs(s ServerList, gc GuiConfig) {

	for _, v := range s {
		config := Configs{Server: v.Server, ServerPort: v.ServerPort, Password: v.Password, Method: v.Method, Group: Group, Enable: true}
		gc.Configs = append(gc.Configs, config)
	}
	gcJson, _ := json.MarshalIndent(gc, "", "	")
	ioutil.WriteFile(WriteFilePath, gcJson, os.ModePerm)
}

package main

import (
	"fmt"
	"os/exec"
)

func main() {

	serverList := GetServerList()
	aserverList := SpeedTest(serverList, 100)
	guiConfig := getConfig()
	setConfigs(aserverList, guiConfig)

	taskkill := exec.Command("cmd", "/C taskkill /f /t /im ShadowsocksR-dotnet4.0.exe")
	if b, err := taskkill.Output(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", b)
	}

	exec.Command("cmd", "/c D:\\greensoft\\ShadowsocksR-win-4.8.0\\ShadowsocksR-dotnet4.0.exe").Start()
}

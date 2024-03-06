-e package main

import (
	"fmt"
	"os/exec"
	"os"
	"io/ioutil"
)

func main() {
	// 输出Hello World
	fmt.Println("Hello, World!")

	// 获取实时新闻
	news, err := getRealTimeNews()
	if err == nil {
		fmt.Println("Real-time News:", news)
	} else {
		fmt.Println("Error fetching news:", err)
	}
}

func getRealTimeNews() (string, error) {
	// 使用curl获取实时新闻
	cmd := exec.Command("curl", "https://newsapi.org/v2/top-headlines?country=us&apiKey=YOUR_API_KEY")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

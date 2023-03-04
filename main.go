package main

import (
	"github.com/sirupsen/logrus"

	"chatgptSDK/chat"
	"chatgptSDK/constant"
)

func main() {
	c := chat.NewChat(constant.GPT35Turbo0301Model, "<<api key>>", 2)
	messages, cost, err := c.ChatWithMessage(c.PlaySystemRole("You are a very helpful assistant"), c.PlayUserRole("What is VTA call graph construct algorithm"))
	if err != nil {
		logrus.Errorf("chat with chatgpt fail: %s", err)
		return
	}
	logrus.Infof("chat cost %d tokens", cost)
	for idx, msg := range messages {
		logrus.Infof("------------------chat response %d-------------------------", idx)
		logrus.Infof("chatgpt response:\n%s", msg.Content)
	}
}

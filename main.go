package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"chatgptSDK/chat"
	"chatgptSDK/constant"
)

var opts struct {
	APIKey    string `short:"k" long:"key" description:"openai api key" required:"true"`
	ChoiceNum int    `short:"n" long:"number" description:"chat response answer number"`
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		fmt.Println("[!] invalid tool argument ")
	}

	c := chat.NewChat(constant.GPT35Turbo0301Model, opts.APIKey, opts.ChoiceNum)
	messages, cost, err := c.ChatWithMessage(c.PlaySystemRole("You are a very helpful assistant"), c.PlayUserRole("What is VTA call graph construct algorithm"))
	if err != nil {
		fmt.Printf("[!] chat with chatgpt fail: %s\n", err)
		return
	}
	fmt.Printf("[+] chat cost %d tokens\n", cost)
	for idx, msg := range messages {
		fmt.Printf("------------------chat response %d-------------------------\n", idx+1)
		fmt.Printf("%s\n", msg.Content)
	}
}

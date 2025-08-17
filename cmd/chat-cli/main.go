package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vaibhavkothari33/go-gemini-cli/internal/conversation"
	"github.com/vaibhavkothari33/go-gemini-cli/internal/gemini"
)

func main() {
	godotenv.Load()
	reader := bufio.NewReader(os.Stdin)

	memory := conversation.NewMemory()

	for {
		fmt.Println("--------------------------Gemini-Go-Bot--------------------- ")
		fmt.Println("Please write your message")
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]

		if input == "exit" {
			fmt.Print("Bye Bye! ")
			break
		}
		memory.Add("user", input)

		prompt := memory.BuildPrompt(input)
		reply, err := gemini.AskGemini((prompt))
		if err != nil {
			panic(err)
		}

		memory.Add("Assistant", reply)

		fmt.Println("BOT:", reply)
	}
}

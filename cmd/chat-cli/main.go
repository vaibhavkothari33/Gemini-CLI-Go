package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vaibhavkothari33/go-gemini-cli/internal/conversation"
	"github.com/vaibhavkothari33/go-gemini-cli/internal/gemini"
)

func main() {
    // Load .env
    godotenv.Load()

    reader := bufio.NewReader(os.Stdin)
    memory := conversation.NewMemory()

    fmt.Println("🤖 Gemini Chatbot CLI")
    fmt.Println("Type 'exit' to quit.")
    fmt.Println("----------------------------")

    for {
        fmt.Print("You: ")
        input, _ := reader.ReadString('\n')
        input = input[:len(input)-1] // remove newline

        if input == "exit" {
            fmt.Println("Bye! 👋")
            break
        }

        // Add user input to memory
        memory.Add("user", input)

        // Build prompt with history
        prompt := memory.BuildPrompt(input)

        // Ask Gemini
        reply, err := gemini.AskGemini(prompt)
        if err != nil {
            log.Println("Error:", err)
            continue
        }

        // Save AI response
        memory.Add("assistant", reply)

        fmt.Println("Bot:", reply)
    }
}

# Go Gemini CLI Chat

A command-line interface chat application built in Go that uses Google's Gemini AI model for conversation.

## Features

- Interactive CLI chat interface
- Conversation memory that maintains chat history
- Integration with Google's Gemini AI model
- Environment variable support for API key configuration

## Prerequisites

- Go 1.24.3 or higher
- Google Gemini API key

## Setup

1. Clone the repository
2. Create a `.env` file in the root directory with your Gemini API key:
```
GOOGLE_API_KEY=your_api_key_here
```

## Installation

```bash
go mod download
go build ./cmd/chat-cli
```

## Usage

Run the application:
```bash
./chat-cli
```

- Type your message and press Enter to chat with the bot
- Type 'exit' to end the conversation

## Project Structure

- `cmd/chat-cli/`: Contains the main application entry point
- `internal/conversation/`: Handles chat history and memory management
- `internal/gemini/`: Implements the Gemini API client

## Dependencies

- github.com/joho/godotenv - For environment variable management
- google.golang.org/genai - Google Gemini AI
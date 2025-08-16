package conversation

type Message struct {
    Role    string
    Content string
}

type Memory struct {
    History []Message
}

func NewMemory() *Memory {
    return &Memory{History: []Message{}}
}

func (m *Memory) Add(role, content string) {
    m.History = append(m.History, Message{Role: role, Content: content})
}

func (m *Memory) BuildPrompt(userMsg string) string {
    prompt := ""
    for _, msg := range m.History {
        prompt += msg.Role + ": " + msg.Content + "\n"
    }
    prompt += "user: " + userMsg
    return prompt
}

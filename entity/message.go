package entity

type Message struct {
	Sender  string   `json:"sender"`
	Name    string   `json:"name"`
	Content []string `json:"content"`
}

func NewMessage(name string, content ...string) *Message {
	return &Message{
		Sender:  "server",
		Name:    name,
		Content: content,
	}
}

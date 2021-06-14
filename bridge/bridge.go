package bridge

import "fmt"

type Message interface {
	BuildMessage(to string, text string) string
}

type MessageSMS struct{}

func (m *MessageSMS) BuildMessage(to string, text string) string {
	return fmt.Sprintf("hello %s,%s via sms,", to, text)
}

type MessageEmail struct{}

func (m *MessageEmail) BuildMessage(to string, text string) string {
	return fmt.Sprintf("hello %s,%s via email,", to, text)
}

type SendMessage interface {
	Send(string, string, Message) string
}

type SendCommonMessage struct {
}

func (s *SendCommonMessage) Send(to, text string, m Message) string {
	return m.BuildMessage(to, text) + "it is common level"
}

type SendUrgencyMessage struct {
}

func (s *SendUrgencyMessage) Send(to, text string, m Message) string {
	return m.BuildMessage(to, text) + "it is urgency level"
}

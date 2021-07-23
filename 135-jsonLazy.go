package main

import (
	"encoding/json"
	"fmt"
)

type MessageType string
type Message struct {
	MessageType MessageType `json:"messageType"`

	MessageId int64 `json:"messageId"`
	//payload暂时不要执行反序列化
	Payload json.RawMessage `json:"payload"`
}
type OnePayload struct {
	One int32
	Two int32
}

func main() {
	m := &Message{}
	data := `{"messageType":"one","messageId":1,"payload":{"one":1,"two":2}}`
	err := json.Unmarshal(([]byte)(data), m)
	fmt.Println(err)
	if err != nil {
		panic("error")
	}
	if m.MessageType == "one" {
		payloadA := &OnePayload{}
		_ = json.Unmarshal(m.Payload, payloadA)
		s, _ := json.Marshal(payloadA)
		fmt.Println((string)(s))
	}
}

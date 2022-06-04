package exampleProtocol

import(
	"testing"
)

func TestMessageParse(t *testing.T) {
	message := "Test"
	data := createMessage(MessageTypeText, message)
	mtype, mlen, msg := readMessage(data)
	
	if mtype != MessageTypeText {
		t.Errorf("Message type is not correct")
	}

	if mlen != uint32(len(message)) {
		t.Errorf("Message length is not correct")
	}

	if msg != message {
		t.Errorf("Message is not correct")
	}

}
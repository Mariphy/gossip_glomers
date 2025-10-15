package main

import (
    "encoding/json"
    "log"
     maelstrom "github.com/jepsen-io/maelstrom/demo/go"
    "github.com/google/uuid"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		id := uuid.New()

		// Update the message type to return back.
		body["type"] = "generate_ok"
		body["id"] = id.String()

		// Echo the original message back with the updated message type.
		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
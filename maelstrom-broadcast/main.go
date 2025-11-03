package main

import (
    "encoding/json"
    "log"
    maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type TopologyPayload struct {
	Type string `json:"type"`
	Topology map[string][]string `json:"topology"`
}

func main() {
	n := maelstrom.NewNode()
	var messages []int
	n.Handle("broadcast", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		response := make(map[string]any)
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.

		messages = append(messages, int(body["message"].(float64)))
		response["type"] = "broadcast_ok"
		// Echo the original message back with the updated message type.
		return n.Reply(msg, response)
	})
	n.Handle("read", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		response := make(map[string]any)
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.
		response["type"] = "read_ok"
		response["messages"] = messages

		// Echo the original message back with the updated message type.
		return n.Reply(msg, response)
	})
	n.Handle("topology", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		response := make(map[string]any)
		var payload TopologyPayload
		if err := json.Unmarshal(msg.Body, &payload); err != nil {
			return err
		}

		//topology = payload.Topology
		
		// Update the message type to return back.
		response["type"] = "topology_ok"

		// Echo the original message back with the updated message type.
		return n.Reply(msg, response)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
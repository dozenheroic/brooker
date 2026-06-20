package broker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"message-broker/manager"
	"message-broker/storage"
)

var state = manager.LoadState()

func StartBroker() {
	http.HandleFunc("/publish", publishHandler)
	http.HandleFunc("/poll", pollHandler)

	fmt.Println("Broker running :8080")
	_ = http.ListenAndServe(":8080", nil)
}

// 📦 publish
func publishHandler(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")
	msg := r.URL.Query().Get("msg")

	if topic == "" || msg == "" {
		http.Error(w, "missing params", 400)
		return
	}

	_ = storage.AppendMessage(topic, msg)
	w.Write([]byte("OK"))
}

// 📥 poll (100% fixed)
func pollHandler(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")
	group := r.URL.Query().Get("group")

	offset := state.GetOffset(group, topic)

	file := "storage/" + topic + ".log"

	data, err := os.ReadFile(file)
	if err != nil {
		json.NewEncoder(w).Encode([]string{})
		return
	}

	raw := strings.TrimSpace(string(data))
	if raw == "" {
		json.NewEncoder(w).Encode([]string{})
		return
	}

	lines := strings.Split(raw, "\n")

	if offset >= len(lines) {
		json.NewEncoder(w).Encode([]string{})
		return
	}

	// ✔ читаем только 1 сообщение
	msg := lines[offset]

	// ✔ сдвигаем offset
	state.SetOffset(group, topic, offset+1)
	_ = state.Save()

	json.NewEncoder(w).Encode([]string{msg})
}

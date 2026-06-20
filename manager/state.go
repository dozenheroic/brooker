package manager

import (
	"encoding/json"
	"os"
)

type State struct {
	Offsets map[string]map[string]int `json:"offsets"`
}

var file = "storage/state.json"

func LoadState() *State {
	s := &State{
		Offsets: map[string]map[string]int{},
	}

	data, err := os.ReadFile(file)
	if err == nil {
		_ = json.Unmarshal(data, s)
	}

	return s
}

func (s *State) Save() error {
	data, _ := json.MarshalIndent(s, "", "  ")
	return os.WriteFile(file, data, 0644)
}

func (s *State) GetOffset(group, topic string) int {
	if s.Offsets[group] == nil {
		return 0
	}
	return s.Offsets[group][topic]
}

func (s *State) SetOffset(group, topic string, val int) {
	if s.Offsets[group] == nil {
		s.Offsets[group] = map[string]int{}
	}
	s.Offsets[group][topic] = val
}

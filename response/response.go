package response

import "encoding/json"

type Success struct {
	Data any
}

func (s *Success) ToJSON() ([]byte, error) {
	return json.Marshal(s.Data)
}

type Failed struct {
	Message string
}

func (f *Failed) ToJSON() ([]byte, error) {
	return json.Marshal(f.Message)
}

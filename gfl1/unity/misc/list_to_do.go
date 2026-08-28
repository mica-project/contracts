package misc

import "encoding/json/v2"

type ListToDo struct {
	Description string `json:"description"`
	Group       int64  `json:"group"`
	Id          int64  `json:"id"`
	Sort        int64  `json:"sort"`
	Type        int64  `json:"type"`
}

type ListToDoList []ListToDo

func (s *ListToDo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

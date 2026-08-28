package bingo

import "encoding/json/v2"

type BingoTaskType struct {
	Content  string `json:"content"`
	Id       int64  `json:"id"`
	TaskType string `json:"task_type"`
	Title    string `json:"title"`
}

type BingoTaskTypeList []BingoTaskType

func (s *BingoTaskType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

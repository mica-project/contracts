package bingo

import "encoding/json/v2"

type BingoTaskInfo struct {
	Content    string `json:"content"`
	IsMutex    string `json:"is_mutex"`
	Size       string `json:"size"`
	TaskId     string `json:"task_id"`
	TaskName   string `json:"task_name"`
	TaskTypeId string `json:"task_type_id"`
	TicketNum  string `json:"ticket_num"`
	Title      string `json:"title"`
	Type       string `json:"type"`
}

type BingoTaskInfoList []BingoTaskInfo

func (s *BingoTaskInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

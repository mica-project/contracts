package quest

import "encoding/json/v2"

type DailyNote struct {
	Code  string `json:"code"`
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type DailyNoteList []DailyNote

func (s *DailyNote) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

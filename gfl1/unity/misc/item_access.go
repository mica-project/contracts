package misc

import "encoding/json/v2"

type ItemAccess struct {
	CarnivalGotoPageId int64  `json:"carnival_goto_page_id"`
	Id                 int64  `json:"id"`
	Name               string `json:"name"`
}

type ItemAccessList []ItemAccess

func (s *ItemAccess) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

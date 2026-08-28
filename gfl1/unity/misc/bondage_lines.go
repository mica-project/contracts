package misc

import "encoding/json/v2"

type BondageLines struct {
	Id        int64    `json:"id"`
	LinesCode string   `json:"lines_code"`
	LinesTxt  []string `json:"lines_txt"`
	MainId    string   `json:"main_id"`
	SubId     string   `json:"sub_id"`
}

type BondageLinesList []BondageLines

func (s *BondageLines) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

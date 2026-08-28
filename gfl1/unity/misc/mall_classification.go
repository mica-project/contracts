package misc

import "encoding/json/v2"

type MallClassification struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Sort        int64  `json:"sort"`
}

type MallClassificationList []MallClassification

func (s *MallClassification) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

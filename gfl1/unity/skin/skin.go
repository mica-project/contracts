package skin

import "encoding/json/v2"

type Skin struct {
	Ai            int64     `json:"ai"`
	ClassId       int64     `json:"class_id"`
	Dialog        string    `json:"dialog"`
	ExploreTag    string    `json:"explore_tag"`
	FitGun        int64     `json:"fit_gun"`
	GiftPosition  []float64 `json:"gift_position"`
	Id            int64     `json:"id"`
	IllustratorCv []string  `json:"illustrator_cv"`
	Name          string    `json:"name"`
	Note          []string  `json:"note"`
	Order         int64     `json:"order"`
	SkinSource    int64     `json:"skin_source"`
	Voice         int64     `json:"voice"`
}

type SkinList []Skin

func (s *Skin) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

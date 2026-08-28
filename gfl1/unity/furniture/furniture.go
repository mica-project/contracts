package furniture

import "encoding/json/v2"

type Furniture struct {
	Classes             int64     `json:"classes"`
	Code                string    `json:"code"`
	DecoRate            int64     `json:"deco_rate"`
	DecomposeGift       int64     `json:"decompose_gift"`
	Description         []string  `json:"description"`
	Id                  int64     `json:"id"`
	InteractPoint       []int64   `json:"interact_point"`
	InteractPointOffset []string  `json:"interact_point_offset"`
	Name                string    `json:"name"`
	Offset              []float64 `json:"offset"`
	Position            []int64   `json:"position"`
	Rotate              string    `json:"rotate"`
	Space               []int64   `json:"space"`
	TextureType         string    `json:"texture_type"`
	Type                int64     `json:"type"`
}

type FurnitureList []Furniture

func (s *Furniture) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

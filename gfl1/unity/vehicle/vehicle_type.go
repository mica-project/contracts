package vehicle

import "encoding/json/v2"

type VehicleType struct {
	Code    string  `json:"code"`
	FixTime int64   `json:"fix_time"`
	Id      int64   `json:"id"`
	MpFix   int64   `json:"mp_fix"`
	Name    string  `json:"name"`
	PartFix float64 `json:"part_fix"`
}

type VehicleTypeList []VehicleType

func (s *VehicleType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

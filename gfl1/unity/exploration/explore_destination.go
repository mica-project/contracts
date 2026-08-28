package exploration

import "encoding/json/v2"

type ExploreDestination struct {
	AreaId int64 `json:"area_id"`
	Id     int64 `json:"id"`
}

type ExploreDestinationList []ExploreDestination

func (s *ExploreDestination) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

package theater

import "encoding/json/v2"

type TheaterIncident struct {
	Description       string  `json:"description"`
	Icon              string  `json:"icon"`
	Id                int64   `json:"id"`
	IsActive          int64   `json:"is_active"`
	Name              string  `json:"name"`
	ScoutMajorityCoef int64   `json:"scout_majority_coef"`
	ScoutManorityCoef int64   `json:"scout_manority_coef"`
	ScoutSelectionId  []int64 `json:"scout_selection_id"`
	TheaterEventId    int64   `json:"theater_event_id"`
	Timing            int64   `json:"timing"`
	Type              int64   `json:"type"`
}

type TheaterIncidentList []TheaterIncident

func (s *TheaterIncident) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

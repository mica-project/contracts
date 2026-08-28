package misc

import "encoding/json/v2"

type FunctionControlInfo struct {
	Description     string `json:"description"`
	FunctionName    string `json:"function_name"`
	Id              string `json:"id"`
	MallId          string `json:"mall_id"`
	MaxCostId       string `json:"max_cost_id"`
	NeedCopperMedal string `json:"need_copper_medal"`
	TurnOn          string `json:"turn_on"`
}

type FunctionControlInfoList []FunctionControlInfo

func (s *FunctionControlInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

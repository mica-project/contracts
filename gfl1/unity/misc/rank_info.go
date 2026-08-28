package misc

import "encoding/json/v2"

type RankInfo struct {
	EveryTableNum       string `json:"every_table_num"`
	Id                  string `json:"id"`
	Name                string `json:"name"`
	RankingOnlistNumber string `json:"ranking_onlist_number"`
	Refresh             string `json:"refresh"`
	SubRank             string `json:"sub_rank"`
	SubRankName         string `json:"sub_rank_name"`
	Title               string `json:"title"`
	Type                string `json:"type"`
	VisableCount        string `json:"visable_count"`
	VisableRank         string `json:"visable_rank"`
}

type RankInfoList []RankInfo

func (s *RankInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

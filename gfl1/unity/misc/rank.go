package misc

import "encoding/json/v2"

type Rank struct {
	Id                  int64  `json:"id"`
	Name                string `json:"name"`
	RankingOnlistNumber int64  `json:"ranking_onlist_number"`
	Refresh             string `json:"refresh"`
	SubRank             int64  `json:"sub_rank"`
	SubRankName         string `json:"sub_rank_name"`
	Title               string `json:"title"`
	Type                int64  `json:"type"`
	VisableCount        int64  `json:"visable_count"`
	VisableRank         int64  `json:"visable_rank"`
}

type RankList []Rank

func (s *Rank) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

package bingo

import "encoding/json/v2"

type BingoInfo struct {
	ChooseNumberFromCredit string    `json:"choose_number_from_credit"`
	CostRaffleTicket       string    `json:"cost_raffle_ticket"`
	Endtime                string    `json:"endtime"`
	EventStage             string    `json:"event_stage"`
	FinalPrizeConfig       []float64 `json:"final_prize_config"`
	GridNum                string    `json:"grid_num"`
	Id                     string    `json:"id"`
	Img1                   string    `json:"img_1"`
	LinePrizeConfig        []int64   `json:"line_prize_config"`
	NormalTaskNum          string    `json:"normal_task_num"`
	PayTaskNum             string    `json:"pay_task_num"`
	RepeatNumberForCredit  string    `json:"repeat_number_for_credit"`
	Starttime              string    `json:"starttime"`
	Text1                  string    `json:"text_1"`
	Text2                  string    `json:"text_2"`
	Text3                  string    `json:"text_3"`
	Text4                  string    `json:"text_4"`
	Text5                  string    `json:"text_5"`
}

type BingoInfoList []BingoInfo

func (s *BingoInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}

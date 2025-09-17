package game

type DrawEvent struct {
	Id               int    `json:"id"`
	ItemId           int    `json:"item_id"`
	Type             int    `json:"type"`
	StartTime        string `json:"start_time"`
	EndTime          string `json:"end_time"`
	DropIds          string `json:"drop_ids"`
	TitleRes         string `json:"title_res"`
	AmountCoordinate string `json:"amount_coordinate"`
	BgRes            string `json:"bg_res"`
	IsShow           int    `json:"is_show"`
	PrizeSkip        string `json:"prize_skip"`
	CanTenDraws      int    `json:"can_ten_draws"`
	UseAnimation     int    `json:"use_animation"`
	GotoToMall       int    `json:"goto_to_mall"`
}

type DrawEventInfo struct {
	Id               string `json:"id"`
	ItemId           string `json:"item_id"`
	StartTime        string `json:"start_time"`
	EndTime          string `json:"end_time"`
	DropIds          string `json:"drop_ids"`
	TitleRes         string `json:"title_res"`
	AmountCoordinate string `json:"amount_coordinate"`
	BgRes            string `json:"bg_res"`
	IsShow           string `json:"is_show"`
	PrizeSkip        string `json:"prize_skip"`
	CanTenDraws      string `json:"can_ten_draws"`
	Type             string `json:"type"`
	UseAnimation     string `json:"use_animation"`
	GotoToMall       string `json:"goto_to_mall"`
}

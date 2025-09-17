package game

type BingoInfo struct {
	Id                     string `json:"id"`
	EventStage             string `json:"event_stage"`
	GridNum                string `json:"grid_num"`
	NormalTaskNum          string `json:"normal_task_num"`
	PayTaskNum             string `json:"pay_task_num"`
	RepeatNumberForCredit  string `json:"repeat_number_for_credit"`
	ChooseNumberFromCredit string `json:"choose_number_from_credit"`
	LinePrizeConfig        string `json:"line_prize_config"`
	FinalPrizeConfig       string `json:"final_prize_config"`
	CostRaffleTicket       string `json:"cost_raffle_ticket"`
	Starttime              string `json:"starttime"`
	Endtime                string `json:"endtime"`
	Img1                   string `json:"img_1"`
	Text1                  string `json:"text_1"`
	Text2                  string `json:"text_2"`
	Text3                  string `json:"text_3"`
	Text4                  string `json:"text_4"`
	Text5                  string `json:"text_5"`
}

type BingoTaskInfo struct {
	TaskId     string `json:"task_id"`
	Type       string `json:"type"`
	TaskName   string `json:"task_name"`
	Size       string `json:"size"`
	TicketNum  string `json:"ticket_num"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	IsMutex    string `json:"is_mutex"`
	TaskTypeId string `json:"task_type_id"`
}

type BingoTaskType struct {
	Id       int    `json:"id"`
	TaskType string `json:"task_type"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

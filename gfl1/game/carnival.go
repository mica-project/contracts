package game

type CarnivalInfo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	PtItemId  string `json:"pt_item_id"`
	PtGift    string `json:"pt_gift"`
	StartTime string `json:"start_time"`
	Banner    string `json:"banner"`
	EndTime   string `json:"end_time"`
	LabelIds  string `json:"label_ids"`
}

type CarnivalGotoPageInfo struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	GotoPage string `json:"goto_page"`
}

type CarnivalLabelInfo struct {
	Id            string `json:"id"`
	StartTime     string `json:"start_time"`
	LabelText     string `json:"label_text"`
	CarnivalTasks string `json:"carnival_tasks"`
}

type CarnivalTask struct {
	Id             int    `json:"id"`
	Type           string `json:"type"`
	CarnivalTypeId int    `json:"carnival_type_id"`
	Count          int    `json:"count"`
	PrizeId        int    `json:"prize_id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
}

type CarnivalTaskInfo struct {
	Id             string `json:"id"`
	Type           string `json:"type"`
	Count          string `json:"count"`
	PrizeId        string `json:"prize_id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	CarnivalTypeId string `json:"carnival_type_id"`
}

type CarnivalTaskType struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

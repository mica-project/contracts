package game

type CareerQuest struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Count   int    `json:"count"`
	PrizeId int    `json:"prize_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	NewType int    `json:"new_type"`
	GradeId int    `json:"grade_id"`
	Sort    int    `json:"sort"`
}

type CareerQuestGrade struct {
	Id          int    `json:"id"`
	Group       int    `json:"group"`
	Grade       int    `json:"grade"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CareerQuestGroup struct {
	Id                int    `json:"id"`
	Type              int    `json:"type"`
	SubType           int    `json:"sub_type"`
	Title             string `json:"title"`
	SubTitle          string `json:"sub_title"`
	FunctionControlId int    `json:"function_control_id"`
}

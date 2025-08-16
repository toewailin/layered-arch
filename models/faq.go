package models

type Faq struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Question    string `json:"question"`
	Answer      string `json:"answer"`
}

func (f *Faq) TableName() string {
	return "faqs"
}


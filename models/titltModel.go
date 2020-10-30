package models

import (
	"github.com/astaxie/beego/orm"
)

type Title struct {
	Id           int    `json:"id",orm:"pk;auto"`
	Title        string `json:"title"`
	Code         string `json:"code"`
	Url          string `json:"url"`
	HasSpidered  int    `json:"has_spidered"`
	Md5CodeTitle string `orm:"column(md_5_code_title)",json:"md_5_code_title"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(Title))
}
func (u *Title) TableName() string {
	return "title"
}

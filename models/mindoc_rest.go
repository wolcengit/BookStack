package models

import "github.com/astaxie/beego/orm"

// MinDocRest ...
type MinDocRest struct {
	Token    string `json:"token"`
	Folder   string `json:"folder"`
	Title    string `json:"title"`
	Identify string `json:"identify"`
	TextMD   string `json:"textmd"`
	TextHTML string `json:"texthtml"`
}

// NewMinDocRest ...
func NewMinDocRest() *MinDocRest {
	return &MinDocRest{}
}

func (m *MinDocRest) FindByPrivateToken(token string) (book *Book, err error) {
	o := orm.NewOrm()

	err = o.QueryTable(NewBook().TableNameWithPrefix()).Filter("private_token", token).One(book)

	return book, err
}

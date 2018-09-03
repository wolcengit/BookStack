package controllers

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/TruthHun/BookStack/models"
	"github.com/astaxie/beego/orm"
	"github.com/TruthHun/BookStack/utils"
)

// MinDocRestController struct.
type MinDocRestController struct {
	BaseController
}

// PostContent API 设置文档内容.
func (c *MinDocRestController) PostContent() {
	c.Prepare()

	var req models.MinDocRest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	folderkey := req.Folder
	doctitle := req.Title
	dockey := req.Identify
	tokenkey := req.Token
	textmd := req.TextMD

	if doctitle == "" {
		c.JsonResult(6004, "文档名称不能为空")
	}
	if dockey == "" {
		c.JsonResult(6004, "文档标识不能为空")
	}
	if ok, err := regexp.MatchString(`^[a-z]+[a-zA-Z0-9_\-]*$`, dockey); !ok || err != nil {

		c.JsonResult(6003, "文档标识只能包含小写字母、数字，以及“-”和“_”符号,并且只能小写字母开头")
	}
	var book *models.Book
	book, err := models.NewBook().FindByIdentify(tokenkey)
	if err != nil {
		c.JsonResult(6002, "文档不存在["+tokenkey+"]")
	}
	beego.Info("req tokenkey =>" + tokenkey + "  -->" + fmt.Sprintf("%d", book.BookId))

	folder, err := models.NewDocument().FindByBookIdAndDocIdentify(book.BookId,folderkey)
	if err != nil {
		beego.Error("folder => ", err)
		c.JsonResult(6002, "项目类目不存在")
	}
	beego.Info("req folderkey =>" + folderkey + "  -->" + fmt.Sprintf("%d", book.BookId))

	doc, err := models.NewDocument().FindByBookIdAndDocIdentify(book.BookId,dockey)
	if err != nil {
		if err != orm.ErrNoRows {
			c.JsonResult(6006, "文档获取失败")
		}
		doc = models.NewDocument()
		doc.Identify = dockey
		doc.BookId = book.BookId
		doc.MemberId = book.MemberId
		doc.OrderSort = 0
	}
	doc.ParentId = folder.DocumentId
	doc.DocumentName = doctitle
	ModelStore := new(models.DocumentStore)
	doc_id, err := doc.InsertOrUpdate()
	if err == nil {
		err := ModelStore.InsertOrUpdate(models.DocumentStore{
			DocumentId: int(doc_id),
			Markdown:   textmd,
			Content: "",
		}, "markdown")
		if  err != nil {
			beego.Error(err)
		}else {
			utils.RenderDocumentById(doc.DocumentId)
		}
	} else {
		beego.Error(err.Error())
	}

	c.JsonResult(0, "ok", doc)
}

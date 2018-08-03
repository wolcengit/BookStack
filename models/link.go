package models

import (
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"strings"
)
type LinkDocument struct {
	LinkId           int    `json:"link_id"`
}

type LinkDocumentTree struct {
	DocumentId   int      `json:"id"`
	ParentId     int      `json:"pId"`
	DocumentName string   `json:"name"`
	Checked 	 bool 	  `json:"checked"`
	Open   		 bool     `json:"open"`
}


func NewLinkDocument() *LinkDocument {
	return &LinkDocument{}
}


func (m *LinkDocument) FindForLinksByBookId(book_id, pageIndex, pageSize int) ([]*Book, int, error) {
	o := orm.NewOrm()

	var books []*Book

	sql1 := "SELECT * FROM md_books WHERE link_id = ? ORDER BY book_id DESC  LIMIT ?,?"

	sql2 := "SELECT count(*) AS total_count FROM md_books WHERE link_id = ?"

	var total_count int

	err := o.Raw(sql2, book_id).QueryRow(&total_count)

	if err != nil {
		return books, 0, err
	}

	offset := (pageIndex - 1) * pageSize

	_, err = o.Raw(sql1, book_id, offset, pageSize).QueryRows(&books)

	if err != nil {
		return books, 0, err
	}

	return books, total_count, nil
}

func (m *LinkDocument) SetLinkBookDocuments(book_id int,link_docs string) (err error) {

	book, err := NewBook().Find(book_id)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	doc_cnt := strings.Count(link_docs,",")
	o.Raw("UPDATE md_books SET link_docs = ?,doc_count=? WHERE book_id = ?", link_docs, doc_cnt, book_id).Exec()

	sql := `DELETE FROM md_documents WHERE book_id = ? AND FIND_IN_SET(link_id,?) = 0 `
	o.Raw(sql,book.BookId,link_docs).Exec()

	sql = `INSERT INTO md_documents (document_name,identify,book_id,parent_id,order_sort,create_time,member_id,modify_time,modify_at,version,vcnt,link_id ) SELECT document_name,CONCAT(identify,document_id),?,parent_id,order_sort,create_time,member_id,modify_time,modify_at,version,0,document_id  FROM md_documents  WHERE book_id = ? AND FIND_IN_SET(document_id,?) > 0  AND document_id NOT IN(SELECT link_id FROM md_documents WHERE book_id = ? )`
	o.Raw(sql,book.BookId,book.LinkId,link_docs,book.BookId).Exec()


	sql = `UPDATE md_documents as a, (select b.parent_id as bpid,b.document_id as bdid,c.document_id as cpid from md_documents b ,md_documents c where b.parent_id = c.link_id and c.book_id = ?) as b SET a.parent_id = b.cpid  WHERE a.book_id = ? AND a.parent_id > 0;`
	o.Raw(sql,book.BookId,book.BookId).Exec()

	return nil
}


func (m *LinkDocument) UpdateLinkBookDocuments(book_id int) (err error) {

	o := orm.NewOrm()
	sql := `UPDATE md_documents a,md_documents b SET a.document_name = b.document_name,a.identify = b.identify,a.release = b.release,a.order_sort = b.order_sort,a.modify_time = b.modify_time,a.version = b.version WHERE a.book_id = ? AND a.link_id = b.document_id AND a.modify_at <> b.modify_at `
	o.Raw(sql,book_id).Exec()

	return nil
}


// GetLinkBookDocuments ...[{ id:1, pId:0, name:"数据错误"}];
func (m *LinkDocument) GetLinkBookDocuments(book_id int) (doclinks string, doclist string, err error) {

	book, err := NewBook().Find(book_id)
	if err != nil {
		return "", "", err
	}
	doclinks = book.LinkDocs

	var docs []*Document
	o := orm.NewOrm()
	sql2 := `SELECT document_id,parent_id,document_name,FIND_IN_SET(document_id,?) AS modify_at FROM md_documents WHERE book_id = ? ORDER BY order_sort  ,document_id  `
	count, _ := o.Raw(sql2, book.LinkDocs, book.LinkId).QueryRows(&docs)
	doclist = ""
	if count > 0 {
		trees := make([]*LinkDocumentTree, count)
		for index, item := range docs {
			tree := &LinkDocumentTree{}
			tree.DocumentId = item.DocumentId
			tree.ParentId = item.ParentId
			tree.DocumentName = item.DocumentName
			if item.ModifyAt > 0 {
				tree.Checked = true
			}else{
				tree.Checked = false
			}
			if item.ParentId == 0{
				tree.Open = true
			}else{
				tree.Open = false
			}
			trees[index] = tree
		}
		data,err:=json.Marshal(trees)
		if err != nil {
			return doclinks, "", err
		}
		doclist = string(data)
	}
	return doclinks, doclist, nil
}

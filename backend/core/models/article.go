package models

import "golang.org/x/tools/godoc"

type Article struct {
	Id       uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Title    string `gorm:"column:title;type:varchar(200);NOT NULL" json:"title"`
	Status   int    `gorm:"column:status;type:tinyint(1);default:0;NOT NULL" json:"status"`
	ImageUrl string `gorm:"column:image_url;type:varchar(2000);NOT NULL" json:"imageUrl"`
	CateId   int    `gorm:"column:cate_id;type:int(11);default:0;NOT NULL" json:"cateId"`
	IsTop    int    `gorm:"column:is_top;type:tinyint(1);default:0;NOT NULL" json:"isTop"`
	Views    int    `gorm:"column:views;type:int(11);default:0;NOT NULL" json:"views"`
	Desc     string `gorm:"column:desc;type:varchar(2000);NOT NULL" json:"desc"`
	Content  string `gorm:"column:content;type:longtext" json:"content"`
	Model
}

func (m *Article) TableNmae() string { return "article" }

type StatRes struct {
	ArticleCount  int64 `json:articleCount`
	CategoryCount int64 `json:categoryCount`
	PageCount     int64 `json:pageCount`
	TagCount      int64 `json:tagCount`
}

type ArticleListReq struct {
	godoc.PageInfo
	Keyword string `form:"keyword" default:"a"`
	CateId  int    `form:"cateId" default:"0"`
}
type ArticleListRes struct {
	godoc.PageInfo
	Data []Article `json:data`
}

type ArticleInfoReq struct {
	//
}
type ArticleAddReq struct {
	Title string `json:"title" validate:"required"`
}

type ArticleUpdateReq struct {
	Id    uint   `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type ArticleDelReq struct {
	Id uint `json:"id" validate:"required"`
}

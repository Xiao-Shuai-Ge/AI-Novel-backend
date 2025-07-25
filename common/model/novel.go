package model

type Novel struct {
	ID int64 `json:"id" gorm:"column:id;primaryKey;type:bigint;type:bigint"`
	TimeModel
	AuthorID int64  `json:"author_id" gorm:"column:author_id;type:bigint"`
	Title    string `json:"title" gorm:"column:title;type:varchar(255)"`
	Avatar   string `json:"avatar" gorm:"column:avatar;type:varchar(255)"`
	Summary  string `json:"summary" gorm:"column:summary;type:text"`
	Status   int    `json:"status" gorm:"column:status;type:int"`
	IsPublic bool   `json:"is_public" gorm:"column:is_public;type:tinyint(1)"`

	Chapters   []Chapter   `gorm:"foreignKey:NovelID;constraint:OnDelete:CASCADE;"`
	Characters []Character `gorm:"foreignKey:NovelID;constraint:OnDelete:CASCADE;"`
}

func (t Novel) TableName() string {
	return "novels"
}

type Chapter struct {
	ID int64 `json:"id" gorm:"column:id;primaryKey;type:bigint;type:bigint"`
	TimeModel
	NovelID int64  `json:"novel_id" gorm:"column:novel_id;type:bigint;index"`
	Title   string `json:"title" gorm:"column:title;type:varchar(255)"`
	Content string `json:"content" gorm:"column:content;type:text"`
	Summary string `json:"summary" gorm:"column:summary;type:text"`

	Novel Novel `gorm:"foreignKey:NovelID;references:ID;constraint:OnDelete:CASCADE;"`
}

func (t Chapter) TableName() string {
	return "chapters"
}

type Character struct {
	ID int64 `json:"id" gorm:"column:id;primaryKey;type:bigint;type:bigint"`
	TimeModel
	NovelID int64  `json:"novel_id" gorm:"column:novel_id;type:bigint;index"`
	Name    string `json:"name" gorm:"column:name;type:varchar(255)"`
	Avatar  string `json:"avatar" gorm:"column:avatar;type:varchar(255)"`
	Summary string `json:"summary" gorm:"column:summary;type:text"`

	Novel Novel `gorm:"foreignKey:NovelID;references:ID;constraint:OnDelete:CASCADE;"`
}

func (t Character) TableName() string {
	return "characters"
}

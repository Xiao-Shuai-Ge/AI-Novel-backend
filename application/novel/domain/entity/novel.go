package entity

import "Ai-Novel/common/model"

type Novel struct {
	ID       int64
	Title    string
	Avatar   string
	AuthorID int64
	Summary  string
	Status   int
	IsPublic bool
}

func NewNovel(id int64, title string, avatar string, authorID int64, summary string, status int, isPublic bool) Novel {
	return Novel{
		ID:       id,
		Title:    title,
		Avatar:   avatar,
		AuthorID: authorID,
		Status:   status,
		IsPublic: isPublic,
	}
}

func (u *Novel) Transform() *model.Novel {
	return &model.Novel{
		ID:       u.ID,
		Title:    u.Title,
		Avatar:   u.Avatar,
		AuthorID: u.AuthorID,
		Status:   u.Status,
		IsPublic: u.IsPublic,
	}
}

func FormNovel(novel *model.Novel) Novel {
	return Novel{
		ID:       novel.ID,
		Title:    novel.Title,
		Avatar:   novel.Avatar,
		AuthorID: novel.AuthorID,
		Status:   novel.Status,
		IsPublic: novel.IsPublic,
	}
}

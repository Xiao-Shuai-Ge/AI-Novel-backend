package entity

import "Ai-Novel/common/model"

type Chapter struct {
	ID      int64
	NovelID int64
	Title   string
	Content string
	Summary string
}

func NewChapter(id int64, novelID int64, title string, content string, summary string) Chapter {
	return Chapter{
		ID:      id,
		NovelID: novelID,
		Title:   title,
		Content: content,
		Summary: summary,
	}
}

func (u *Chapter) Transform() *model.Chapter {
	return &model.Chapter{
		ID:      u.ID,
		NovelID: u.NovelID,
		Title:   u.Title,
		Content: u.Content,
		Summary: u.Summary,
	}
}

func FormChapter(chapter *model.Chapter) Chapter {
	return Chapter{
		ID:      chapter.ID,
		NovelID: chapter.NovelID,
		Title:   chapter.Title,
		Content: chapter.Content,
		Summary: chapter.Summary,
	}
}

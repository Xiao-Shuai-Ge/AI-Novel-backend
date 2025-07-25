package entity

import "Ai-Novel/common/model"

type Character struct {
	ID      int64
	NovelID int64
	Avatar  string
	Name    string
	Summary string
}

func NewCharacter(id int64, novelID int64, avatar string, name string, summary string) Character {
	return Character{
		ID:      id,
		NovelID: novelID,
		Avatar:  avatar,
		Name:    name,
		Summary: summary,
	}
}

func (u *Character) Transform() *model.Character {
	return &model.Character{
		ID:      u.ID,
		NovelID: u.NovelID,
		Avatar:  u.Avatar,
		Name:    u.Name,
		Summary: u.Summary,
	}
}

func FormCharacter(character *model.Character) Character {
	return Character{
		ID:      character.ID,
		NovelID: character.NovelID,
		Avatar:  character.Avatar,
		Name:    character.Name,
		Summary: character.Summary,
	}
}

package repo

import (
	"Ai-Novel/application/novel/domain/entity"
	"Ai-Novel/common/model"
)

func (r *NovelRepo) CreateNovel(novel entity.Novel) (err error) {
	// 转换为数据库模型
	DbNovel := novel.Transform()
	// 存入数据库
	return r.db.Create(DbNovel).Error
}

func (r *NovelRepo) CreateChapter(chapter entity.Chapter) (err error) {
	// 转换为数据库模型
	DbChapter := chapter.Transform()
	// 存入数据库
	return r.db.Create(DbChapter).Error
}

func (r *NovelRepo) CreateCharacter(character entity.Character) (err error) {
	// 转换为数据库模型
	DbCharacter := character.Transform()
	// 存入数据库
	return r.db.Create(DbCharacter).Error
}

func (r *NovelRepo) GetNovel(id int64) (novel entity.Novel, err error) {
	var DbNovel *model.Novel
	err = r.db.Where("id=?", id).First(&DbNovel).Error
	if err != nil {
		return
	}
	novel = entity.FormNovel(DbNovel)
	return
}

func (r *NovelRepo) GetChapter(id int64) (chapter entity.Chapter, err error) {
	var DbChapter *model.Chapter
	err = r.db.Where("id=?", id).First(&DbChapter).Error
	if err != nil {
		return
	}
	chapter = entity.FormChapter(DbChapter)
	return
}

func (r *NovelRepo) GetCharacter(id int64) (character entity.Character, err error) {
	var DbCharacter *model.Character
	err = r.db.Where("id=?", id).First(&DbCharacter).Error
	if err != nil {
		return
	}
	character = entity.FormCharacter(DbCharacter)
	return
}

func (r *NovelRepo) UpdateNovel(novel entity.Novel) (err error) {
	// 仅修改部分字段
	err = r.db.Model(&model.Novel{}).Where("id=?", novel.ID).Updates(map[string]interface{}{
		"title":     novel.Title,
		"avatar":    novel.Avatar,
		"summary":   novel.Summary,
		"status":    novel.Status,
		"is_public": novel.IsPublic,
	}).Error
	return
}

func (r *NovelRepo) UpdateChapter(chapter entity.Chapter) (err error) {
	// 仅修改部分字段
	err = r.db.Model(&model.Chapter{}).Where("id=?", chapter.ID).Updates(map[string]interface{}{
		"title":   chapter.Title,
		"content": chapter.Content,
		"summary": chapter.Summary,
	}).Error
	return
}

func (r *NovelRepo) UpdateCharacter(character entity.Character) (err error) {
	// 仅修改部分字段
	err = r.db.Model(&model.Character{}).Where("id=?", character.ID).Updates(map[string]interface{}{
		"name":    character.Name,
		"avatar":  character.Avatar,
		"summary": character.Summary,
	}).Error
	return
}

func (r *NovelRepo) DeleteNovel(id int64) (err error) {
	err = r.db.Where("id=?", id).Delete(&model.Novel{}).Error
	return
}

func (r *NovelRepo) DeleteChapter(id int64) (err error) {
	err = r.db.Where("id=?", id).Delete(&model.Chapter{}).Error
	return
}

func (r *NovelRepo) DeleteCharacter(id int64) (err error) {
	err = r.db.Where("id=?", id).Delete(&model.Character{}).Error
	return
}

func (r *NovelRepo) GetChapterList(novel_id int64) (chapters []entity.Chapter, err error) {
	var DbChapters []*model.Chapter
	err = r.db.Where("novel_id=?", novel_id).Order("id").Find(&DbChapters).Error
	if err != nil {
		return
	}
	for _, chapter := range DbChapters {
		chapters = append(chapters, entity.FormChapter(chapter))
	}
	return
}

func (r *NovelRepo) GetCharacterList(novel_id int64) (characters []entity.Character, err error) {
	var DbCharacters []*model.Character
	err = r.db.Where("novel_id=?", novel_id).Order("id").Find(&DbCharacters).Error
	if err != nil {
		return
	}
	for _, character := range DbCharacters {
		characters = append(characters, entity.FormCharacter(character))
	}
	return
}

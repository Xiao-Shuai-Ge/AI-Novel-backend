package app

import (
	"Ai-Novel/application/novel/domain/entity"
	"Ai-Novel/application/novel/domain/services"
	"Ai-Novel/application/novel/infrastructure/repo"
	"Ai-Novel/common/utils/snowflake"
	"Ai-Novel/common/zlog"
	"context"
	"errors"
)

type NovelApp struct {
	Repo          *repo.NovelRepo
	SnowflakeNode *snowflake.Node
}

func NewNovelApp(repo *repo.NovelRepo, snowflakeNode *snowflake.Node) NovelApp {
	return NovelApp{
		Repo:          repo,
		SnowflakeNode: snowflakeNode,
	}
}

// CreateNovel
//
//	@Description: 创建小说
//	@receiver a
//	@param ctx
//	@param title
//	@param avatar
//	@param authorID
//	@param summary
//	@param status
//	@param isPublic
//	@return id
//	@return err
func (a *NovelApp) CreateNovel(ctx context.Context, title, avatar string, authorID int64, summary string, status int, isPublic bool) (id int64, err error) {
	// 1.生成ID
	id = a.SnowflakeNode.Generate().Int64()
	// 2. 创建实体
	novel := entity.NewNovel(id, title, avatar, authorID, summary, status, isPublic)
	// 3. 调用服务保存到数据库
	err = services.NewNovelService(ctx, a.Repo, a.SnowflakeNode).CreateNovel(novel)
	if err != nil {
		zlog.ErrorfCtx(ctx, "创建小说失败: %v", err)
		err = errors.New("创建小说失败")
		return
	}
	return
}

// CreateChapter
//
//	@Description: 创建章节
//	@receiver a
//	@param ctx
//	@param novelID
//	@param operatorID
//	@param title
//	@param content
//	@param summary
//	@return id
//	@return err
func (a *NovelApp) CreateChapter(ctx context.Context, novelID int64, operatorID int64, title, content string, summary string) (id int64, err error) {
	// 1. 调用小说服务获取小说
	novel, err := a.GetNovel(ctx, novelID, operatorID)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	// 2. 验证作者是否有权限创建章节
	if novel.AuthorID != operatorID {
		err = errors.New("无权创建章节")
		return
	}
	// 3. 生成ID
	id = a.SnowflakeNode.Generate().Int64()
	// 4. 创建实体
	chapter := entity.NewChapter(id, novelID, title, content, summary)
	// 5. 调用服务保存到数据库
	err = services.NewNovelService(ctx, a.Repo, a.SnowflakeNode).CreateChapter(chapter)
	if err != nil {
		zlog.ErrorfCtx(ctx, "创建章节失败: %v", err)
		err = errors.New("创建章节失败")
		return
	}
	return
}

// CreateCharacter
//
//	@Description: 创建角色
//	@receiver a
//	@param ctx
//	@param novelID
//	@param operatorID
//	@param name
//	@param avatar
//	@param summary
//	@return id
//	@return err
func (a *NovelApp) CreateCharacter(ctx context.Context, novelID int64, operatorID int64, name, avatar, summary string) (id int64, err error) {
	novel, err := a.GetNovel(ctx, novelID, operatorID)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	if novel.AuthorID != operatorID {
		err = errors.New("无权创建角色")
		return
	}
	id = a.SnowflakeNode.Generate().Int64()
	character := entity.NewCharacter(id, novelID, avatar, name, summary)
	err = services.NewNovelService(ctx, a.Repo, a.SnowflakeNode).CreateCharacter(character)
	if err != nil {
		zlog.ErrorfCtx(ctx, "创建角色失败: %v", err)
		err = errors.New("创建角色失败")
		return
	}
	return
}

// GetNovel
//
//	@Description: 获取小说
//	@receiver a
//	@param ctx
//	@param id
//	@param operatorID
//	@return novel
//	@return err
func (a *NovelApp) GetNovel(ctx context.Context, id int64, operatorID int64) (novel entity.Novel, err error) {
	// 1. 调用服务获取小说
	novel, err = services.NewNovelService(ctx, a.Repo, a.SnowflakeNode).GetNovel(id)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	// 2. 如果为未公开，则需要判断是否有权限查看
	if !novel.IsPublic {
		if novel.AuthorID != operatorID {
			err = errors.New("无权查看此小说")
			return
		}
	}
	return
}

// GetChapter
//
//	@Description: 获取章节内容
//	@receiver a
//	@param ctx
//	@param chapterID
//	@param operatorID
//	@return chapters
//	@return err
func (a *NovelApp) GetChapter(ctx context.Context, chapterID int64, operatorID int64) (chapter entity.Chapter, err error) {
	s := services.NewNovelService(ctx, a.Repo, a.SnowflakeNode)
	// 1. 篇章
	chapter, err = s.GetChapter(chapterID)
	if err != nil {
		err = errors.New("获取章节失败")
		return
	}
	// 2. 获取小说
	novel, err := s.GetNovel(chapter.NovelID)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	// 3. 判断是否有权限查看
	if !novel.IsPublic {
		if novel.AuthorID != operatorID {
			err = errors.New("无权查看此小说")
			return
		}
	}
	// 4. 返回章节
	return
}

// GetCharacter
//
//	@Description: 获取角色信息
//	@receiver a
//	@param ctx
//	@param characterID
//	@param operatorID
//	@return character
//	@return err
func (a *NovelApp) GetCharacter(ctx context.Context, characterID int64, operatorID int64) (character entity.Character, err error) {
	s := services.NewNovelService(ctx, a.Repo, a.SnowflakeNode)
	// 1. 获取角色
	character, err = s.GetCharacter(characterID)
	if err != nil {
		err = errors.New("获取角色失败")
		return
	}
	// 2. 获取小说
	novel, err := s.GetNovel(character.NovelID)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	// 3. 判断是否有权限查看
	if !novel.IsPublic {
		if novel.AuthorID != operatorID {
			err = errors.New("无权查看此小说")
			return
		}
	}
	// 4. 返回角色
	return
}

// UpdateNovel
//
//	@Description: 更新小说
//	@receiver a
//	@param ctx
//	@param id
//	@param title
//	@param avatar
//	@param summary
//	@param authorID
//	@param status
//	@param isPublic
//	@return err
func (a *NovelApp) UpdateNovel(ctx context.Context, id int64, title, avatar, summary string, operatorID int64, status int, isPublic bool) (err error) {
	s := services.NewNovelService(ctx, a.Repo, a.SnowflakeNode)
	// 1. 获取小说
	novel, err := s.GetNovel(id)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	// 2. 如果做出修改的人不是作者，则不允许修改
	if novel.AuthorID != operatorID {
		err = errors.New("无权修改此小说")
		return
	}
	// 3. 修改小说信息
	novel.Title = title
	novel.Avatar = avatar
	novel.Summary = summary
	novel.Status = status
	novel.IsPublic = isPublic
	// 4. 修改小说
	err = s.UpdateNovel(novel)
	if err != nil {
		err = errors.New("修改小说失败")
		return
	}
	// 5. 返回成功
	return
}

// UpdateChapter
//
//	@Description: 更新篇章内容
//	@receiver a
//	@param ctx
//	@param id
//	@param operatorID
//	@param title
//	@param content
//	@param summary
//	@return err
func (a *NovelApp) UpdateChapter(ctx context.Context, id int64, operatorID int64, title, content, summary string) (err error) {
	s := services.NewNovelService(ctx, a.Repo, a.SnowflakeNode)
	// 1. 获取章节
	chapter, err := s.GetChapter(id)
	if err != nil {
		err = errors.New("获取章节失败")
		return
	}
	// 2. 获取小说
	novel, err := s.GetNovel(chapter.NovelID)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	// 3. 判断是否有权限修改
	if novel.AuthorID != operatorID {
		err = errors.New("无权修改此小说")
		return
	}
	// 4. 修改章节信息
	chapter.Title = title
	chapter.Content = content
	chapter.Summary = summary
	// 5. 修改章节
	err = s.UpdateChapter(chapter)
	if err != nil {
		err = errors.New("修改章节失败")
		return
	}
	// 6. 返回成功
	return
}

func (a *NovelApp) UpdateCharacter(ctx context.Context, id int64, operatorID int64, name, avatar, summary string) (err error) {
	s := services.NewNovelService(ctx, a.Repo, a.SnowflakeNode)
	// 1. 获取角色
	character, err := s.GetCharacter(id)
	if err != nil {
		err = errors.New("获取角色失败")
		return
	}
	// 2. 获取小说
	novel, err := s.GetNovel(character.NovelID)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	// 3. 判断是否有权限修改
	if novel.AuthorID != operatorID {
		err = errors.New("无权修改此小说")
		return
	}
	// 4. 修改角色信息
	character.Name = name
	character.Avatar = avatar
	character.Summary = summary
	// 5. 修改角色
	err = s.UpdateCharacter(character)
	if err != nil {
		err = errors.New("修改角色失败")
		return
	}
	// 6. 返回成功
	return
}

// DeleteNovel
//
//	@Description: 删除小说
//	@receiver a
//	@param ctx
//	@param id
//	@param operatorID
//	@return err
func (a *NovelApp) DeleteNovel(ctx context.Context, id int64, operatorID int64) (err error) {
	s := services.NewNovelService(ctx, a.Repo, a.SnowflakeNode)
	// 1. 获取小说
	novel, err := s.GetNovel(id)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	// 2. 如果做出删除的人不是作者，则不允许删除
	if novel.AuthorID != operatorID {
		err = errors.New("无权删除此小说")
		return
	}
	// 3. 删除小说
	err = s.DeleteNovel(id)
	if err != nil {
		err = errors.New("删除小说失败")
		return
	}
	// 4. 返回成功
	return
}

// DeleteChapter
//
//	@Description: 删除章节
//	@receiver a
//	@param ctx
//	@param id
//	@param operatorID
//	@return err
func (a *NovelApp) DeleteChapter(ctx context.Context, id int64, operatorID int64) (err error) {
	s := services.NewNovelService(ctx, a.Repo, a.SnowflakeNode)
	// 1. 获取章节
	chapter, err := s.GetChapter(id)
	if err != nil {
		err = errors.New("获取章节失败")
		return
	}
	// 2. 获取小说
	novel, err := s.GetNovel(chapter.NovelID)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	// 3. 判断是否有权限删除
	if novel.AuthorID != operatorID {
		err = errors.New("无权删除此小说")
		return
	}
	// 4. 删除章节
	err = s.DeleteChapter(id)
	if err != nil {
		err = errors.New("删除章节失败")
		return
	}
	// 5. 返回成功
	return
}

// DeleteCharacter
//
//	@Description: 删除角色
//	@receiver a
//	@param ctx
//	@param id
//	@param operatorID
//	@return err
func (a *NovelApp) DeleteCharacter(ctx context.Context, id int64, operatorID int64) (err error) {
	s := services.NewNovelService(ctx, a.Repo, a.SnowflakeNode)
	// 1. 获取角色
	character, err := s.GetCharacter(id)
	if err != nil {
		err = errors.New("获取角色失败")
		return
	}
	// 2. 获取小说
	novel, err := s.GetNovel(character.NovelID)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	// 3. 判断是否有权限删除
	if novel.AuthorID != operatorID {
		err = errors.New("无权删除此小说")
		return
	}
	// 4. 删除角色
	err = s.DeleteCharacter(id)
	if err != nil {
		err = errors.New("删除角色失败")
		return
	}
	// 5. 返回成功
	return
}

// GetChapterList
//
//	@Description: 获取章节列表
//	@receiver a
//	@param ctx
//	@param novelID
//	@param operatorID
//	@return chapters
//	@return err
func (a *NovelApp) GetChapterList(ctx context.Context, novelID int64, operatorID int64) (chapters []entity.Chapter, err error) {
	s := services.NewNovelService(ctx, a.Repo, a.SnowflakeNode)
	// 1. 获取小说
	novel, err := s.GetNovel(novelID)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	// 2. 判断是否有权限查看
	if !novel.IsPublic {
		if novel.AuthorID != operatorID {
			err = errors.New("无权查看此小说")
			return
		}
	}
	// 3. 获取章节列表
	chapters, err = s.GetChapterList(novelID)
	if err != nil {
		err = errors.New("获取章节列表失败")
		return
	}
	// 4. 返回章节列表
	return
}

// GetCharacterList
//
//	@Description: 获取角色列表
//	@receiver a
//	@param ctx
//	@param novelID
//	@param operatorID
//	@return characters
//	@return err
func (a *NovelApp) GetCharacterList(ctx context.Context, novelID int64, operatorID int64) (characters []entity.Character, err error) {
	s := services.NewNovelService(ctx, a.Repo, a.SnowflakeNode)
	// 1. 获取小说
	novel, err := s.GetNovel(novelID)
	if err != nil {
		err = errors.New("获取小说失败")
		return
	}
	// 2. 判断是否有权限查看
	if !novel.IsPublic {
		if novel.AuthorID != operatorID {
			err = errors.New("无权查看此小说")
			return
		}
	}
	// 3. 获取角色列表
	characters, err = s.GetCharacterList(novelID)
	if err != nil {
		err = errors.New("获取角色列表失败")
		return
	}
	// 4. 返回角色列表
	return
}

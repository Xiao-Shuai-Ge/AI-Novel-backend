package services

import (
	"Ai-Novel/application/novel/domain/entity"
	"Ai-Novel/application/novel/infrastructure/repo"
	"Ai-Novel/common/utils/snowflake"
	"Ai-Novel/common/zlog"
	"context"
)

type NovelService struct {
	ctx           context.Context
	Repo          *repo.NovelRepo
	SnowFlakeNode *snowflake.Node
}

func NewNovelService(ctx context.Context, repo *repo.NovelRepo, SnowFlakeNode *snowflake.Node) NovelService {
	return NovelService{
		ctx:           ctx,
		Repo:          repo,
		SnowFlakeNode: SnowFlakeNode,
	}
}

func (s NovelService) CreateNovel(novel entity.Novel) (err error) {
	// 1. 使用数据库保存
	err = s.Repo.CreateNovel(novel)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return err
	}
	// 2. 返回
	return nil
}

func (s NovelService) CreateChapter(chapter entity.Chapter) (err error) {
	// 1. 使用数据库保存
	err = s.Repo.CreateChapter(chapter)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return err
	}
	// 2. 返回
	return nil
}

func (s NovelService) CreateCharacter(character entity.Character) (err error) {
	err = s.Repo.CreateCharacter(character)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return err
	}
	// 2. 返回
	return nil
}

func (s NovelService) GetNovel(id int64) (novel entity.Novel, err error) {
	// 1. 从数据库获取
	novel, err = s.Repo.GetNovel(id)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return novel, err
	}
	// 2. 返回
	return
}

func (s NovelService) GetChapter(id int64) (chapter entity.Chapter, err error) {
	// 1. 从数据库获取
	chapter, err = s.Repo.GetChapter(id)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return chapter, err
	}
	// 2. 返回
	return
}

func (s NovelService) GetCharacter(id int64) (character entity.Character, err error) {
	// 1. 从数据库获取
	character, err = s.Repo.GetCharacter(id)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return character, err
	}
	// 2. 返回
	return
}

func (s NovelService) UpdateNovel(novel entity.Novel) (err error) {
	// 1. 使用数据库更新
	err = s.Repo.UpdateNovel(novel)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return err
	}
	// 2. 返回
	return
}

func (s NovelService) UpdateChapter(chapter entity.Chapter) (err error) {
	err = s.Repo.UpdateChapter(chapter)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return err
	}
	// 2. 返回
	return nil
}

func (s NovelService) UpdateCharacter(character entity.Character) (err error) {
	err = s.Repo.UpdateCharacter(character)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return err
	}
	// 2. 返回
	return nil
}

func (s NovelService) DeleteNovel(id int64) (err error) {
	// 1. 使用数据库删除
	err = s.Repo.DeleteNovel(id)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return err
	}
	// 2. 返回
	return nil
}

func (s NovelService) DeleteChapter(id int64) (err error) {
	// 1. 使用数据库删除
	err = s.Repo.DeleteChapter(id)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return err
	}
	// 2. 返回
	return nil
}

func (s NovelService) DeleteCharacter(id int64) (err error) {
	// 1. 使用数据库删除
	err = s.Repo.DeleteCharacter(id)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return err
	}
	// 2. 返回
	return nil
}

func (s NovelService) GetChapterList(novel_id int64) (chapters []entity.Chapter, err error) {
	// 1. 从数据库获取
	chapters, err = s.Repo.GetChapterList(novel_id)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return chapters, err
	}
	// 2. 返回
	return
}

func (s NovelService) GetCharacterList(novel_id int64) (characters []entity.Character, err error) {
	// 1. 从数据库获取
	characters, err = s.Repo.GetCharacterList(novel_id)
	if err != nil {
		zlog.ErrorfCtx(s.ctx, "数据库错误:%v", err)
		return characters, err
	}
	// 2. 返回
	return
}

package repo

import (
	"Ai-Novel/application/user/domain/entity"
)

func (r *UserRepo) ModifyUserProfile(user entity.User) (err error) {
	// 仅替换部分用户信息
	err = r.db.Model(&entity.User{}).Where("id =?", user.ID).Updates(map[string]interface{}{
		"username": user.Username,
		"avatar":   user.Avatar,
	}).Error
	return
}

package service

import (
	"encoding/base64"
	"errors"
	"goer/config"
	MyErr "goer/error"
	"goer/model"
	"goer/util"
	"goer/view"
	"os"
	"path/filepath"
	"time"
)

func CountClipForUser(userId int) (int, error) {
	var count int64
	queryRes := config.DataBase.Model(&model.UserClip{}).Where("user_id = ?", userId).Count(&count)
	if queryRes.Error != nil {
		return MyErr.IntErrValue, queryRes.Error
	}
	return int(count), nil
}

func DeleteUnusedClipForUser(userId int) error {
	var clip model.UserClip
	if err := config.DataBase.Where("user_id = ?", userId).Order("use_time asc").First(&clip).Error; err != nil {
		return err
	}
	if err := config.DataBase.Delete(&clip).Error; err != nil {
		return err
	}
	return nil
}

func PageUserClips(userId, pageNum, pageSize int) (*view.PageDataView, error) {
	// 计算总记录数
	var total int64
	if err := config.DataBase.Model(&model.UserClip{}).Where("user_id = ?", userId).Count(&total).Error; err != nil {
		return nil, err
	}
	totalInt := int(total)
	// 分页查询
	offset := util.Max((pageNum-1)*pageSize, totalInt-pageSize)
	var userClips []model.UserClip
	if err := config.DataBase.Model(&model.UserClip{}).
		Where("user_id = ?", userId).
		Order("use_time desc").
		Offset(offset).
		Limit(pageSize).
		Find(&userClips).Error; err != nil {
		return nil, err
	}
	// 对数据进行封装
	var data []interface{}
	for _, uc := range userClips {
		content := uc.Content
		if uc.Type == model.ClipImage {
			// 如果是图片数据，则转码为base64后返回
			imagePath := filepath.Join(config.ImageDirPath, uc.Content)
			contentTmp, err := os.ReadFile(imagePath)
			if err != nil {
				contentTmp = []byte("")
			}
			content = base64.StdEncoding.EncodeToString(contentTmp)
		}
		ucv := &view.UserClipView{
			ID:         uc.ID,
			Content:    content,
			Type:       uc.Type,
			CreateTime: uc.CreatedAt,
		}
		data = append(data, ucv)
	}
	return view.NewPageDataView(data, totalInt, pageNum, pageSize), nil
}

func DeleteUserClip(userId, clipId int) error {
	return config.DataBase.Where("user_id = ? and id = ?", userId, clipId).Delete(&model.UserClip{}).Error
}

func UpdateUserClip(userId, clipId int, content string) error {
	// 先将原本的数据查出
	var userClip model.UserClip
	if err := config.DataBase.Where("user_id = ? and id = ?", userId, clipId).First(&userClip).Error; err != nil {
		return err
	}
	// 检查是否为文本类型的数据
	if userClip.Type != model.ClipText {
		return errors.New(MyErr.DataCannotModify)
	}
	return config.DataBase.Model(&model.UserClip{}).
		Where("user_id = ? and id = ?", userId, clipId).
		Update("content", content).
		Update("use_time", time.Now()).
		Error
}

func UpdateUserClipUseTime(userId, clipId int) error {
	return config.DataBase.Model(&model.UserClip{}).
		Where("user_id = ? and id = ?", userId, clipId).
		Update("use_time", time.Now()).
		Error
}

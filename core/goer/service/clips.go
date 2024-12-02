package service

import (
	"github.com/robfig/cron/v3"
	"goer/config"
	"goer/model"
	"log"
	"time"
)

func GetClipByIdentifier(identifier string) (clip *model.Clip, err error) {
	clip = &model.Clip{}
	err = config.DataBase.Where("identifier = ?", identifier).First(clip).Error
	return
}

func TimedDeleteExpiredClips() {
	// 创建调度器
	c := cron.New()
	// 添加定时任务
	//time4AM := "0 4 * * *" // 每日凌晨4点 上线之后使用
	time2Min := "*/2 * * * *" // 每两分钟 测试用
	_, err := c.AddFunc(time2Min, func() {
		deleteExpiredClips()
	})
	if err != nil {
		log.Fatalf("failed to add cron job: %v", err)
	}
	// 启动调度器
	c.Start()
	// 防止主程序退出
	select {}
}

func deleteExpiredClips() {
	var clips []model.Clip
	// 过期时间是从现在向前推的
	expireTime := time.Now().Add(-time.Duration(config.NormalClipExpireDays) * 24 * time.Hour)
	if err := config.DataBase.Where("created_at < ?", expireTime).Find(&clips).Error; err != nil {
		// 直接查询
		log.Panicf("failed to find expired clips: %v", err)
		return
	}
	// 进行删除
	for _, clip := range clips {
		if err := config.DataBase.Delete(&clip).Error; err != nil {
			log.Panicf("failed to delete expired clip: %v", err)
		} else {
			log.Printf("deleted expired clip: %v", clip)
		}
	}
}

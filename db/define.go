package db

import "time"

type Messages struct {
	Id        int32     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	UserId    int32     `gorm:"column:user_id;NOT NULL"`
	Content   string    `gorm:"column:content;NOT NULL"`
	Tag       int8      `gorm:"column:tag;default:0;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
}

type Templates struct {
	Id        int32     `gorm:"column:id;AUTO_INCREMENT;NOT NULL"`
	Content   string    `gorm:"column:content;NOT NULL"`
	CreatorId int32     `gorm:"column:creator_id;NOT NULL"`
	Category  int8      `gorm:"column:category;default:0;NOT NULL"`
	Tag       int8      `gorm:"column:tag;default:0;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
}

var TagMp = map[string]int8{
	"Happy.":       1,
	"Angry.":       2,
	"Complaining.": 3,
	"Praising.":    4,
	"Flat.":        5,
}

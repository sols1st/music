package model

import (
	"time"
)

// User 用户模型
type User struct {
	ID           uint      `gorm:"column:id;primarykey" json:"id"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`
	Username     string    `gorm:"column:username;type:varchar(50);uniqueIndex;not null" json:"username"`
	Password     string    `gorm:"column:password;type:varchar(100);not null" json:"password"`
	Sex          uint8     `gorm:"column:sex;type:tinyint" json:"sex"`
	PhoneNum     string    `gorm:"column:phone_num;type:varchar(20)" json:"phoneNum"`
	Email        string    `gorm:"column:email;type:varchar(100);uniqueIndex" json:"email"`
	Birth        time.Time `gorm:"column:birth;type:date" json:"birth"`
	Introduction string    `gorm:"column:introduction;type:text" json:"introduction"`
	Location     string    `gorm:"column:location;type:varchar(100)" json:"location"`
	Avatar       string    `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
}

func (*User) TableName() string {
	return "user"
}

// Song 歌曲模型
type Song struct {
	ID           uint      `gorm:"column:id;primarykey" json:"id"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`
	SingerID     uint      `gorm:"column:singer_id;not null" json:"singerId"`
	Name         string    `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Introduction string    `gorm:"column:introduction;type:text" json:"introduction"`
	Pic          string    `gorm:"column:pic;type:varchar(255)" json:"pic"`
	Lyric        string    `gorm:"column:lyric;type:text" json:"lyric"`
	Url          string    `gorm:"column:url;type:varchar(255);not null" json:"url"`
}

func (*Song) TableName() string {
	return "song"
}

// Comment 评论模型
type Comment struct {
	ID         uint      `gorm:"column:id;primarykey" json:"id"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	Content    string    `gorm:"column:content;type:text;not null" json:"content"`
	UserID     uint      `gorm:"column:user_id" json:"userId"`
	SongID     uint      `gorm:"column:song_id" json:"songId"`
	SongListID uint      `gorm:"column:song_list_id" json:"songListId"`
	Type       byte      `gorm:"column:type" json:"type"`
	Up         uint      `gorm:"column:up;type:int" json:"up"`
}

func (*Comment) TableName() string {
	return "comment"
}

// Singer 歌手模型
type Singer struct {
	ID           uint      `gorm:"column:id;primarykey" json:"id"`
	Name         string    `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Sex          uint8     `gorm:"column:sex;type:tinyint" json:"sex"`
	Pic          string    `gorm:"column:pic;type:varchar(255)" json:"pic"`
	Birth        time.Time `gorm:"column:birth;type:date" json:"birth"`
	Location     string    `gorm:"column:location;type:varchar(100)" json:"location"`
	Introduction string    `gorm:"column:introduction;type:text" json:"introduction"`
}

func (*Singer) TableName() string {
	return "singer"
}

// SongList 歌单模型
type SongList struct {
	ID           uint      `gorm:"column:id;primarykey" json:"id"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`
	Title        string    `gorm:"column:title;type:varchar(100);not null" json:"title"`
	Pic          string    `gorm:"column:pic;type:varchar(255)" json:"pic"`
	Introduction string    `gorm:"column:introduction;type:text" json:"introduction"`
	Style        string    `gorm:"column:style;type:varchar(50)" json:"style"`
}

func (*SongList) TableName() string {
	return "song_list"
}

// ListSong 歌单歌曲关联模型
type ListSong struct {
	ID         uint `gorm:"column:id;primarykey" json:"id"`
	SongID     uint `gorm:"column:song_id;not null" json:"songId"`
	SongListID uint `gorm:"column:song_list_id;not null" json:"songListId"`
}

func (*ListSong) TableName() string {
	return "list_song"
}

// Collect 收藏模型
type Collect struct {
	ID         uint      `gorm:"column:id;primarykey" json:"id"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UserID     uint      `gorm:"column:user_id;not null" json:"userId"`
	SongID     uint      `gorm:"column:song_id" json:"songId"`
	SongListID uint      `gorm:"column:song_list_id" json:"songListId"`
	Type       uint8     `gorm:"column:type;type:tinyint;not null" json:"type"`
}

func (*Collect) TableName() string {
	return "collect"
}

// Admin 管理员模型
type Admin struct {
	ID       uint   `gorm:"column:id;primarykey" json:"id"`
	Name     string `gorm:"column:name;type:varchar(50);not null" json:"username"`
	Password string `gorm:"column:password;type:varchar(100);not null" json:"password"`
}

func (*Admin) TableName() string {
	return "admin"
}

// Banner 轮播图模型
type Banner struct {
	ID  uint   `gorm:"column:id;primarykey" json:"id"`
	Pic string `gorm:"column:pic;type:varchar(255);not null" json:"pic"`
}

func (*Banner) TableName() string {
	return "banner"
}

type Order struct {
	ID       uint   `gorm:"column:id;primarykey" json:"id"`
	Name     string `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Password string `gorm:"column:password;type:varchar(100);not null" json:"password"`
}

func (*Order) TableName() string {
	return "order"
}

// ResetPasswordRequest 重置密码请求模型
type ResetPasswordRequest struct {
	Email           string `json:"email" gorm:"column:email;type:varchar(100)"`
	Code            string `json:"code" gorm:"column:code;type:varchar(10)"`
	Password        string `json:"password" gorm:"column:password;type:varchar(100)"`
	ConfirmPassword string `json:"confirmPassword" gorm:"column:confirm_password;type:varchar(100)"`
}

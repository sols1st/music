package model

import "time"

type Config struct {
	ID               uint      `gorm:"primary_key" json:"id"`
	MaxLoginAttempts int       `json:"max_login_attempts"` // 最大登录尝试次数
	LockoutDuration  int       `json:"lockout_duration"`   // 锁定持续时间（分钟）
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

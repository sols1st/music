package service

import (
	"music-server-gin/global"
	"music-server-gin/model"
)

type BannerService struct{}

var BannerServiceApp = new(BannerService)

func (s *BannerService) AddBanner(banner *model.Banner) error {
	return global.DB.Create(banner).Error
}

func (s *BannerService) UpdateBannerMsg(banner *model.Banner) error {
	return global.DB.Model(&model.Banner{}).Where("id = ?", banner.ID).Updates(banner).Error
}

func (s *BannerService) UpdateBannerPic(id uint, pic string) error {
	return global.DB.Model(&model.Banner{}).Where("id = ?", id).Update("pic", pic).Error
}

func (s *BannerService) DeleteBanner(id uint) error {
	return global.DB.Delete(&model.Banner{}, id).Error
}

func (s *BannerService) AllBanner() ([]model.Banner, error) {
	var banners []model.Banner
	err := global.DB.Find(&banners).Error
	return banners, err
}

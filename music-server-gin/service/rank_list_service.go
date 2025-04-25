package service

import (
	"music-server-gin/global"
	"music-server-gin/model"
)

type RankListService struct{}

func (s *RankListService) AddRank(rankList *model.RankList) error {
	return global.DB.Create(rankList).Error
}

func (s *RankListService) RankOfSongListId(songListId uint) ([]model.RankList, error) {
	var rankLists []model.RankList
	err := global.DB.Where("song_list_id = ?", songListId).Find(&rankLists).Error
	return rankLists, err
}

func (s *RankListService) GetUserRank(songListId, consumerId uint) (model.RankList, error) {
	var rankList model.RankList
	err := global.DB.Where("song_list_id = ? AND consumer_id = ?", songListId, consumerId).First(&rankList).Error
	return rankList, err
}

func (s *RankListService) GetAverageScore(songListId uint) (float64, error) {
	var avg float64
	err := global.DB.Model(&model.RankList{}).Where("song_list_id = ?", songListId).Select("AVG(score)").Row().Scan(&avg)
	return avg, err
}

var RankListServiceApp = new(RankListService)

package dto

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/entity"
	"DouyinSimpleProject/utils"
	"time"
)

type VideoDTO struct {
	ID            uint        `json:"id"`
	Author        UserInfoDTO `json:"author"`
	PlayURL       string      `json:"play_url"`
	CoverURL      string      `json:"cover_url"`
	FavoriteCount uint        `json:"favorite_count"`
	CommentCount  uint        `json:"comment_count"`
	IsFavorite    bool        `json:"is_favorite"`
	Title         string      `json:"title"`
	CreatedAt     time.Time   `json:"-"`
}

// VideoResponse responses to `/feed/` or `/publish/list`
type VideoResponse struct {
	Response
	NextTime  int64       `json:"next_time,omitempty"`
	VideoList []*VideoDTO `json:"video_list"`
}

// NewVideoDTO creates an instance of VideoDTO
// uid is the LoginUser
func NewVideoDTO(video *entity.Video, loginUID uint) *VideoDTO {
	isFavorite := false
	if loginUID != 0 { // no login user
		fq := dao.Q.Favorite
		cnt, err := fq.Where(fq.UserID.Eq(loginUID)).Where(fq.VideoID.Eq(video.ID)).Count()
		if err == nil && cnt != 0 {
			isFavorite = true
		}
	}
	author := NewUserInfoDTO(&video.User, loginUID)
	return &VideoDTO{
		ID:            video.ID,
		Author:        *author,
		PlayURL:       utils.GetFileURL(video.PlayURL),
		CoverURL:      utils.GetFileURL(video.CoverURL),
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    isFavorite,
		Title:         video.Title,
		CreatedAt:     video.CreatedAt,
	}
}

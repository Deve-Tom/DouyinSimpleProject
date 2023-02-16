package dto

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/entity"
	"DouyinSimpleProject/utils"
	"time"
)

type AuthorDTO struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   uint   `json:"follow_count"`
	FollowerCount uint   `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type VideoDTO struct {
	ID            uint      `json:"id"`
	Author        AuthorDTO `json:"author"`
	PlayURL       string    `json:"play_url"`
	CoverURL      string    `json:"cover_url"`
	FavoriteCount uint      `json:"favorite_count"`
	CommentCount  uint      `json:"comment_count"`
	IsFavorite    bool      `json:"is_favorite"`
	Title         string    `json:"title"`
	CreatedAt     time.Time `json:"-"`
}

// VideoResponse responses to `/feed/` or `/publish/list`
type VideoResponse struct {
	Response
	NextTime  int64       `json:"next_time,omitempty"`
	VideoList []*VideoDTO `json:"video_list"`
}

// NewVideoDTO creates an instance of VideoDTO
// uid is the LoginUser
func NewVideoDTO(video *entity.Video, uid uint) *VideoDTO {
	isFavorite := false
	// TODO: get isFollow
	isFollow := false
	if uid != 0 { // no login user
		fq := dao.Q.Favorite
		cnt, err := fq.Where(fq.UserID.Eq(uid)).Where(fq.VideoID.Eq(video.ID)).Count()
		if err == nil && cnt != 0 {
			isFavorite = true
		}
	}
	return &VideoDTO{
		ID: video.ID,
		Author: AuthorDTO{
			ID:            video.User.ID,
			Name:          video.User.Nickname,
			FollowCount:   video.User.FollowCount,
			FollowerCount: video.User.FollowerCount,
			IsFollow:      isFollow,
		},
		PlayURL:       utils.GetFileURL(video.PlayURL),
		CoverURL:      utils.GetFileURL(video.CoverURL),
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    isFavorite,
		Title:         video.Title,
		CreatedAt:     video.CreatedAt,
	}
}

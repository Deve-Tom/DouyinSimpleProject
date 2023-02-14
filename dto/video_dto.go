package dto

import "time"

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
	PlayURL       string    `json:"paly_url"`
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

package service

import (
	"DouyinSimpleProject/config"
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/entity"
	"DouyinSimpleProject/utils"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var videoSuffixMap = map[string]struct{}{
	".mp4":  {},
	".avi":  {},
	".wmv":  {},
	".flv":  {},
	".mov":  {},
	".mpeg": {},
}

type VideoService interface {
	GetVideoList(user_id uint) []dto.VideoDTO
	Publish(ctx *gin.Context, uid uint, title string, videoFile *multipart.FileHeader) (string, bool)
	getVideoName(uid uint) string
}

type videoService struct {
}

func NewVideoService() VideoService {
	return &videoService{}
}

func (s *videoService) GetVideoList(user_id uint) []dto.VideoDTO {
	vq := dao.Q.Video
	videos, err := vq.Where(vq.UserID.Eq(user_id)).Find()
	if err != nil {
		return nil
	}
	videoDTOList := make([]dto.VideoDTO, len(videos))
	// TODO
	isFollow := true
	isFavorite := true
	for i, video := range videos {
		videoDTOList[i] = dto.VideoDTO{
			ID: video.ID,
			Author: dto.AuthorDTO{
				ID:            video.User.ID,
				Name:          video.User.Nickname,
				FollowCount:   video.User.FollowCount,
				FollowerCount: video.User.FollowerCount,
				IsFollow:      isFollow,
			},
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		}
	}
	return videoDTOList
}

func (s *videoService) Publish(ctx *gin.Context, uid uint, title string, videoFile *multipart.FileHeader) (string, bool) {
	// check video type
	suffix := filepath.Ext(videoFile.Filename)
	if _, ok := videoSuffixMap[suffix]; !ok {
		return "Unsupported video type", false
	}

	// save uploaded video
	videoName := s.getVideoName(uid)
	videoFileName := videoName + suffix
	videoPath := filepath.Join(config.STATIC_ROOT_PATH, videoFileName)
	if err := ctx.SaveUploadedFile(videoFile, videoPath); err != nil {
		return "Save Uploaded File error: " + err.Error(), false
	}

	// extract cover image from video
	coverFilename := utils.ExtractImageFromVideo(videoName, suffix)

	// insert video
	vq := dao.Q.Video
	err := vq.Create(&entity.Video{
		UserID:   uid,
		Title:    title,
		PlayURL:  utils.GetFileURL(videoFileName),
		CoverURL: utils.GetFileURL(coverFilename),
	})
	if err != nil {
		return err.Error(), false
	}
	return "Successfully publish a video", true
}

func (s *videoService) getVideoName(uid uint) string {
	vq := dao.Q.Video
	videoCount, _ := vq.Where(vq.UserID.Eq(uid)).Count()
	videoName := fmt.Sprintf("%d-%d", uid, videoCount+1)
	return videoName
}

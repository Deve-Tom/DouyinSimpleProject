package service

import (
	"DouyinSimpleProject/config"
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/entity"
	"DouyinSimpleProject/utils"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

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
	GetVideoDTOList(limitNum int, latestTime time.Time, uid uint) ([]*dto.VideoDTO, error)
	Publish(ctx *gin.Context, uid uint, title string, videoFile *multipart.FileHeader) error
	genVideoName(uid uint) string
	getVideoList(limitNum int, latestTime time.Time, uid uint) ([]*entity.Video, error)
}

type videoService struct {
}

func NewVideoService() VideoService {
	return &videoService{}
}

// GetVideoDTOList gets a videoDTO list from database according to limitNum, latestTime and uid
func (s *videoService) GetVideoDTOList(limitNum int, latestTime time.Time, uid uint) ([]*dto.VideoDTO, error) {
	videos, err := s.getVideoList(limitNum, latestTime, uid)
	if err != nil {
		return nil, err
	}

	videoDTOList := make([]*dto.VideoDTO, len(videos))
	// TODO: implement `isFollow` and `isFavorite`
	isFollow := true
	isFavorite := true
	for i, video := range videos {
		videoDTOList[i] = dto.NewVideoDTO(video, isFavorite, isFollow)
	}
	return videoDTOList, nil
}

// Publish creates a video and save it into database
func (s *videoService) Publish(ctx *gin.Context, uid uint, title string, videoFile *multipart.FileHeader) error {
	// check video type
	suffix := filepath.Ext(videoFile.Filename)
	if _, ok := videoSuffixMap[suffix]; !ok {
		return errors.New("unsupported video type")
	}

	// save uploaded video
	videoName := s.genVideoName(uid)
	videoFileName := videoName + suffix
	videoPath := filepath.Join(config.STATIC_ROOT_PATH, videoFileName)
	if err := ctx.SaveUploadedFile(videoFile, videoPath); err != nil {
		return errors.New("save uploaded file error")
	}

	// extract cover image from video
	coverFilename := utils.ExtractImageFromVideo(videoName, suffix)

	// insert video
	vq := dao.Q.Video
	err := vq.Create(&entity.Video{
		UserID:   uid,
		Title:    title,
		PlayURL:  videoFileName,
		CoverURL: coverFilename,
	})
	return err
}

// genVideoName generate video name, the format is `{user_id}-{videoCount+1}`
func (s *videoService) genVideoName(uid uint) string {
	vq := dao.Q.Video
	videoCount, _ := vq.Where(vq.UserID.Eq(uid)).Count()
	videoName := fmt.Sprintf("%d-%d", uid, videoCount+1)
	return videoName
}

// getVideoList retrieves videos from database
func (s *videoService) getVideoList(limitNum int, latestTime time.Time, uid uint) ([]*entity.Video, error) {
	vq := dao.Q.Video
	_vq := vq.Preload(vq.User)
	if uid != 0 {
		_vq = _vq.Where(vq.UserID.Eq(uid))
	}
	videos, err := _vq.Where(vq.CreatedAt.Lte(latestTime)).
		Order(vq.CreatedAt.Desc()).
		Limit(limitNum).
		Find()
	if err != nil {
		return nil, err
	}
	return videos, nil
}

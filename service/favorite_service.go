package service

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/entity"
	"errors"
)

type FavoriteService interface {
	Action(uid, vid, actionType uint) error
	DO(uid, vid uint) error
	Cancel(uid, vid uint) error
	GetFavoriteList(uid uint) ([]*dto.VideoDTO, error)
}

type favoriteService struct{}

func NewFavoriteService() FavoriteService {
	return &favoriteService{}
}

func (s *favoriteService) Action(uid, vid, actionType uint) error {
	if actionType == 1 {
		return s.DO(uid, vid)
	} else if actionType == 2 {
		return s.Cancel(uid, vid)
	} else {
		return errors.New("invalid action type")
	}
}

func (s *favoriteService) DO(uid, vid uint) error {
	fq := dao.Q.Favorite
	cnt, err := fq.Where(fq.UserID.Eq(uid)).Where(fq.VideoID.Eq(vid)).Count()
	if err != nil {
		return err
	}
	if cnt != 0 {
		return errors.New("repeat thumbs up")
	}
	err = fq.Create(&entity.Favorite{
		UserID:  uid,
		VideoID: vid,
	})
	if err != nil {
		return errors.New("no such video")
	}
	return nil
}

func (s *favoriteService) Cancel(uid, vid uint) error {
	fq := dao.Q.Favorite
	// use Unscoped to hard delete, not soft delete
	_, err := fq.Unscoped().Where(fq.UserID.Eq(uid)).Where(fq.VideoID.Eq(vid)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *favoriteService) GetFavoriteList(uid uint) ([]*dto.VideoDTO, error) {

	fq := dao.Q.Favorite
	favorites, err := fq.Where(fq.UserID.Eq(uid)).Find()
	if err != nil {
		return nil, err
	}
	vids := make([]uint, len(favorites))
	for i, f := range favorites {
		vids[i] = f.VideoID
	}
	vq := dao.Q.Video
	videos, err := vq.Preload(vq.User).Where(vq.ID.In(vids...)).Order(vq.CreatedAt.Desc()).Find()
	if err != nil {
		return nil, err
	}
	// TODO: isFollow
	videoDTOList := make([]*dto.VideoDTO, len(videos))
	for i, v := range videos {
		videoDTOList[i] = dto.NewVideoDTO(v, true, true)
	}
	return videoDTOList, nil
}

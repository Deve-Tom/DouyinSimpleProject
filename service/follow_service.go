package service

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/dto"
	"DouyinSimpleProject/entity"
	"errors"
)

type FollowService interface {
	Action(uid, fuid, actionType uint) error
	DO(uid, fuid uint) error
	Cancel(uid, fuid uint) error

	GetFollowList(uid uint, isFollow bool) ([]*dto.UserInfoDTO, error)

}

type followService struct{}

func NewFollowService() FollowService {
	return &followService{}
}
func (s *followService) Action(uid, fuid, actionType uint) error {

	if uid == fuid {
		return errors.New("can not follow yourself")
	}

	if actionType == 1 {
		return s.DO(uid, fuid)
	} else if actionType == 2 {
		return s.Cancel(uid, fuid)
	} else {
		return errors.New("invalid action type")
	}
}

func (s *followService) DO(uid, fuid uint) error {
	fq := dao.Q.Follow
	cnt, err := fq.Where(fq.UserID.Eq(uid)).Where(fq.FollowUserID.Eq(fuid)).Count()
	if err != nil {
		return err
	}
	if cnt != 0 {
		return errors.New("repeat follow")
	}
	//use transaction to do follow
	err = dao.Q.Transaction(func(tx *dao.Query) error {
		follow := entity.Follow{
			UserID:       uid,
			FollowUserID: fuid,
		}
		if err := tx.Follow.Create(&follow); err != nil {
			return err
		}
		//update user.follow_count
		if _, err := tx.User.Where(tx.User.ID.Eq(uid)).UpdateSimple(tx.User.FollowCount.Add(1)); err != nil {
			return err
		}
		//update fuser.follower.count
		if _, err := tx.User.Where(tx.User.ID.Eq(fuid)).UpdateSimple(tx.User.FollowerCount.Add(1)); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("no such user")
	}
	return nil
}

func (s *followService) Cancel(uid, fuid uint) error {
	//use transaction to cancel follow
	err := dao.Q.Transaction(func(tx *dao.Query) error {
		fq := tx.Follow
		//use Unscoped to hard delete, not soft delete
		if _, err := fq.Unscoped().Where(fq.UserID.Eq(uid)).Where(fq.FollowUserID.Eq(fuid)).Delete(); err != nil {
			return err
		}
		//update user.follow_count
		if _, err := tx.User.Where(tx.User.ID.Eq(uid)).UpdateSimple(tx.User.FollowCount.Sub(1)); err != nil {
			return err
		}
		//update fuser.follower.count
		if _, err := tx.User.Where(tx.User.ID.Eq(fuid)).UpdateSimple(tx.User.FollowerCount.Sub(1)); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// get followlist(true) & get followerlist(false)
func (s *followService) GetFollowList(uid uint, isFollow bool) ([]*dto.UserInfoDTO, error) {
	uq := dao.Q.User
	fq := dao.Q.Follow
	fq_field := fq.UserID
	if !isFollow {
		fq_field = fq.FollowUserID
	}

	user_follows, err := fq.Where(fq_field.Eq(uid)).Find()
	if err != nil {
		return nil, err
	}
	users := make([]*entity.User, len(user_follows))

	for i, user_follow := range user_follows {
		fq_uint := user_follow.FollowUserID
		if !isFollow {
			fq_uint = user_follow.UserID
		}

		//TODO:sql optimization
		users[i], err = uq.Where(uq.ID.Eq(fq_uint)).First()
		if err != nil {
			return nil, err
		}
	}

	UserDTOList := make([]*dto.UserInfoDTO, len(users))
	for i, user := range users {
		UserDTOList[i] = dto.NewUserInfoDTO(user, uid, true)
	}

	return UserDTOList, nil

}

package service

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/entity"
	"errors"
)

type FollowService interface {
	Action(uid, fuid, actionType uint) error
	DO(uid, fuid uint) error
	Cancel(uid, fuid uint) error
}

type followService struct{}

func NewFollowService() FollowService {
	return &followService{}
}
func (s *followService) Action(uid, fuid, actionType uint) error {
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

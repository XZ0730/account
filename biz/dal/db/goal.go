package db

import (
	"time"
)

type Goal struct {
	UserId     int64
	GoalId     int64
	GoalName   string
	Money      float64
	CreateDate time.Time
	Deadline   time.Time
}

func NewGoal(userid int64, name string, money float64, c_time, deadline time.Time) *Goal {
	return &Goal{
		UserId:     userid,
		GoalName:   name,
		Money:      money,
		CreateDate: c_time,
		Deadline:   deadline,
	}
}

func CreateGoal(goal *Goal) error {
	return DB.Table("t_goal").Create(&goal).Error
}

func GetGoalList(userid int64) ([]Goal, error) {
	goals := make([]Goal, 0)
	err := DB.Table("t_goal").Where("user_id=?", userid).Find(&goals).Error
	if err != nil {
		return nil, err
	}
	return goals, nil
}

func DelGoal(userid int64, goalid int64) error {
	return DB.Table("t_goal").Where("user_id=? AND goal_id=?", userid, goalid).Unscoped().Delete(&Goal{}).Error
}

func UpdateGoal(goal *Goal) error {
	return DB.Table("t_goal").Where("user_id=? AND goal_id=?", goal.UserId, goal.GoalId).
		Updates(&Goal{GoalName: goal.GoalName, CreateDate: goal.CreateDate, Deadline: goal.Deadline, Money: goal.Money}).Error
}

package service

import (
	"TrainSystem/entity"
	"errors"
)


type ScheduleCache map[int]*entity.Schedule

func NewSchduleCache() ScheduleCache{
	return make(map[int]*entity.Schedule)
}

func(c ScheduleCache) Schedule(id int)(*entity.Schedule, error){
	if cat, ok := c[id]; ok {
		return cat, nil
	}
	return nil, errors.New("Catagory was not found")
}

func (c ScheduleCache) StoreSchedule(Schedule *entity.Schedule) error {
	if _, ok := c[Schedule.ID]; !ok {
		c[Schedule.ID]= Schedule
		return nil
	}
	return errors.New("Catagory already exists")
}
package service

import (
	"errors"

	"github.com/betsegawlemma/web-prog-lab-05-mem/entity"
)

// CategoryCache provides an in-memory cache
type ScheduleCache map[int]*entity.Schedule

// NewCategoryCache returns a new category cache
func NewScheduleCache() ScheduleCache {
	return make(map[int]*entity.Schedule)
}

// Category returns a category for a given id from the cache
func (c ScheduleCache) Schedule(id int) (*entity.Schedule, error) {
	if cat, ok := c[id]; ok {
		return cat, nil
	}
	return nil, errors.New("Schedule was not found")
}

// StoreCategory stores category data to the cache
func (c ScheduleCache) StoreSchedule(Schedule *entity.Schedule) error {
	if _, ok := c[Schedule.ID]; !ok {
		c[Schedule.ID] = Schedule
		return nil
	}
	return errors.New("Category already exists")
}

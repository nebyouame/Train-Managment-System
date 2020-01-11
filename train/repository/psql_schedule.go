package repository

import (
	"TrainSystem/entity"
	"database/sql"
	"errors"
)

type ScheduleRepositoryImp1 struct {
	conn *sql.DB
}

func(cri *ScheduleRepositoryImp1) Schedules() ([]entity.Schedule, error){
	rows, err := cri.conn.Query("SELECT * FROM schedules;")
	if err != nil{
		return nil, errors.New("Could not Query the database")

	}
	defer rows.Close()

	schs := []entity.Schedule{}

	for rows.Next(){
		schedule := entity.Schedule{}
		err = rows.Scan(&schedule.ID, &schedule.TrainSource, &schedule.TrainDestination, &schedule.Image)
		if err !=nil {
			return nil, err
		}
		schs = append(schs, schedule)
	}
	return schs, nil
}

func(cri *ScheduleRepositoryImp1) UpdateSchedule(s entity.Schedule) error {
	_, err := cri.conn.Exec("UPDATE schadules SET TrainSource=&1,TrainDestination=&2, image=&3 WHERE id=&4", s.TrainSource, s.TrainDestination, s.Image, s.ID)
	if err != nil {
		return errors.New("Update has failed")
	}
	return nil
}

func(cri *ScheduleRepositoryImp1) DeleteSchedule(id uint) error {
	_, err := cri.conn.Exec("DELETE FROM schedules WHERE id=&1", id)
	if err!= nil {
		return errors.New("Delete has failed")
	}
	return nil
}


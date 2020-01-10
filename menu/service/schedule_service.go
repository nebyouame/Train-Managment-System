package service

import (
	"bytes"

	"encoding/gob"
	"errors"

	"io/ioutil"
	
	"TrainSystem/entity"
)

type ScheduleService struct {
	FileName string
}

func NewScheduleService(fileName string) *ScheduleService {
	return &ScheduleService{FileName:fileName}
}
func (cs ScheduleService) Schedules() ([]entity.Schedule, error) {

	raw, err := ioutil.ReadFile(cs.FileName)
	if err != nil {
		return nil, errors.New("File could not be open")
	}
	buffer := bytes.NewBuffer(raw)

	dec := gob.NewDecoder(buffer)

	var ctgs []entity.Schedule

	err = dec.Decode(&ctgs)

	if err != nil {
		return nil, errors.New("Decoding error")
	}

	return ctgs, nil
}

func (cs ScheduleService) StoreSchedules(ctgs []entity.Schedule) error {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)

	err := encoder.Encode(ctgs)

	if err != nil {
		return errors.New("Data encoding has failed")
	}

	err = ioutil.WriteFile(cs.FileName, buffer.Bytes(), 0644)

	if err != nil {
		return errors.New("Writing to a file has failed")
	}

	return nil
}

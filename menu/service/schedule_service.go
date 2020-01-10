package service

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"TrainSystem/entity"
)

type ScheduleService struct {
	FileName string
}

func NewScheduleService(fileName string) *ScheduleService {
	return &ScheduleService{FileName:fileName}
}
func (cs ScheduleService) Categories() ([]entity.Schedule, error) {
	file, err := os.Open(cs.FileName)
	if err != nil {
		return nil, errors.New("File could not be open")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("File could not be open")
	}
	var ctgs []entity.Schedule
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		c := entity.Schedule{ID: int(id), TrainSource: item[1],
			TrainDestination: item[2], Image: item[3]}
		ctgs = append(ctgs, c)
	}
	return ctgs, nil
}

func (cs ScheduleService) StoreCategories(ctgs []entity.Schedule) error {
	csvFile, err := os.Create(cs.FileName)
	if err != nil {
		return errors.New("File could not be created")
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, c := range ctgs {
		line := []string{strconv.Itoa(c.ID), c.TrainSource, c.TrainDestination, c.Image}
		writer.Write(line)
	}
	writer.Flush()
	return nil
}

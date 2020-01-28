package entity

var ScheduleMock = Schedule{
	ID: 4,
	TrainSource: "Mock Schedule 01",
	TrainDestination: "Mock Schedule 01 Destination",
	Image : "mock_train.png",
	Infos: []Info{},
}


var InfoMock = Info {
	ID: 1,
	TrainSource: "Mock Info 01",
	Price: 100.0,
	TrainDestination: "Mock Info 01 Destination",
	Schedules: []Schedule{},
	Image: "mock_info.png",
	Paths: []Path{},
}

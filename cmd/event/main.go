package event

type TimeInterval struct {
	Start string
	End string
}

type Event struct {
	Id int
	Name string
	Description string
	Date TimeInterval
}


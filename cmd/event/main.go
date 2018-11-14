package event

type TimeInterval struct {
	Start string
	End string
}

type Event struct {
	Id int `bson:"id,omitempty" json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Date TimeInterval `json:"date"`
}


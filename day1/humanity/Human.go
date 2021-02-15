package humanity

type Human struct {
	Name    string
	age     int
	Country string
}

func NewHumanFromCSV(csv []string) (*Human, error) {
	return &Human{
		Name:    csv[0],
		age:     csv[1],
		Country: csv[2],
	}, nil
}

func NewHumansFromCsvFile(path string) ([]*Human, error) {
	for _, i := range()
}
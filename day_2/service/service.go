package service

import (
	"encoding/json"
	"sort"
)

var _ Decoder = (*Service)(nil)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type Place struct {
	City    string `json:"city"`
	Country string `json:"country"`
}
type Thing struct {
	*Person
	*Place
}
type Data struct {
	Things []Thing `json:"things"`
}
type Decoder interface {
	Decode(data []byte) ([]Person, []Place)
	Sort(dataToSort interface{})

	Print(interface{})
	Printlen(persons []Person, places []Place)
}

type Logger interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Service struct {
	log Logger
}

func New(log Logger) *Service {
	return &Service{log: log}
}
func (s *Service) Decode(data []byte) ([]Person, []Place) {
	var d Data
	if err := json.Unmarshal(data, &d); err != nil {
		s.log.Fatalf("decode data: ", err)
		return nil, nil
	}
	var persons []Person
	var places []Place
	for _, thing := range d.Things {
		if thing.Person != nil {
			persons = append(persons, *thing.Person)
		}
		if thing.Place != nil {
			places = append(places, *thing.Place)
		}
	}
	s.Printlen(persons, places)
	s.Sort(persons)
	s.Print(persons)
	s.Sort(places)
	s.Print(places)
	return persons, places
}

func (s *Service) Sort(dataToSort interface{}) {
	switch data := dataToSort.(type) {
	case []Person:
		sort.Slice(data, func(i, j int) bool { return data[i].Age > data[j].Age })
	case []Place:
		sort.Slice(data, func(i, j int) bool { return len(data[i].City) < len(data[j].City) })
	}
}
func (s *Service) Print(data interface{}) {
	s.log.Println(data)
}

func (s *Service) Printlen(persons []Person, places []Place) {
	s.log.Println(len(persons), len(places))
}

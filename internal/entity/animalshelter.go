package entity

import "github.com/mcherdakov/serialization-test/internal/gen/proto"

type EmployeeID int64

type Animal struct {
	Name   string
	Age    int64
	Weight float64
}

type Employee struct {
	FirstName  string
	SecondName string
	Manager    *EmployeeID
	Hobbies    []string
}

type AnimalShelter struct {
	Address     string
	PhoneNumber string
	Animals     []Animal
	Employees   []Employee
}

func NewAnimalShelter() AnimalShelter {
	return AnimalShelter{
		Address:     "Hogwarts Castle, Highlands, Scotland, Great Britain",
		PhoneNumber: "8-800-555-35-35",
		Animals: []Animal{
			{
				Name:   "Fawkes",
				Age:    100,
				Weight: 8,
			},
			{
				Name:   "Fang",
				Age:    10,
				Weight: 25,
			},
			{
				Name:   "Crookshanks",
				Age:    5,
				Weight: 10,
			},
			{
				Name:   "Mrs. Norris",
				Age:    15,
				Weight: 6,
			},
			{
				Name:   "Hedwig",
				Age:    5,
				Weight: 4,
			},
			{
				Name:   "Buckbeak",
				Age:    11,
				Weight: 325,
			},
		},
		Employees: []Employee{
			{
				FirstName:  "Albus",
				SecondName: "Dumbledore",
				Hobbies:    []string{"duelling", "magic items collecting"},
			},
			{
				FirstName:  "Harry",
				SecondName: "Potter",
				Manager:    ptr(EmployeeID(1)),
				Hobbies:    []string{"getting in trouble", "quidditch"},
			},
			{
				FirstName:  "Hermione",
				SecondName: "Granger",
				Manager:    ptr(EmployeeID(1)),
				Hobbies:    []string{"reading", "muggle studies"},
			},
			{
				FirstName:  "Ron",
				SecondName: "Weasley",
				Manager:    ptr(EmployeeID(3)),
				Hobbies:    []string{"quidditch", "chess"},
			},
		},
	}
}

func (shelter *AnimalShelter) ToProto() *proto.AnimalShelter {
	animals := make([]*proto.Animal, 0, len(shelter.Animals))
	for _, animal := range shelter.Animals {
		animals = append(animals, &proto.Animal{
			Name:   animal.Name,
			Age:    animal.Age,
			Weight: float32(animal.Weight),
		})
	}

	employyes := make([]*proto.Employee, 0, len(shelter.Employees))
	for _, employee := range shelter.Employees {
		employyes = append(employyes, &proto.Employee{
			FirstName:  employee.FirstName,
			SecondName: employee.SecondName,
			Manager:    (*int64)(employee.Manager),
			Hobbies:    employee.Hobbies,
		})
	}

	return &proto.AnimalShelter{
		Address:     shelter.Address,
		PhoneNumber: shelter.PhoneNumber,
		Animals:     animals,
		Employees:   employyes,
	}
}

func FromProto(shelter *proto.AnimalShelter) AnimalShelter {
	animals := make([]Animal, 0, len(shelter.Animals))
	for _, animal := range shelter.Animals {
		animals = append(animals, Animal{
			Name:   animal.Name,
			Age:    animal.Age,
			Weight: float64(animal.Weight),
		})
	}

	employyes := make([]Employee, 0, len(shelter.Employees))
	for _, employee := range shelter.Employees {
		employyes = append(employyes, Employee{
			FirstName:  employee.FirstName,
			SecondName: employee.SecondName,
			Manager:    (*EmployeeID)(employee.Manager),
			Hobbies:    employee.Hobbies,
		})
	}

	return AnimalShelter{
		Address:     shelter.Address,
		PhoneNumber: shelter.PhoneNumber,
		Animals:     animals,
		Employees:   employyes,
	}
}

func ptr[T any](val T) *T {
	return &val
}

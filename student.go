package student

import (
	"errors"
	"strconv"
)

type Student struct {
	Name   string
	Age    int
	Course int
}

func (s Student) Info() string {
	return s.Name + ", " + strconv.Itoa(s.Age) + ", " + strconv.Itoa(s.Course)
}

func Put(list map[string]*Student, value *Student) (int, error) {
	list[value.Name] = value

	if list[value.Name] == nil {
		return -1, errors.New("Ошибка добавления данных в список!")
	} else {
		return 0, nil
	}
}

func Get(list map[string]*Student, name string) (*Student, error) {
	if list[name] == nil {
		return nil, errors.New("Студент в списке не найден!")
	} else {
		return list[name], nil
	}
}

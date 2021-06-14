package filter

type Gender int

const (
	Male   Gender = 0
	Female Gender = 1
)

type Person struct {
	Name   string
	Age    int
	Gender Gender
}

type Criteria interface {
	Filter([]Person) []Person
}

type MaleCriteria struct{}

func (m *MaleCriteria) Filter(persons []Person) []Person {
	var res []Person
	for _, v := range persons {
		if v.Gender == Male {
			res = append(res, v)
		}
	}
	return res
}

type FemaleCriteria struct{}

func (m *FemaleCriteria) Filter(persons []Person) []Person {
	var res []Person
	for _, v := range persons {
		if v.Gender == Female {
			res = append(res, v)
		}
	}
	return res
}

type AgeGreaterCriteria struct {
	Age int
}

func (a *AgeGreaterCriteria) Filter(persons []Person) []Person {
	var res []Person
	for _, v := range persons {
		if v.Age > a.Age {
			res = append(res, v)
		}
	}
	return res
}

type AgeLessCriteria struct {
	Age int
}

func (a *AgeLessCriteria) Filter(persons []Person) []Person {
	var res []Person
	for _, v := range persons {
		if v.Age < a.Age {
			res = append(res, v)
		}
	}
	return res
}

type AndCriteria struct {
	Left, Right Criteria
}

func (a *AndCriteria) Filter(persons []Person) []Person {
	return a.Left.Filter(a.Right.Filter(persons))
}

type OrCriteria struct {
	Left, Right Criteria
}

func (a *OrCriteria) Filter(persons []Person) []Person {
	x := a.Left.Filter(persons)
	y := a.Right.Filter(persons)
	m := make(map[Person]bool)
	for _, v := range x {
		m[v] = true
	}

	for _, v := range y {
		m[v] = true
	}

	var res []Person

	for v := range m {
		res = append(res, v)
	}
	return res
}

type PersonCollection []Person

func (p PersonCollection) Where(c Criteria) PersonCollection {
	persons := c.Filter(p)
	return persons
}

package filter

import (
	"fmt"
)

func ExampleFilter() {
	//Output:
	//[{ff 18 0}]
	var persons []Person
	persons = append(persons, Person{"aa", 31, Female})
	persons = append(persons, Person{"bb", 22, Male})
	persons = append(persons, Person{"cc", 23, Female})
	persons = append(persons, Person{"dd", 33, Male})
	persons = append(persons, Person{"ee", 40, Female})
	persons = append(persons, Person{"ff", 18, Male})

	var pc PersonCollection = persons

	var res []Person = pc.Where(&MaleCriteria{}).Where(&AgeLessCriteria{20})

	fmt.Println(res)
}

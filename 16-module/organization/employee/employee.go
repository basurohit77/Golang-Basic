package employee

import . "testmodel/organization"

type Employee struct {
	Id   int
	Name string
	Org  Organization
}

func NewEmployee(id int, name string, orgName string, orgCity string) *Employee {
	return &Employee{
		Id:   id,
		Name: name,
		Org: Organization{
			Name: orgName,
			City: orgCity,
		},
	}
}


package census

type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

func (r *Resident) HasRequiredInfo() bool {
	return r.Name != "" && r.Address["street"] != ""
}

func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

func Count(residents []*Resident) int {
	var count int
	for _, v := range residents {
		if v.HasRequiredInfo() {
			count++
		}
	}

	return count
}

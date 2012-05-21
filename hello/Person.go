package hello

type Person struct{
	First string
	Last string
}

func NewPerson(first, last string) *Person{
	p := new(Person)
	p.First = first
	p.Last = last

	return p
}

func (p *Person) Name() string{
	return p.First + " " + p.Last
}
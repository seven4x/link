package learn

type Person struct {
	name string
}

func (p *Person) Hi(context string)  {
	println(context)
}
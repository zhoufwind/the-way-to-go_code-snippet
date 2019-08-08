package person2

type Person struct {
    firstName	string
    lastName	string
}

// Get方法
func (p *Person) FirstName() string {
    return p.firstName
}

// Set方法
func (p *Person) SetFirstName(newName string) {
    p.firstName = newName
}

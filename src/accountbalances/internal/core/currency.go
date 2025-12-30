package core

type Currency struct {
	ID   uint
	Code string
}

func (c *Currency) IsValidCode() bool {
	return len(c.Code) == 3
}

package numbers

type PrimaryPort interface {
	AddNumber(string) error
	UpdateNumber() error
	ListAvailableNumbers() ([]string, error)
}

type SecondaryPort interface {
	AddNumber(string) error
	UpdateNumber() error
	ListAvailableNumbers() ([]string, error)
}

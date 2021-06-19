package numbers

type PrimaryPort interface {
	UpdateNumber() error
}

type SecondaryPort interface {
	UpdateNumber() error
}

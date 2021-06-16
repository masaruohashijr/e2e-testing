package calls

type PrimaryPort interface {
	MakeCall() error
}

type SecondaryPort interface {
	MakeCall() error
}

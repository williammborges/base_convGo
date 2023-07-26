package bases

type Base int

const (
	Binary Base = iota + 1
	Octal
	Decimal
	Hexadecimal
)

func (s Base) String() string {
	switch s {
	case Binary:
		return "Binary"
	case Octal:
		return "Octal"
	case Decimal:
		return "Decimal"
	case Hexadecimal:
		return "Hexadecimal"
	}
	return "unknown"
}

func (b Base) EnumIndex() int {
	return int(b)
}

package alterationtype

type Enum = int

const (
	Maintenance Enum = iota
	Cosmetic
	Performance
	Utility // Canopy, Towbar, etc.
)

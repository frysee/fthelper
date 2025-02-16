package fs

type Mode string

const (
	SINGLE   Mode = "single"
	MULTIPLE Mode = "multiple"
)

func ToMode(s string) (Mode, bool) {
	switch s {
	case "single", "s":
		return SINGLE, true
	case "multiple", "mul", "m":
		return MULTIPLE, true
	}

	return SINGLE, false
}

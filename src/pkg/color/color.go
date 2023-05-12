package color

type colorType struct {
	Reset  string
	Red    string
	Green  string
	Yellow string
	Blue   string
	Purple string
	Cyan   string
	Gray   string
	White  string
}

var Color colorType

const (
	charEscape = "\033["
)

func init() {
	Color = colorType{
		Reset:  charEscape + "0m",
		Red:    charEscape + "31m",
		Green:  charEscape + "32m",
		Yellow: charEscape + "33m",
		Blue:   charEscape + "34m",
		Purple: charEscape + "35m",
		Cyan:   charEscape + "36m",
		Gray:   charEscape + "37m",
		White:  charEscape + "97m",
	}
}

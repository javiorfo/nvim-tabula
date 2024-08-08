package border

type Border struct {
	CornerUpLeft      string
	CornerUpRight     string
	CornerBottomLeft  string
	CornerBottomRight string
	DivisionUp        string
	DivisionBottom    string
	Horizontal        string
	Vertical          string
	Intersection      string
	VerticalLeft      string
	VerticalRight     string
}

func GetBorder(opt BorderOption) Border {
	return borders[opt]
}

type BorderOption int

const (
	Default BorderOption = iota + 1
	Simple
	Rounded
	Double
	SimpleDouble
)

var borders map[BorderOption]Border

func init() {
	borders = map[BorderOption]Border{
		1: {
			CornerUpLeft:      "┏",
			CornerUpRight:     "┓",
			CornerBottomLeft:  "┗",
			CornerBottomRight: "┛",
			DivisionUp:        "┳",
			DivisionBottom:    "┻",
			Horizontal:        "━",
			Vertical:          "┃",
			Intersection:      "╋",
			VerticalLeft:      "┣",
			VerticalRight:     "┫",
		},
		2: {
			CornerUpLeft:      "┌",
			CornerUpRight:     "┐",
			CornerBottomLeft:  "└",
			CornerBottomRight: "┘",
			DivisionUp:        "┬",
			DivisionBottom:    "┴",
			Horizontal:        "─",
			Vertical:          "│",
			Intersection:      "┼",
			VerticalLeft:      "├",
			VerticalRight:     "┤",
		},
		3: {
			CornerUpLeft:      "╭",
			CornerUpRight:     "╮",
			CornerBottomLeft:  "╰",
			CornerBottomRight: "╯",
			DivisionUp:        "┬",
			DivisionBottom:    "┴",
			Horizontal:        "─",
			Vertical:          "│",
			Intersection:      "┼",
			VerticalLeft:      "├",
			VerticalRight:     "┤",
		},
		4: {
			CornerUpLeft:      "╔",
			CornerUpRight:     "╗",
			CornerBottomLeft:  "╚",
			CornerBottomRight: "╝",
			DivisionUp:        "╦",
			DivisionBottom:    "╩",
			Horizontal:        "═",
			Vertical:          "║",
			Intersection:      "╬",
			VerticalLeft:      "╠",
			VerticalRight:     "╣",
		},
		5: {
			CornerUpLeft:      "╒",
			CornerUpRight:     "╕",
			CornerBottomLeft:  "╘",
			CornerBottomRight: "╛",
			DivisionUp:        "╤",
			DivisionBottom:    "╧",
			Horizontal:        "═",
			Vertical:          "│",
			Intersection:      "╪",
			VerticalLeft:      "╞",
			VerticalRight:     "╡",
		},
	}
}

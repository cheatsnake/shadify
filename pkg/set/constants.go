package set

const (
	DeckSize        int    = 81
	StartLayoutSize int    = 12
	SetSize         int    = 3
	MaxLayerSize    int    = 21
	idSeparator     string = "-"
	stateSeparator  string = "@"
)

var (
	numbers  = []int{1, 2, 3}
	shapes   = []string{"diamond", "squiggle", "oval"}
	shadings = []string{"solid", "striped", "open"}
	colors   = []string{"red", "green", "purple"}
	Deck     = generateCards()
)
package set

const (
	deckSize        int    = 81
	startLayoutSize int    = 12
	setSize         int    = 3
	maxLayerSize    int    = 21
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
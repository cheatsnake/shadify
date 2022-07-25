package set

const (
	// Total number of cards in game
	DeckSize int = 81
	// Default size of game layout
	StartLayoutSize int = 12
	// Number of cards in set
	SetSize int = 3
	// Max possible game layout size (in this size is guaranteed to be a set)
	MaxLayerSize   int    = 21
	idSeparator    string = "-"
	stateSeparator string = "@"
)

//error constants
const (
	setIsNotValid        string = "this combination is not a set"
	peakMaxLayoutSize    string = "the operation of adding cards has been canceled (in the layout with size of more then 20 cards is guaranteed to be a set)"
	freeCardsOver        string = "there are no more free cards left"
	gameOver             string = "the game is over there are no more sets"
	setsNotFound         string = "sets not found"
	notEnoughCards       string = "the size of the layout and the number of won cards in total must have at least 12 cards"
	layoutIsTooBig       string = "layout is too big (the maximum size of the layout is 21 cards)"
	incorrectCardsAmount string = "the layout size and the number of won cards must be a multiple of 3"
	tooMuchCards         string = "cannot be more than 81 cards in the game"
)

var (
	numbers  = []int{1, 2, 3}
	shapes   = []string{"diamond", "squiggle", "oval"}
	shadings = []string{"solid", "striped", "open"}
	colors   = []string{"red", "green", "purple"}
	Deck     = generateCards()
)

package countries

// app constant
const (
	flagImageUrl     string = "https://flagcdn.com/w320/"
	flagImageFormat  string = ".png"
	maxVariantsCount int    = 6
	minVariantsCount int    = 2
	maxQuizAmount    int    = 20
)

// error constant
const (
	variantsCountToBig   string = "the maximum number of variants should not exceed 6"
	variantsCountToSmall string = "the number of variants should not be less than 2"
	quizAmountToBig      string = "the maximum value of amount should not exceed 20"
	quizAmountToSmall    string = "the value of amount should be positive"
)

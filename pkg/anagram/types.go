package anagram

// Core struct of anagram module
type Core struct {
	Task  string   `json:"task"`  // Task is random word from which you need to make other words
	Words []string `json:"words"` // Array of all possible words that can be made up from the Task
}

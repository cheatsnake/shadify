package anagram

// Core struct of anagram module
type Core struct {
	Task  string   // Task is random word from which you need to make other words
	Words []string // Array of all possible words that can be made up from the Task
}
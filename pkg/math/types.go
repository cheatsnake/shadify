package math

type Core struct{}

type BasicExpression struct {
	First      int    `json:"first"`
	Second     int    `json:"second"`
	Operation  string `json:"operation"`
	Expression string `json:"expression"`
	Answer     int    `json:"answer"`
}

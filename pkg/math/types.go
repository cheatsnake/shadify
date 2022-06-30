package math

type Core struct{}

type BasicExpression struct {
	First      int    `json:"first"`
	Second     int    `json:"second"`
	Operation  string `json:"operation"`
	Expression string `json:"expression"`
	Answer     int    `json:"answer"`
}

type QuadraticEquation struct {
	Equation     string `json:"equation"`
	A            int    `json:"a"`
	B            int    `json:"b"`
	C            int    `json:"c"`
	Discriminant int    `json:"discriminant"`
	X1           string `json:"x1"`
	X2           string `json:"x2"`
}

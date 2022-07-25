package math

// Basic structure of the expressions with addition, subtraction, multiplication, and division
type BasicExpression struct {
	First      int    `json:"first"`      // First number of the expression
	Second     int    `json:"second"`     // Second number of the expression
	Operation  string `json:"operation"`  // Symbol of the operation
	Expression string `json:"expression"` // Full expression as a string
	Answer     int    `json:"answer"`     // Result of the expression
}

// Structure for quadratic equation
type QuadraticEquation struct {
	Equation     string `json:"equation"`     // Full quadratic equations as a string
	A            int    `json:"a"`            // A coefficient of the quadratic equations
	B            int    `json:"b"`            // B coefficient of the quadratic equations
	C            int    `json:"c"`            // C coefficient of the quadratic equations
	Discriminant int    `json:"discriminant"` // Discriminant of the quadratic equations
	X1           string `json:"x1"`           // First root of the quadratic equation
	X2           string `json:"x2"`           // Second root of the quadratic equation
}

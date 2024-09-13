package tool

/* Conditional or Ternary Operator */
func TernaryOperator(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
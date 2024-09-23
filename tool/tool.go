package tool

/* Conditional or Ternary Operator */
func TernaryOperator(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

/* Opening a file */
func OpenFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err = f.Close(); err != nil {
				fmt.Printf("defer close file %v %v\n", fileName, err)
			}
		}(f) // By taking advantaage of defer's property of evaluating first and then delaying execution, the file to be closed is specified at the time of registration,
		// and it will be correctly closed when the function ends. This method is useful when mutiple files are opened.
	}

	return nil
}
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

// The Function is for dertermining wheather the given slice contains the targets.
func Contains(slice []interface{}, targets ...interface{}) bool {
	for _, target := range targets {
		found := false
		for _, item := range slice {
			switch item.(type) {
			case int:
				if v, ok := target.(int); ok && v == item {
					found = true
					break
				}
			case int64:
				if v, ok := target.(int64); ok && v == item {
					found = true
					break
				}
			case int32:
				if v, ok := target.(int32); ok && v == item {
					found = true
					break
				}
			case string:
				if v, ok := target.(string); ok && v == item {
					found = true
					break
				}
				// add more cases as needed...
			}
		}
		if !found {
			return false
		}
	}
	return true
}

/* Record the time cost of this function */
func comsumeTime() {
	defer func(startTime int64) {
		log.Printf("time cost: %v ms", (time.Now().UnixNano()-startTime)/1e6)
	}(time.Now().Unix())

	/* business logic below*/
}
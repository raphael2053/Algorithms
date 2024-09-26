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

// 通过反射机制来获取interface的字段和方法
// func main() {

// 	user := User{1, "Allen.Wu", 25}

// 	DoFiledAndMethod(user)

// }
type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}
func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}
func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	// 2. 再分别通过reflect.Type的Method获取对应的真实的方法（函数）
	// 3. 最后对结果取其Name和Type得知具体的方法名
	// 4. 也就是说反射可以将“反射类型对象”再重新转换为“接口类型变量”
	// 5. struct 或者 struct 的嵌套都是一样的判断处理方式
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
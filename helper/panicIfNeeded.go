package helper

// PanicIfNeeded function    
func PanicIfNeeded(err error) {
	if err != nil {
		panic(err)
	}
}

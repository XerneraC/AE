package main


// This function is to replace the ternaries.
// It is not currently used alot (if even) for performance reasons
// (I have a hunch, that interface{} is performance intensive)
func ternary(condition bool, ifTrue interface{}, ifFalse interface{}) interface{} {
	if condition {
		return ifTrue
	}
	return ifFalse
}
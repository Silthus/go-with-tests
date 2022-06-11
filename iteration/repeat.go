package iteration

func Repeat(s string, count int) (r string) {
	if count < 0 {
		panic("Repeat called with negative count.")
	}
	for i := 0; i < count; i++ {
		r += s
	}
	return
}

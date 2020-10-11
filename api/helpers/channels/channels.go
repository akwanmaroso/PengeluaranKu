package channels

func OK(done <-chan bool) bool {
	if ok := <-done; !ok {
		return false
	}
	return true
}

package util

// Must checks if the passed err is nil. If not - it logs the error and exits the program.
func Must(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}

package worker

type Job struct {
	Data struct {
		ID         int
		Started    bool
		Progress   float64
		InputFile  string
		OutputFile string
	}
}

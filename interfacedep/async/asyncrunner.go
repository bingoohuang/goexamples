package async

type Worker interface {
	Do()
}

type Runner struct{}

func (r Runner) Run(w interface{ Worker }) {
	go w.Do()
}

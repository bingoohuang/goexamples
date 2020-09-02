package runner

type Worker interface {
	Do()
}

type Runner interface {
	Run(w interface{ Worker })
}

func Run(r Runner, w interface{ Worker }) {
	r.Run(w)
}

package kai

type Middleware func(next HandleFunc) HandleFunc

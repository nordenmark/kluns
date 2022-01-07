package kluns

type config struct {
	version string
}

func Config() config {
	return config{version: "1.0.0"}
}

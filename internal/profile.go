package internal

type Profile struct {
	Type string

	Username string
	Tagline  string
	URL      string
	TwtURL   string
	BlogsURL string

	Followers map[string]string
	Following map[string]string
}

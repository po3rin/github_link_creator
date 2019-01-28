package entity

// User has name field.
type User struct {
	AvatarURL string `json:"avatar_url"`
}

// Repo has repository data.
type Repo struct {
	Name        string `json:"full_name"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
	Forks       int    `json:"forks_count"`
	Stars       int    `json:"stargazers_count"`
	Owner       User   `json:"owner"`
}

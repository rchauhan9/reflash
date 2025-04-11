package study

type StudyProject struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Icon *string `json:"icon"`
}

type StudyProjectCard struct {
	ID             string
	StudyProjectID string `db:"project_id"`
	Question       string
	Answer         string
}

package study

type StudyProject struct {
	ID   string
	Name string
	Icon *string
}

type StudyProjectCard struct {
	ID             string
	StudyProjectID string `db:"project_id"`
	Question       string
	Answer         string
}

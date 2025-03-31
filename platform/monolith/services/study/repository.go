package study

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	db "github.com/rchauhan9/reflash/monolith/common/database"
	"strings"
)

type Repository interface {
	CreateStudyProject(ctx context.Context, name string, icon *string) (string, error)
	CreateCards(ctx context.Context, cards []CreateCard) ([]StudyProjectCard, error)
	DeleteCards(ctx context.Context, studyProjectID string) error
}

func NewRepository(dbPool db.Pool) Repository {
	return &repository{
		pool: dbPool,
	}
}

type repository struct {
	pool db.Pool
}

func (r *repository) CreateStudyProject(ctx context.Context, name string, icon *string) (string, error) {
	query := `
		INSERT INTO project (name, icon)
		VALUES ($1, $2)
		RETURNING id
	`
	row := r.pool.QueryRow(ctx, query, name, icon)
	var studyProjectID string
	err := row.Scan(
		&studyProjectID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to insert study project into database")
	}
	return studyProjectID, nil

}

type CreateCard struct {
	StudyProjectID string `json:"study_project_id"`
	Question       string `json:"question"`
	Answer         string `json:"answer"`
}

func (r *repository) CreateCards(ctx context.Context, cards []CreateCard) ([]StudyProjectCard, error) {
	query := `
		INSERT INTO project_card (project_id, question, answer)
		VALUES %s
		RETURNING id, project_id, question, answer
	`
	numberOfColumns := 3
	values := make([]string, 0, len(cards))
	args := make([]interface{}, 0, len(cards)*numberOfColumns)
	for i, row := range cards {
		values = append(values, fmt.Sprintf("($%d, $%d, $%d)", i*numberOfColumns+1, i*numberOfColumns+2, i*numberOfColumns+3))
		args = append(args, row.StudyProjectID, row.Question, row.Answer)
	}

	sql := fmt.Sprintf(query, strings.Join(values, ", "))
	rows, err := r.pool.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to insert study project cards into database")
	}
	var cardsInserted []StudyProjectCard
	for rows.Next() {
		var card StudyProjectCard
		if err := rows.Scan(&card.ID, &card.StudyProjectID, &card.Question, &card.Answer); err != nil {
			return nil, errors.Wrap(err, "error scanning row from database")
		}
		cardsInserted = append(cardsInserted, card)
	}
	return cardsInserted, nil
}

func (r *repository) DeleteCards(ctx context.Context, studyProjectID string) error {
	query := `
		DELETE FROM project_card
		WHERE project_id = $1
	`
	_, err := r.pool.Exec(ctx, query, studyProjectID)
	if err != nil {
		return errors.Wrapf(err, "failed to delete cards for study project with id %s", studyProjectID)
	}
	return nil
}

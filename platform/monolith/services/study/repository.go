package study

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	db "github.com/rchauhan9/reflash/monolith/common/database"
	"strings"
)

type Repository interface {
	ListStudyProjects(ctx context.Context, userID string) ([]StudyProject, error)
	CreateStudyProject(ctx context.Context, userID string, name string, icon *string) (string, error)

	ListCards(ctx context.Context, userID string, studyProjectID string) ([]StudyProjectCard, error)
	CreateCards(ctx context.Context, userID string, cards []CreateCard) ([]StudyProjectCard, error)
	DeleteCards(ctx context.Context, userID string, studyProjectID string) error
}

func NewRepository(dbPool db.Pool) Repository {
	return &repository{
		pool: dbPool,
	}
}

type repository struct {
	pool db.Pool
}

func (r *repository) ListStudyProjects(ctx context.Context, userID string) ([]StudyProject, error) {
	query := `
		SELECT id, name, icon
		FROM project
		WHERE user_id = $1
	`
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to query study projects")
	}
	studyProjects, err := pgx.CollectRows(rows, pgx.RowToStructByName[StudyProject])
	if err != nil {
		return nil, errors.Wrapf(err, "failed to collect study projects")
	}
	return studyProjects, nil
}

func (r *repository) CreateStudyProject(ctx context.Context, userID string, name string, icon *string) (string, error) {
	query := `
		INSERT INTO project (user_id, name, icon)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	row := r.pool.QueryRow(ctx, query, userID, name, icon)
	var studyProjectID string
	err := row.Scan(
		&studyProjectID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to insert study project into database")
	}
	return studyProjectID, nil

}

func (r *repository) ListCards(ctx context.Context, userID string, studyProjectID string) ([]StudyProjectCard, error) {
	query := `
		SELECT 
		    id, 
		    project_id, 
		    question, 
		    answer
		FROM project_card
		WHERE user_id = $1
			AND project_id = $2
	`
	rows, err := r.pool.Query(ctx, query, userID, studyProjectID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to query study project cards")
	}
	cards, err := pgx.CollectRows(rows, pgx.RowToStructByName[StudyProjectCard])
	if err != nil {
		return nil, errors.Wrapf(err, "failed to collect study project cards")
	}
	return cards, nil
}

type CreateCard struct {
	StudyProjectID string `json:"study_project_id"`
	Question       string `json:"question"`
	Answer         string `json:"answer"`
}

func (r *repository) CreateCards(ctx context.Context, userID string, cards []CreateCard) ([]StudyProjectCard, error) {
	query := `
		INSERT INTO project_card (user_id, project_id, question, answer)
		VALUES %s
		RETURNING id, project_id, question, answer
	`
	numberOfColumns := 4
	values := make([]string, 0, len(cards))
	args := make([]interface{}, 0, len(cards)*numberOfColumns)
	for i, row := range cards {
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*numberOfColumns+1, i*numberOfColumns+2, i*numberOfColumns+3, i*numberOfColumns+4))
		args = append(args, userID, row.StudyProjectID, row.Question, row.Answer)
	}

	sql := fmt.Sprintf(query, strings.Join(values, ", "))
	rows, err := r.pool.Query(ctx, sql, args...)
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

func (r *repository) DeleteCards(ctx context.Context, userID string, studyProjectID string) error {
	query := `
		DELETE FROM project_card
		WHERE user_id = $1
			AND project_id = $2
	`
	_, err := r.pool.Exec(ctx, query, userID, studyProjectID)
	if err != nil {
		return errors.Wrapf(err, "failed to delete cards for study project with id %s", studyProjectID)
	}
	return nil
}

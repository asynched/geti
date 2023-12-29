package repositories

import (
	"database/sql"

	"github.com/asynched/geti/domain/entities"
)

type VisitRepository interface {
	Create(visit entities.Visit) (entities.Visit, error)
	FindAll(linkId int) ([]entities.Visit, error)
}

type visitRepositoryImpl struct {
	db *sql.DB
}

func NewVisitRepository(db *sql.DB) VisitRepository {
	return &visitRepositoryImpl{db}
}

func (r *visitRepositoryImpl) Create(visit entities.Visit) (entities.Visit, error) {
	query := `
		INSERT INTO visits (referrer, user_agent, ip, link_id)
		VALUES (?, ?, ?, ?)
		RETURNING *
	`

	row := r.db.QueryRow(query, visit.Referrer, visit.UserAgent, visit.Ip, visit.LinkId)

	err := row.Scan(
		&visit.Id,
		&visit.Referrer,
		&visit.UserAgent,
		&visit.Ip,
		&visit.LinkId,
		&visit.CreatedAt,
	)

	if err != nil {
		return entities.Visit{}, err
	}

	return visit, nil
}

func (r *visitRepositoryImpl) FindAll(linkId int) ([]entities.Visit, error) {
	query := `
		SELECT * FROM visits
		WHERE link_id = ?
	`

	rows, err := r.db.Query(query, linkId)

	if err != nil {
		return nil, err
	}

	var visits []entities.Visit

	for rows.Next() {
		var visit entities.Visit

		err := rows.Scan(
			&visit.Id,
			&visit.Referrer,
			&visit.UserAgent,
			&visit.Ip,
			&visit.LinkId,
			&visit.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		visits = append(visits, visit)
	}

	return visits, nil
}

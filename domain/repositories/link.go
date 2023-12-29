package repositories

import (
	"database/sql"

	"github.com/asynched/geti/domain/entities"
)

type LinkRepository interface {
	Create(link entities.Link) (entities.Link, error)
	FindBySlug(slug string) (entities.Link, error)
}

type linkRepositoryImpl struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) LinkRepository {
	return &linkRepositoryImpl{db}
}

func (r *linkRepositoryImpl) Create(link entities.Link) (entities.Link, error) {
	query := `
		INSERT INTO links (slug, redirect_to)
		VALUES (?, ?)
		RETURNING *
	`

	row := r.db.QueryRow(query, link.Slug, link.RedirectTo)

	err := row.Scan(
		&link.Id,
		&link.Slug,
		&link.RedirectTo,
		&link.CreatedAt,
		&link.UpdatedAt,
	)

	if err != nil {
		return entities.Link{}, err
	}

	return link, nil
}

func (r *linkRepositoryImpl) FindBySlug(slug string) (entities.Link, error) {
	query := `
		SELECT *
		FROM links
		WHERE slug = ?
	`

	row := r.db.QueryRow(query, slug)

	link := entities.Link{}

	err := row.Scan(
		&link.Id,
		&link.Slug,
		&link.RedirectTo,
		&link.CreatedAt,
		&link.UpdatedAt,
	)

	if err != nil {
		return entities.Link{}, err
	}

	return link, nil
}

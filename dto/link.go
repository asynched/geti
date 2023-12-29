package dto

import (
	"errors"

	"github.com/asynched/geti/domain/entities"
)

type CreateLink struct {
	Slug        string `json:"slug"`
	RedirectUrl string `json:"redirectUrl"`
}

func (c *CreateLink) Validate() error {
	if c.Slug == "" {
		return errors.New("slug is required")
	}

	if c.RedirectUrl == "" {
		return errors.New("redirectUrl is required")
	}

	return nil
}

func (c *CreateLink) ToEntity() entities.Link {
	return entities.Link{
		Slug:       c.Slug,
		RedirectTo: c.RedirectUrl,
	}
}

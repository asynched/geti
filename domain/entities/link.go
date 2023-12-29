package entities

type Link struct {
	Id         int    `json:"id"`
	Slug       string `json:"slug"`
	RedirectTo string `json:"redirectTo"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

package models

import "time"

type (
	BlogResp struct {
		ID    string `query:"id"`
		Title string `query:"title"`
		Body  string `query:"body"`
	}

	Post struct {
		Title       string    `toml:"title"`
		Slug        string    `toml:"slug"`
		Description string    `toml:"description"`
		Date        time.Time `toml:"date"`
		Author      Author    `toml:"author"`
	}

	Author struct {
		Name  string `toml:"name"`
		Email string `toml:"email"`
	}
)

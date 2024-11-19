package routes

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/deej-tsn/blog-go/internal/helper"
	"github.com/deej-tsn/blog-go/internal/models"
	"github.com/deej-tsn/blog-go/web/components"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
)

type (
	// handler has a folder pointer in it
	SlugReader struct {
		pathToFile string
	}
)

func NewSlugReader(pathToFile string) *SlugReader {
	return &SlugReader{
		pathToFile: pathToFile,
	}
}

func (sl *SlugReader) ReadFile(slug string) (string, error) {
	f, err := os.Open(sl.pathToFile + slug + ".md")
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (sl *SlugReader) GetPost(c echo.Context) error {
	post := new(models.Post)
	slug := c.Param("slug")
	postMarkdown, err := sl.ReadFile(slug)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	mdRenderer := goldmark.New(
		goldmark.WithExtensions(
			highlighting.Highlighting,
			extension.GFM,
		),
	)
	rest, err := frontmatter.Parse(strings.NewReader(postMarkdown), &post)
	if err != nil {
		return c.String(http.StatusForbidden, err.Error())
	}
	var buf bytes.Buffer
	err = mdRenderer.Convert(rest, &buf)
	if err != nil {
		// TODO: Handle different errors in the future
		return c.String(201, err.Error())
	}
	return helper.Render(c, http.StatusAccepted, components.Index(slug, components.Post(buf.String())))
}

func (sl *SlugReader) ParseFile(slug string) (*models.Post, error) {
	post := new(models.Post)
	post.Slug = slug
	postMarkdown, err := sl.ReadFile(post.Slug)
	if err != nil {
		return nil, err
	}
	_, err = frontmatter.Parse(strings.NewReader(postMarkdown), &post)
	if err != nil {
		log.Printf("Could not parse content file %s due to following error:\n %s\n", slug, err)
	}
	return post, nil
}

func (sl *SlugReader) GetAllPosts(c echo.Context) error {
	posts := []models.Post{}
	sortBy := c.QueryParam("sort")
	fileNames := helper.FindFilenamesInDirectory(sl.pathToFile, ".md")
	for _, file := range fileNames {
		post, err := sl.ParseFile(file)
		if err != nil {
			log.Printf("Could not parse toml frontmatter to following error:\n %s\n", err)
		} else {
			posts = append(posts, *post)
		}

	}

	helper.SortPosts(sortBy, &posts)

	return helper.Render(c, http.StatusAccepted, components.PostsList(posts))
}

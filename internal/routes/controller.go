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
)

type (
	// handler has a db pointer in it
	SlugReader interface {
		Read(slug string) (string, error)
	}

	FileReader struct{}
)

func NewFileReader() *FileReader {
	return &FileReader{}
}

func (fsr FileReader) Read(slug string) (string, error) {
	f, err := os.Open("./data/posts/" + slug + ".md")
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

func getAllPostNames(dir string) []string {
	fileNames := []string{}
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fileNames = append(fileNames, strings.TrimSuffix(file.Name(), ".md"))
	}
	return fileNames
}

func (sl *FileReader) GetPost(c echo.Context) error {
	post := new(models.Post)
	slug := c.Param("slug")
	postMarkdown, err := sl.Read(slug)
	if err != nil {
		return c.String(http.StatusForbidden, err.Error())
	}
	mdRenderer := goldmark.New(
		goldmark.WithExtensions(
			highlighting.Highlighting,
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
	return helper.Render(c, http.StatusAccepted, components.Post(buf.String()))
}

func (sl *FileReader) PostHandler(c echo.Context) error {
	post := sl.ParseFile("post_1")
	return helper.Render(c, http.StatusAccepted, components.PostHero(*post))
}

func (sl *FileReader) ParseFile(slug string) *models.Post {
	post := new(models.Post)
	post.Slug = slug
	postMarkdown, err := sl.Read(post.Slug)
	if err != nil {
		log.Fatal(err)
	}
	_, err = frontmatter.Parse(strings.NewReader(postMarkdown), &post)
	if err != nil {
		log.Fatal(err)
	}
	return post
}

func (sl *FileReader) GetAllPosts(c echo.Context) error {
	posts := []models.Post{}
	log.Println("enter")
	fileNames := getAllPostNames("./data/posts")
	log.Println(fileNames)

	for _, file := range fileNames {
		posts = append(posts, *sl.ParseFile(file))
	}

	return helper.Render(c, http.StatusAccepted, components.PostsList(posts))
}

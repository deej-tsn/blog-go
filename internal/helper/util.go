package helper

import (
	"context"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/a-h/templ"
	"github.com/deej-tsn/blog-go/internal/models"
	"github.com/labstack/echo/v4"
)

type (
	sortFunction func(a, b models.Post) int
)

func Render(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)
	err := t.Render(context.Background(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}
	return nil
}

func SortPosts(sortBy string, posts *[]models.Post) {
	var sortFunction sortFunction
	if sortBy == "newest" || sortBy == "" {
		sortFunction = func(a, b models.Post) int { return -1 * a.Date.Compare(b.Date) }
	} else {
		sortFunction = func(a, b models.Post) int { return a.Date.Compare(b.Date) }
	}
	slices.SortFunc(*posts, sortFunction)
}

func FindFilenamesInDirectory(dir string, extension string) []string {
	fileNames := []string{}
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("No filenames return due to following error:\n %s\n", err)
		return fileNames
	}
	for _, file := range files {
		fileNames = append(fileNames, strings.TrimSuffix(file.Name(), extension))
	}
	return fileNames
}

package models

import (
	"bytes"
	"fmt"
	"io/fs"
	"sort"
	"time"

	"github.com/vlkvch/bloggee/internal/markdown"
	gm "github.com/yuin/goldmark"
	gmmeta "github.com/yuin/goldmark-meta"
	gmparser "github.com/yuin/goldmark/parser"
	gmutil "github.com/yuin/goldmark/util"
)

type Post struct {
	ID        string
	Title     string
	Author    string
	Content   string
	Published time.Time
}

type PostModel struct {
	Dir fs.FS
}

func (m *PostModel) Get(id string) (*Post, error) {
	postDir, err := fs.Sub(m.Dir, id)
	if err != nil {
		return nil, err
	}

	fileContents, err := fs.ReadFile(postDir, "index.md")
	if err != nil {
		return nil, ErrNoPost
	}

	r := &markdown.ImageLinkRewriter{PostDirName: id}

	md := gm.New(gm.WithExtensions(gmmeta.Meta), gm.WithParserOptions(gmparser.WithASTTransformers(gmutil.Prioritized(r, 100))))
	ctx := gmparser.NewContext()

	content := new(bytes.Buffer)

	if err = md.Convert(fileContents, content, gmparser.WithContext(ctx)); err != nil {
		return nil, err
	}

	metadata := gmmeta.Get(ctx)

	title := fmt.Sprint(metadata["title"])
	author := fmt.Sprint(metadata["author"])
	published, err := time.Parse(time.DateOnly, fmt.Sprint(metadata["published"]))
	if err != nil {
		return nil, err
	}

	post := &Post{
		ID:        id,
		Title:     title,
		Author:    author,
		Content:   string(content.Bytes()),
		Published: published,
	}

	return post, nil
}

func (m *PostModel) All() ([]*Post, error) {
	postDirs, err := fs.ReadDir(m.Dir, ".")
	if err != nil {
		return nil, err
	}

	posts := []*Post{}

	for _, dir := range postDirs {
		p, err := m.Get(dir.Name())
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Published.After(posts[j].Published)
	})

	return posts, nil
}

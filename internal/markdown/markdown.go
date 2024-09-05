package markdown

import (
	"fmt"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type ImageLinkRewriter struct {
	PostDirName string
}

func (r *ImageLinkRewriter) Transform(doc *ast.Document, reader text.Reader, pctx parser.Context) {
	ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		img, ok := node.(*ast.Image)
		if !ok {
			return ast.WalkContinue, nil
		}

		imgName := img.Destination

		img.Destination = []byte(fmt.Sprintf("/posts/%s/%s", r.PostDirName, imgName))

		return ast.WalkSkipChildren, nil
	})
}

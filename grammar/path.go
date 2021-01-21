package grammar

import (
	"github.com/di-wu/parser/ast"
	"github.com/di-wu/parser/op"
	"github.com/scim2/filter-parser/types"
)

func Path(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(ast.Capture{
		Type: typ.Path,
		Value: op.Or{
			op.And{
				ValuePath,
				op.Optional(SubAttr),
			},
			AttrPath,
		},
	})
}

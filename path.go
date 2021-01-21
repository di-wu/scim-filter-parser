package filter

import (
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"github.com/scim2/filter-parser/grammar"
	typ "github.com/scim2/filter-parser/types"
)

// ParsePath parses the given raw data as an Path.
func ParsePath(raw []byte) (Path, error) {
	p, err := ast.New(raw)
	if err != nil {
		return Path{}, err
	}
	node, err := grammar.Path(p)
	if err != nil {
		return Path{}, err
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		return Path{}, err
	}
	return parsePath(node)
}

func parsePath(node *ast.Node) (Path, error) {
	children := node.Children()
	if len(children) == 0 {
		return Path{}, invalidLengthError(typ.Path, 1, 0)
	}

	// AttrPath
	if node.Type == typ.AttrPath {
		attrPath, err := parseAttrPath(node)
		if err != nil {
			return Path{}, err
		}
		return Path{
			AttributePath: attrPath,
		}, nil
	}

	if node.Type != typ.Path {
		return Path{}, invalidTypeError(typ.Path, node.Type)
	}

	// ValuePath SubAttr?
	valuePath, err := parseValuePath(children[0])
	if err != nil {
		return Path{}, err
	}

	var subAttr *string
	if len(children) == 2 {
		node := children[1]
		if node.Type != typ.AttrName {
			return Path{}, invalidTypeError(typ.AttrName, node.Type)
		}
		value := node.ValueString()
		subAttr = &value
	}

	return Path{
		AttributePath:   valuePath.AttributePath,
		ValueExpression: valuePath.ValueFilter,
		SubAttribute:    subAttr,
	}, nil
}

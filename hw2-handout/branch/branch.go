package branch

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func branchCount(fn *ast.FuncDecl) uint {
    x := 0
    for _, v := range fn.Body.List {
       ast.Inspect(v, func(n ast.Node) bool {
            switch n.(type) {
            case *ast.IfStmt:
                x = x+1
            case *ast.ForStmt:
                x = x+1
            case *ast.RangeStmt:
                x = x+1
            case *ast.SwitchStmt:
                x = x+1
            case *ast.TypeSwitchStmt:
                x = x+1
            }
       		return true
       	})
    }
	return uint(x)
}

// ComputeBranchFactors returns a map from the name of the function in the given
// Go code to the number of branching statements it contains.
func ComputeBranchFactors(src string) map[string]uint {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	m := make(map[string]uint)
	for _, decl := range f.Decls {
		switch fn := decl.(type) {
		case *ast.FuncDecl:
			m[fn.Name.Name] = branchCount(fn)
		}
	}

	return m
}

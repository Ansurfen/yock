package walk

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type UnimplementedGoWalk struct{}

func (UnimplementedGoWalk) VisitDir() {}

func (UnimplementedGoWalk) VisitExpr() {}

func (UnimplementedGoWalk) Visit(dir string) {
	pkgs, err := parser.ParseDir(token.NewFileSet(), dir, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				switch v := decl.(type) {
				case *ast.FuncDecl:
				case *ast.BadDecl:
				case *ast.GenDecl:
					for _, spec := range v.Specs {
						switch sv := spec.(type) {
						case *ast.ImportSpec:
							fmt.Println(sv.Path.Value)
						case *ast.TypeSpec:
							// ast.Print(&token.FileSet{}, sv)
							fmt.Println(sv)
						case *ast.ValueSpec:
							ast.Print(&token.FileSet{}, sv)
						default:
							fmt.Println(sv)
						}
					}
				default:
					// fmt.Println(reflect.TypeOf(v))
				}
			}
		}
	}
}

type GoWalk struct{}

const (
	PackageHandle = iota
	DeclFunc
	DeclBad
	DeclGen

	TypeArray
	TypeStruct
	TypeFunc
	TypeInterface
	TypeMap
	TypeChan

	ExprStar
	ExprSelector

	SymbolEllipsis
	SymbolIdent

	HandleDefault
)

type (
	GoDecl = ast.Decl
	GoExpr = ast.Expr

	FuncDecl = *ast.FuncDecl
	BadDecl  = *ast.BadDecl
	GenDecl  = *ast.GenDecl

	ArrayType     = *ast.ArrayType
	StructType    = *ast.StructType
	FuncType      = *ast.FuncType
	InterfaceType = *ast.InterfaceType
	MapType       = *ast.MapType
	ChanType      = *ast.ChanType

	StarExpr     = *ast.StarExpr
	SelectorExpr = *ast.SelectorExpr

	Ellipsis    = *ast.Ellipsis
	IdentSymbol = *ast.Ident
)

type (
	VisitDirHandle  map[uint8]func(pkg string, decl ast.Decl) bool
	VisitExprHandle map[uint8]func(idx int, expr ast.Expr)
)

func (walk *GoWalk) VisitDir(dir string, handle VisitDirHandle) {
	pkgs, err := parser.ParseDir(token.NewFileSet(), dir, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	var pkgHanle func(pkg string) bool
	if method, ok := handle[PackageHandle]; ok {
		pkgHanle = func(pkg string) bool {
			return method(pkg, nil)
		}
	} else {
		pkgHanle = func(pkg string) bool {
			return true
		}
	}
	for pkgName, pkg := range pkgs {
		if !pkgHanle(pkgName) {
			continue
		}
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				switch v := decl.(type) {
				case FuncDecl:
					if fn, ok := handle[DeclFunc]; ok {
						if !fn(pkgName, v) {
							continue
						}
					}
				case BadDecl:
					if fn, ok := handle[DeclBad]; ok {
						if !fn(pkgName, v) {
							continue
						}
					}
				case GenDecl:
					if fn, ok := handle[DeclGen]; ok {
						if !fn(pkgName, v) {
							continue
						}
					}
				default:
					if fn, ok := handle[HandleDefault]; ok {
						if !fn(pkgName, v) {
							continue
						}
					}
				}
			}
		}
	}
}

func (walk *GoWalk) VisitExprs(exprs []ast.Expr, handle VisitExprHandle) {
	for idx, expr := range exprs {
		walk.VisitExpr(idx, expr, handle)
	}
}

func (walk *GoWalk) VisitExpr(idx int, expr ast.Expr, handle VisitExprHandle) {
	switch v := expr.(type) {
	case ArrayType:
		if fn, ok := handle[TypeArray]; ok {
			fn(idx, v)
		}
	case StructType:
		if fn, ok := handle[TypeStruct]; ok {
			fn(idx, v)
		}
	case FuncType:
		if fn, ok := handle[TypeFunc]; ok {
			fn(idx, v)
		}
	case InterfaceType:
		if fn, ok := handle[TypeInterface]; ok {
			fn(idx, v)
		}
	case MapType:
		if fn, ok := handle[TypeMap]; ok {
			fn(idx, v)
		}
	case ChanType:
		if fn, ok := handle[TypeChan]; ok {
			fn(idx, v)
		}
	case IdentSymbol:
		if fn, ok := handle[SymbolIdent]; ok {
			fn(idx, v)
		}
	case Ellipsis:
		if fn, ok := handle[SymbolEllipsis]; ok {
			fn(idx, v)
		}
	case StarExpr:
		if fn, ok := handle[ExprStar]; ok {
			fn(idx, v)
		}
	case SelectorExpr:
		if fn, ok := handle[ExprSelector]; ok {
			fn(idx, v)
		}
	default:
		if fn, ok := handle[HandleDefault]; ok {
			fn(idx, v)
		}
	}
}

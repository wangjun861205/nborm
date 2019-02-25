package nborm

import (
	"errors"
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
)

var commentAttributeRe = regexp.MustCompile(`(?m)([\w_]+?):(.*)$`)
var Pkg string

var attrHandlers = map[string]func(*modelAttrs, []string){
	"DB":  func(a *modelAttrs, group []string) { a.dbName = strings.Trim(group[2], " ") },
	"Tab": func(a *modelAttrs, group []string) { a.tabName = strings.Trim(group[2], " ") },
	"PrimaryKey": func(a *modelAttrs, group []string) {
		l := strings.Split(group[2], ",")
		for i, pk := range l {
			l[i] = strings.Trim(pk, " ")
		}
		a.primaryKey = l
	},
	"Index": func(a *modelAttrs, group []string) {
		l := strings.Split(group[2], ",")
		for i, idx := range l {
			l[i] = strings.Trim(idx, " ")
		}
		a.indices = append(a.indices, l)
	},
	"UniqueKey": func(a *modelAttrs, group []string) {
		l := strings.Split(group[2], ",")
		for i, idx := range l {
			l[i] = strings.Trim(idx, " ")
		}
		a.uniqueKeys = append(a.uniqueKeys, l)
	},
	"Charset": func(a *modelAttrs, group []string) { a.charset = strings.Trim(group[2], " ") },
	"Collate": func(a *modelAttrs, group []string) { a.collate = strings.Trim(group[2], " ") },
}

type modelAttrs struct {
	dbName     string
	tabName    string
	primaryKey []string
	indices    [][]string
	uniqueKeys [][]string
	charset    string
	collate    string
}

func newModelAttrs(comment string) *modelAttrs {
	attrs := commentAttributeRe.FindAllStringSubmatch(comment, -1)
	modelAttrs := &modelAttrs{charset: "utf8mb4", collate: "utf8mb4_bin"}
	for _, attr := range attrs {
		attrHandlers[attr[1]](modelAttrs, attr)
	}
	return modelAttrs
}

type comment map[string]*modelAttrs

func ParseComment(dir string, filenames ...string) (comment, error) {
	if len(filenames) == 0 {
		return nil, errors.New("nborm.ParseComment() error: no file supported")
	}
	tempDir, err := NewTempDir("./", "temp")
	if err != nil {
		return nil, err
	}
	defer tempDir.Remove()
	files, err := tempDir.CopyFiles(filenames...)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if err := file.Close(); err != nil {
			return nil, err
		}
	}
	fs := token.NewFileSet()
	pkgs, err := parser.ParseDir(fs, string(*tempDir), nil, parser.ParseComments)
	if err != nil {
		return nil, nil
	}
	if len(pkgs) != 1 {
		l := make([]string, 0, len(pkgs))
		for _, pkg := range pkgs {
			l = append(l, pkg.Name)
		}
		return nil, fmt.Errorf("nborm.ParseComment() error: multiple packages %s", strings.Join(l, ", "))
	}
	cmt := make(comment)
	for _, pkg := range pkgs {
		if Pkg == "" {
			Pkg = pkg.Name
		}
		docPkg := doc.New(pkg, string(*tempDir), 0)
		for _, typ := range docPkg.Types {
			cmt[typ.Name] = newModelAttrs(typ.Doc)
		}

	}
	return cmt, nil
}

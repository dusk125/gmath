package main

import (
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var (
	outDir string
	skip   = []string{
		"Float32bits", "Float32frombits", "Float64bits", "Float64frombits",
		"Nextafter", "Nextafter32",
		"NaN",
	}
)

func main() {
	root := cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			generate(cmd.Context())

			format(cmd.Context())
			return
		},
	}

	root.Flags().StringVar(&outDir, "dir", ".", "directory to output to")
	root.Flags().StringSliceVar(&skip, "skip", skip, "functions to skip")

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}

func generate(parent context.Context) {
	cmd := exec.CommandContext(parent, "go", "doc", "-short", "math")
	cmd.Stderr = os.Stderr
	b, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	dummyPkg := strings.Builder{}

	fmt.Fprintln(&dummyPkg, "package dummy")

	for l := range strings.SplitSeq(string(b), "\n") {
		if !strings.HasPrefix(l, "func") {
			continue
		}

		fmt.Fprintln(&dummyPkg, l)
	}

	f, err := os.Create(path.Join(outDir, "math.go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprint(f, `// GENERATED! DO NOT EDIT
package gmath
	
import (
	"math"
	
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}
	
`)

	ex, err := parser.ParseFile(token.NewFileSet(), "", dummyPkg.String(), 0)
	if err != nil {
		panic(err)
	}
	for _, d := range ex.Decls {
		var (
			v    = d.(*ast.FuncDecl)
			name = v.Name.Name
		)

		if slices.Contains(skip, name) {
			continue
		}

		var (
			doc            = funcDoc(parent, name)
			params, pCons  = argsFromFieldList(v.Type.Params)
			returns, rCons = argsFromFieldList(v.Type.Results)
		)

		for _, s := range doc {
			fmt.Fprintf(f, "// %v\n", strings.TrimSpace(s))
		}

		fmt.Fprintf(f, "func %v", name)

		if pCons.Has() || rCons.Has() {
			fmt.Fprint(f, "[")
			mult := false
			if pCons.Float || rCons.Float {
				mult = true
				fmt.Fprint(f, "N Number")
			}
			if pCons.Int || rCons.Int {
				if mult {
					fmt.Fprint(f, ",")
				}
				fmt.Fprint(f, "I constraints.Integer")
			}
			fmt.Fprint(f, "]")
		}
		fmt.Fprint(f, "(")
		for i, p := range params {
			if i > 0 {
				fmt.Fprint(f, ",")
			}
			fmt.Fprint(f, p.String())
		}
		fmt.Fprint(f, ")")

		switch len(returns) {
		case 0:
		case 1:
			r := returns[0]
			hasName := r.Name != ""
			if hasName {
				fmt.Fprint(f, "(")
			}
			fmt.Fprint(f, r.String())
			if hasName {
				fmt.Fprint(f, ")")
			}
		default:
			fmt.Fprint(f, "(")
			for i, r := range returns {
				if i > 0 {
					fmt.Fprint(f, ",")
				}
				fmt.Fprint(f, r.String())
			}
			fmt.Fprint(f, ")")
		}

		paramS := strings.Builder{}
		for i, p := range params {
			if i > 0 {
				fmt.Fprint(&paramS, ",")
			}
			switch {
			case p.Constraint.Float:
				fmt.Fprintf(&paramS, "%v(%v)", p.Type, p.Name)
			case p.Constraint.Int:
				fmt.Fprintf(&paramS, "%v(%v)", p.Type, p.Name)
			default:
				fmt.Fprint(&paramS, p.Name)
			}
		}

		returnS := strings.Builder{}
		fmt.Fprint(f, "{\n")
		if len(returns) == 0 {
			fmt.Fprintf(f, "\tmath.%v(%v)", name, paramS.String())
		} else {
			if len(returns) == 1 {
				r := returns[0]
				switch {
				case r.Constraint.Float:
					fmt.Fprint(&returnS, "N(")
				case r.Constraint.Int:
					fmt.Fprint(&returnS, "I(")
				}
				fmt.Fprintf(&returnS, "math.%v(%v)", name, paramS.String())
				if r.Constraint.Has() {
					fmt.Fprint(&returnS, ")")
				}
			} else {
				for i, r := range returns {
					if i > 0 {
						fmt.Fprint(f, ",")
						fmt.Fprint(&returnS, ",")
					}
					fmt.Fprint(f, r.Name+"R")
					switch {
					case r.Constraint.Float:
						fmt.Fprintf(&returnS, "N(%vR)", r.Name)
					case r.Constraint.Int:
						fmt.Fprintf(&returnS, "I(%vR)", r.Name)
					default:
						fmt.Fprint(&returnS, r.Name+"R")
					}
				}
				fmt.Fprintf(f, "\t := math.%v(%v)\n", name, paramS.String())
			}
			fmt.Fprintf(f, "\treturn %v", returnS.String())
		}

		fmt.Fprint(f, "\n}")
		fmt.Fprintln(f)
	}
}

type Arg struct {
	Name       string
	Type       string
	Constraint Constraint
}

func (a Arg) String() string {
	s := strings.Builder{}
	s.WriteString(a.Name)
	s.WriteRune(' ')
	if a.Constraint.Float {
		s.WriteString("N")
	} else if a.Constraint.Int {
		s.WriteString("I")
	} else {
		s.WriteString(a.Type)
	}
	return s.String()
}

type Constraint struct {
	Float bool
	Int   bool
}

func (c Constraint) Has() bool {
	return c.Float || c.Int
}

func argsFromFieldList(fields *ast.FieldList) (args []Arg, constraints Constraint) {
	if l := len(fields.List); l > 0 {
		args = make([]Arg, 0, l)

		for _, p := range fields.List {
			tipe := p.Type.(*ast.Ident).Name
			constr := Constraint{
				Float: tipe == "float64",
				Int:   tipe == "int",
			}
			constraints.Float = constr.Float || constraints.Float
			constraints.Int = constr.Int || constraints.Int

			if len(p.Names) > 0 {
				for _, n := range p.Names {
					args = append(args, Arg{
						Name:       n.Name,
						Type:       tipe,
						Constraint: constr,
					})
				}
			} else {
				args = append(args, Arg{Type: tipe, Constraint: constr})
			}
		}
	}
	return
}

func funcDoc(parent context.Context, name string) []string {
	b := strings.Builder{}
	cmd := exec.CommandContext(parent, "go", "doc", "-short", "math", name)
	cmd.Stdout = &b
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	return strings.Split(b.String(), "\n")[1:]
}

func format(parent context.Context) {
	cmd := exec.CommandContext(parent, "go", "fmt", path.Join(outDir, "math.go"))
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

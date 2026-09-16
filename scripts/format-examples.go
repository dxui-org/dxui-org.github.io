// Command format-examples expands Go examples for reading in documentation.
// It formats complete programs and Go fences without changing API declarations.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var fence = regexp.MustCompile("(?ms)^```go[^\\n]*\\n(.*?)^```[ \\t]*$")

func expandPass(source string) (string, error) {
	// Examples may be full files, declarations, or statements with explanatory comments.
	prefix, suffix := "", ""
	var file *ast.File
	var fs *token.FileSet
	var err error
	for _, wrapper := range [][2]string{{"", ""}, {"package main\n\n", ""}, {"package main\n\nfunc example() {\n", "\n}"}} {
		prefix, suffix = wrapper[0], wrapper[1]
		fs = token.NewFileSet()
		file, err = parser.ParseFile(fs, "example.go", prefix+source+suffix, parser.ParseComments)
		if err == nil {
			break
		}
	}
	if err != nil {
		return "", err
	}
	input := prefix + source + suffix
	insertions := map[int]string{}
	position := func(pos token.Pos) int { return fs.Position(pos).Offset }
	newline := func(offset int) {
		if _, exists := insertions[offset]; exists {
			return
		}
		left, right := offset, offset
		for left > 0 && strings.ContainsRune(" \t\r\n", rune(input[left-1])) {
			left--
		}
		for right < len(input) && strings.ContainsRune(" \t\r\n", rune(input[right])) {
			right++
		}
		if !strings.Contains(input[left:right], "\n") {
			insertions[offset] = "\n"
		}
	}
	list := func(open, close token.Pos, elements []ast.Expr) {
		if len(elements) == 0 {
			return
		}
		newline(position(open) + 1)
		for _, element := range elements {
			newline(position(element.Pos()))
		}
		last := position(elements[len(elements)-1].End())
		end := position(close)
		if !strings.HasPrefix(strings.TrimSpace(input[last:end]), ",") {
			insertions[last] = ",\n"
		}
		newline(end)
	}
	ast.Inspect(file, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.CompositeLit:
			// Keep small coordinate/color literals compact, but give every configured
			// property its own line, including nested one-field props and styles.
			keyed := false
			for _, element := range n.Elts {
				_, ok := element.(*ast.KeyValueExpr)
				keyed = keyed || ok
			}
			if keyed || len(n.Elts) > 1 || position(n.End())-position(n.Pos()) > 88 {
				list(n.Lbrace, n.Rbrace, n.Elts)
			}
		case *ast.CallExpr:
			if len(n.Args) > 1 && !n.Ellipsis.IsValid() && position(n.End())-position(n.Pos()) > 88 {
				list(n.Lparen, n.Rparen, n.Args)
			}
		case *ast.BlockStmt:
			if len(n.List) > 0 {
				newline(position(n.Lbrace) + 1)
				newline(position(n.Rbrace))
				for _, statement := range n.List {
					newline(position(statement.Pos()))
				}
			}
		}
		return true
	})
	positions := make([]int, 0, len(insertions))
	for offset := range insertions {
		positions = append(positions, offset)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(positions)))
	for _, offset := range positions {
		input = input[:offset] + insertions[offset] + input[offset:]
	}
	formatted, err := format.Source([]byte(input))
	if err != nil {
		return "", fmt.Errorf("%w\n%s", err, input)
	}
	result := string(formatted)
	if prefix != "" {
		result = strings.TrimPrefix(result, "package main\n\n")
	}
	if suffix != "" {
		result = strings.TrimPrefix(result, "func example() {\n")
		result = strings.TrimSuffix(result, "}\n")
		lines := strings.Split(result, "\n")
		for i := range lines {
			lines[i] = strings.TrimPrefix(lines[i], "\t")
		}
		result = strings.Join(lines, "\n")
	}
	return strings.TrimSpace(result) + "\n", nil
}

func expand(source string) (string, error) {
	for i := 0; i < 20; i++ {
		next, err := expandPass(source)
		if err != nil {
			return "", err
		}
		if next == source {
			return next, nil
		}
		source = next
	}
	return "", fmt.Errorf("example formatting did not converge")
}

func main() {
	check := flag.Bool("check", false, "report unformatted examples without writing")
	flag.Parse()
	root := "docs"
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}
	changed := 0
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if entry.Name() == ".vitepress" || entry.Name() == "public" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, ".md") {
			return nil
		}
		original, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		var result string
		if strings.HasSuffix(path, ".go") {
			result, err = expand(string(original))
		} else {
			result = fence.ReplaceAllStringFunc(string(original), func(block string) string {
				parts := fence.FindStringSubmatch(block)
				value, formatErr := expand(parts[1])
				if formatErr != nil {
					err = formatErr
					return block
				}
				return "```go\n" + value + "```"
			})
		}
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		if !bytes.Equal(original, []byte(result)) {
			changed++
			if *check {
				fmt.Println(path)
			} else if err := os.WriteFile(path, []byte(result), 0644); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Go example formatting: %d files %s.\n", changed, map[bool]string{true: "need formatting", false: "updated"}[*check])
	if *check && changed != 0 {
		os.Exit(1)
	}
}

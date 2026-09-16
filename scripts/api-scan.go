// Command api-scan extracts the source contract without loading SDL or dependencies.
package main

import (
 "bytes"
 "encoding/json"
 "go/ast"
 "go/format"
 "go/parser"
 "go/token"
 "os"
 "path/filepath"
 "strings"
)

type Entry struct { Package, Name, Kind, File, Doc, Code string; Line int }
func main() {
 root := os.Args[1]
 fs := token.NewFileSet()
 entries := []Entry{}
 printNode := func(n any) string { var b bytes.Buffer; if err := format.Node(&b, fs, n); err != nil { panic(err) }; return b.String() }
 for _, dir := range []string{".", "icon", "icon/catalog", "internal/icondata"} {
  paths, err := filepath.Glob(filepath.Join(root, dir, "*.go")); if err != nil { panic(err) }
  for _, path := range paths {
   if strings.HasSuffix(path, "_test.go") { continue }
   f, err := parser.ParseFile(fs, path, nil, parser.ParseComments); if err != nil { panic(err) }
   for _, decl := range f.Decls {
    add := func(name, kind, doc, code string, pos token.Pos) { rel, _ := filepath.Rel(root, path); entries = append(entries, Entry{dir,name,kind,filepath.ToSlash(rel),doc,code,fs.Position(pos).Line}) }
    switch d := decl.(type) {
    case *ast.FuncDecl:
     if !d.Name.IsExported() { continue }; name := d.Name.Name
     if d.Recv != nil { recv := strings.TrimPrefix(printNode(d.Recv.List[0].Type), "*"); if !token.IsExported(strings.Split(recv,"[")[0]) { continue }; if dir == "internal/icondata" && recv != "Data" { continue }; name = recv+"."+name }
     if dir == "internal/icondata" && d.Recv == nil { continue }
     doc := d.Doc.Text(); d.Doc = nil; d.Body = nil; add(name,"func",doc,printNode(d),d.Pos())
    case *ast.GenDecl:
     if d.Tok == token.IMPORT { continue }
     constantPrinted := false
     for _, spec := range d.Specs {
      switch s := spec.(type) {
      case *ast.TypeSpec:
       if !s.Name.IsExported() && !(dir == "." && strings.Contains(s.Name.Name,"olorToken")) { continue }
       doc := s.Doc.Text(); if doc == "" { doc = d.Doc.Text() }
       if st, ok := s.Type.(*ast.StructType); ok { fields := []*ast.Field{}; for _, field := range st.Fields.List { if len(field.Names)>0 && field.Names[0].IsExported() { fields=append(fields,field) } }; st.Fields.List=fields }
       add(s.Name.Name,"type",doc,"type "+printNode(s),s.Pos())
      case *ast.ValueSpec:
       for _, name := range s.Names { if name.IsExported() {
        doc := s.Doc.Text(); if doc == "" { doc = d.Doc.Text() }
        code := d.Tok.String()+" "+printNode(s)
        if d.Tok == token.CONST { if constantPrinted { code = "" } else { code = printNode(d); constantPrinted = true } } // preserve implicit types and iota once
        if d.Tok == token.VAR && len(s.Values)>0 { if v,ok:=s.Values[0].(*ast.CompositeLit); ok { code="var "+name.Name+" "+printNode(v.Type) } }
        add(name.Name,d.Tok.String(),doc,code,s.Pos())
       } }
      }
     }
    }
   }
  }
 }
 if err:=json.NewEncoder(os.Stdout).Encode(entries); err!=nil { panic(err) }
}

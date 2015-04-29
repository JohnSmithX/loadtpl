package loadtpl

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"strings"
)

var (
	//default template extension
	DefaultTplExt = []string{"html", "tpl"}
	root          string
	t             *template.Template
)

func hasTemplateExt(paths string) bool {
	for _, v := range DefaultTplExt {
		if strings.HasSuffix(paths, "."+v) {
			return true
		}
	}
	return false
}

// add new template extension
func AddTemplateExt(ext string) {
	for _, v := range DefaultTplExt {
		if v == ext {
			return
		}
	}
	DefaultTplExt = append(DefaultTplExt, ext)
}

//load all template from dir directory
func LoadTemplates(dir string) (*template.Template, error) {
	if _, err := os.Stat(dir); err != nil {
		return nil, err
	}
	root = dir
	err := filepath.Walk(dir, pathHandler)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func pathHandler(paths string, f os.FileInfo, err error) error {
	if err != nil {
		log.Println(err)
	}

	if f == nil {
		return err
	}
	if f.IsDir() || (f.Mode()&os.ModeSymlink) > 0 {
		return nil
	}
	if !hasTemplateExt(paths) {
		return nil
	}

	b, err := ioutil.ReadFile(paths)
	if err != nil {
		return err
	}
	s := string(b)

	replace := strings.NewReplacer("\\", "/")
	a := []byte(paths)
	a = a[len([]byte(root)):]
	name := strings.TrimLeft(replace.Replace(string(a)), "/")

	if t == nil {
		t = template.New(name)
	}
	if name == t.Name() {
	} else {
		t = t.New(name)
	}
	_, err = t.Parse(s)
	if err != nil {
		return err
	}

	return nil
}

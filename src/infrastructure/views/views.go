package views

import (
  "html/template"
  "fmt"
  "strings"
  "log"
)

type templateParser struct {
	HTML string
}

type LayoutData struct {
	Page	string
	User		string
	Data		interface{}
}

func (tP *templateParser) Write(p []byte) (n int, err error) {
	tP.HTML += string(p)
	return len(p), nil
}

type staticTemplate struct {
  Title string
}

func Render(page string, data interface{}) (error, []byte){
	tp := &templateParser{}
	t, _ := template.ParseFiles("views/" + page + ".html")
	
	t = t.Funcs(template.FuncMap{"emailExpand": EmailExpander})
	t = t.Funcs(template.FuncMap{"handlebars": Handlebars})
	t = t.Funcs(template.FuncMap(map[string]interface{}{"equal": CheckSame}))
	
	t.Execute(tp, data)
	return nil, []byte(tp.HTML)
}

func RenderLayout(page string, data interface{}) (error, []byte){
	tp := &templateParser{}
	
	//var t = template.Must(template.ParseFiles("views/layout.html",
	//	page + ".html"))
		
	t := template.New("t")
	t = t.Funcs(template.FuncMap{"emailExpand": EmailExpander})
	t = t.Funcs(template.FuncMap{"handlebars": Handlebars})
	t = t.Funcs(template.FuncMap(map[string]interface{}{"equal": CheckSame}))
	
	t, err := t.ParseFiles("views/layout.html",
		page + ".html")
		
	if err != nil {
		log.Printf("layout error %s ", err.Error())
		return nil, []byte(err.Error())
	} else {
		err = t.ExecuteTemplate(tp, "layout.html", data)
		//err = t.Execute(tp, data)
		if err != nil {
			return nil, []byte(err.Error())
		}
	}
	//err := t.Execute(tp, data)
	//if (err != nil){
	//	log.Printf("layout error %s ", err.Error())
	//}
	return nil, []byte(tp.HTML)
}

func Static(page string) (error, []byte) {
  data := &staticTemplate{"Title"}
  return Render(page, data)
}

func StaticLayout(page string) (error, []byte) {
  data := &staticTemplate{"Title"}
  return RenderLayout(page, data)
}

func Error(err error) []byte {
  e, v := Render("error", err)
  if e != nil {
    return []byte(err.Error());
  } 
  return v
}

func Handlebars(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	s = "{{"  + s + "}}"
	return s
}

func CheckSame(args ...interface{}) bool {
	ok := false
	var s1 string
	var s2 string
	if len(args) != 2 {
		return false
	}
	s1, ok = args[0].(string)
	if !ok {
		return false
	}
	s2, ok = args[1].(string)
	if !ok {
		return false
	}
	return s1 == s2
}

func EmailExpander(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	// replace the @ by " at "
	return (substrs[0] + " at " + substrs[1])
}



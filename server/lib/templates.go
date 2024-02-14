package lib

import (
	"html"
	"os"
	"text/template"
)

func HelloHTMLPage() string {
	// Generic hello world HTML page.
	var x = `
	<html>
	  <body>
	    Hello {{.}}
	  </body>
	</html>
`

	return x
}

func CallHelloHTMLPage() {
	// Call the HTML string, inject some code, make sure we are properly encoding.
	t, err := template.New("hello").Parse(HelloHTMLPage())
	if err != nil {
		panic(err)
	}
	unescapedString := "<script>alert('world');</script>"
	encodedString := html.EscapeString(unescapedString)
	t.Execute(os.Stdout, encodedString)
}

// HAS A RENDER FORM FUNCTION WTHAT TAKES (w reppsonse writer, d data and loads the form template and executes it
package views

import (
	"bytes"
	"html/template"
	"net/http"
)
const formTemplate = `
<h1> Create Book Form </h1>
<form action="" method="POST"> 
<div> 
<label for="title">Title </label>
<input type="text" name="title" id="title">
</div>
<div> 
<label for="author">Author </label>
<input type="text" name="author" id="author">
</div>
<div> 
<label for="title">Year </label>
<input type="text" name="year" id="year">
</div>
<div> 
<label for="type">Type </label>
<label><input type="radio" name="type" id="raw">raw</label>
<label><input type="radio" name="type" id="html">html</label>
</div>
<div>
<button>Send</button>
</div>
</form>
`
func RenderForm(w http.ResponseWriter, data interface{})error{
	t , err := template.New("").Parse(formTemplate)
	if err!=nil {
		return err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf,data)
	if err!=nil{
		return err
	}
	_ , err= buf.WriteTo(w)
	if err!=nil{
		return err
	}
	return nil
}
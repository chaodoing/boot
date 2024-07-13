package o

import (
	`bytes`
	`encoding/json`
	`text/template`
	
	`github.com/chaodoing/boot/assets/vscode`
	`github.com/kataras/iris/v12`
)

func (r *Respond) html(content interface{}) (value string, err error) {
	var data []byte
	tpl, err := template.New("json").Parse(vscode.HTML)
	if err != nil {
		return
	}
	data, err = json.Marshal(content)
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, map[string]string{
		"Title": "JSON",
		"Json":  string(data),
	})
	return buf.String(), err
}

func (r *Respond) Negotiation(ctx iris.Context, value interface{}) error {
	html, err := r.html(value)
	if err != nil {
		return err
	}
	ctx.Negotiation().JSON(value).XML(value).YAML(value).HTML(html).Charset("UTF-8").EncodingGzip()
	_, err = ctx.Negotiate(nil)
	if err != nil {
		return err
	}
	return nil
}

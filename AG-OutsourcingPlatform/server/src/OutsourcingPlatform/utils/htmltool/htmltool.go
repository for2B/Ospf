package htmltool

import (
	"fmt"
	"net/http"
)

func ResponHtml(w http.ResponseWriter, title string, body string) {
	fmt.Fprintln(w, `<title>`+title+`</title><div style="
 display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    line-height: 60px;
    text-align: center;
"><h1 style="text-align:center;margin-top:50px;">`+body+`</h1></div`)
}

func ResponDialog(w http.ResponseWriter, title string, body string) {
	fmt.Fprintln(w, `
		<title>`+title+`</title>
		<h1 style="text-align:center;margin-top:50px;">`+body+
		`</h1>`)
}

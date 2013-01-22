package main

import (
    "flag"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os/exec"
    "bytes"
    "strings"
)

var addr = flag.String("0.0.0.0", ":8080", "") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
    
    switchState("off", "B")  

    flag.Parse()
    http.Handle("/", http.HandlerFunc(QR))
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func switchState(state string, name string) {
  cmd := exec.Command("tdtool", "--off", name)
  cmd.Stdin = strings.NewReader("")
  var out bytes.Buffer
  
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Print("got: %q\n", out.String())
  
}

func QR(w http.ResponseWriter, req *http.Request) {
    templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>Go Home!</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`

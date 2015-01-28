package importer

import (
 "github.com/codegangsta/negroni"
 "github.com/codegangsta/cli"
  "html/template"
  "net/http"
  "github.com/parnurzeal/gorequest"
  "fmt"
  "encoding/json"
)

const (
  tmplPocket = `<form method='GET' action='https://getpocket.com/auth/authorize'>
      <input type="hidden" name="request_token" value="{{ .Token }}" />
      <input type="hidden" name="redirect_uri" value="http://local.dev:23000/auth" />
      <button type='submit'>Login to pocket</button>
    </form>`
  tmplAuth = `Authentication. Indexing now`
)

var (
  TokenCode string
  ConsumerKey string
)

func Run(c *cli.Context) {
  ConsumerKey := "37202-e57b812865ccd79bfd739dc4"

  println("Will start server with API %s", c.Args()[1])
  mux := http.NewServeMux()

  mux.HandleFunc("/1", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, TokenCode)
    TokenCode = "100"
  })
  mux.HandleFunc("/2", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, TokenCode)
    TokenCode = "200"
  })



  mux.HandleFunc("/auth", func(w http.ResponseWriter, req *http.Request) {
    url := "https://getpocket.com/v3/oauth/authorize"
    fmt.Fprintf(w, tmplAuth + TokenCode)
    
    request := gorequest.New()
    resp, body, errs := request.Post(url).
      /*Send(payload).*/
      Set("X-Accept", "application/json").
      Send(`{"consumer_key":"` + ConsumerKey + `","code":"` + TokenCode + `"}`).
      End()

    fmt.Print(resp)
    fmt.Print(body)

    if errs != nil {
      fmt.Print(errs)
    }

    var pocketResponse struct {
      AccessToken string `json:"access_token"`
      UserName string `json:"username"`
    }

    if err := json.Unmarshal([]byte(body), &pocketResponse); err == nil {
      fmt.Println("Access token ", pocketResponse.AccessToken)
      fmt.Println(pocketResponse.UserName)
      go doFetch(ConsumerKey, pocketResponse.AccessToken)
    }

  })


  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    url := "https://getpocket.com/v3/oauth/request"

    /*payload := struct {*/
      /*Consumer_Key string `json:"scopes"`*/
      /*Redirect_Uri string `json:"note"`*/
    /*}{*/
      /*Consumer_Key: "37202-e57b812865ccd79bfd739dc4",*/
      /*Redirect_Uri: "http://local.dev:23000/auth",*/
    /*}*/

    var code struct {
      Code string `json:"code"`
    }

    request := gorequest.New()
    resp, body, errs := request.Post(url).
      /*Send(payload).*/
      Set("X-Accept", "application/json").
      Send(`{"consumer_key":"37202-e57b812865ccd79bfd739dc4","redirect_uri":"http://local.dev:23000/auth"}`).
      End()

    fmt.Print(resp)

    if errs != nil {
      fmt.Print(errs)
    }

    if err := json.Unmarshal([]byte(body), &code); err == nil {
      fmt.Println(code)
    }

    TokenCode = code.Code

    /*fmt.Fprintf(w, tmplPocket)*/
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    tmpl, _ := template.New("pocket").Parse(tmplPocket)
    tmpl.Execute(w, struct{ Token string }{ Token: code.Code, })
  })

  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":23000")
}

func doFetch(consumerkey, token string) {
  url := "https://getpocket.com/v3/get"

  request := gorequest.New()
  resp, body, errs := request.Post(url).
    Set("X-Accept", "application/json").
    Send(`{"consumer_key":"` + consumerkey + `","access_token":"` + token  + `"}`).
    End()

  fmt.Println(resp)
  fmt.Println(body)
  fmt.Println(errs)

}


package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

	// This is the main function
	log.Println("listening on :8083")
	var site Site
	site.Init()
	log.Fatal(http.ListenAndServe(":8083", &site))
}

type User struct {
	Name, Email, Pass string
}

type Site struct {
	views    *template.Template
	users    map[string]*User
	cookies  map[string]*User
	handlers map[string]Handler
	loginURI string
}

type Handler func(http.ResponseWriter, *http.Request) error

func (this *Site) Init() {
	this.views = template.Must(template.New("views").Parse(views))
	this.users = map[string]*User{
		"john": &User{"john", "allan@sc.", "12345"},
	}

	//This is to handle the differnt endpoints to be invoked of the client
	
	this.handlers = map[string]Handler{
		"/":          this.home,
		"/home":      this.home,
		"/about":     this.about,
		"/login":     this.login,
		"/protected": this.protected,
	}
	this.cookies = make(map[string]*User)
	this.loginURI = "/login?from="
}

func (this *Site) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, hasHandler := this.handlers[r.URL.Path]; hasHandler {
		if err := handler(w, r); err != nil {
			log.Println("Site.ServeHTTP:", err.Error())
		}
	} else {
		http.NotFound(w, r)
	}
}

func (this *Site) home(w http.ResponseWriter, r *http.Request) error {
	vm := &HomeVM{
		Header:  "Home",
		Message: "Just an ordinary home page",
		Items:   r.URL.Query(),
	}
	return this.views.ExecuteTemplate(w, "home", vm)
}

func (this *Site) about(w http.ResponseWriter, r *http.Request) error {
	vm := &AboutVM{
		Message: "Just an ordinary about page",
	}
	return this.views.ExecuteTemplate(w, "about", vm)
}

func (this *Site) login(w http.ResponseWriter, r *http.Request) (err error) {
	username, password := "", ""
	from := r.URL.Query().Get("from")
	log.Println("Login from:", from)

	if r.Method != "POST" {
		if _, isValid := this.validateCookie(r); isValid {
			goto redirect
		}
		vm := &LoginVM{Action: this.loginURI + from}
		return this.views.ExecuteTemplate(w, "login", vm)
	}

	if err = r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	username, password = r.Form.Get("username"), r.Form.Get("password")
	log.Printf("%s=%s",username, password)
	if _, isValid := this.validateUserpass(w, username, password); !isValid {
		vm := &LoginVM{
			Username: username,
			Action:   this.loginURI + from,
			Message:  "Invalid user name or password",
		}

		fmt.Println("Printing the User Details...")
		fmt.Println("=>", vm.Message)
		fmt.Println("=>", vm.Action)
		fmt.Println("=>",vm.Username)
		
		return this.views.ExecuteTemplate(w, "login", vm)
	}

redirect:
	if len(from) > 0 {
		from, err = url.QueryUnescape(from)
	}
	if len(from) == 0 || err != nil {
		from = "/"
	}
	http.Redirect(w, r, from, http.StatusFound)
	return
}

func (this *Site) protected(w http.ResponseWriter, r *http.Request) error {
	user, validUser := this.authenticate(w, r)
	if !validUser {
		return nil
	}
	vm := &HomeVM{
		Header:  "Protected",
		Message: "Hey " + user.Email + " how ya doin' ?",
		Items:   r.URL.Query(),
	}
	return this.views.ExecuteTemplate(w, "home", vm)
}

func (this *Site) authenticate(w http.ResponseWriter, r *http.Request) (user *User, isValid bool) {
	if user, isValid = this.validateCookie(r); isValid {
		return
	}
	if name, pass, isBasicAuth := r.BasicAuth(); isBasicAuth {
		if user, isValid = this.validateUserpass(w, name, pass); isValid {
			return
		}
	}

	if len(this.loginURI) > 0 {
		from := r.URL.RequestURI()
		log.Println("Redirect to login:", r.RemoteAddr, "from:", from)
		http.Redirect(w, r, this.loginURI+url.QueryEscape(from), http.StatusFound)
	} else {
		log.Println("Unauthorized:", r.RemoteAddr)
		w.Header().Set("WWW-Authenticate", "Basic realm=\"Site\"")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	return nil, false
}

func (this *Site) validateCookie(r *http.Request) (user *User, isValid bool) {
	if cookie, err := r.Cookie("ZeCookie"); err == nil && len(cookie.Value) > 0 {
		log.Println("authenticating cookie:", cookie.Value)
		user, isValid = this.cookies[cookie.Value]
	}
	return
}

func (this *Site) validateUserpass(w http.ResponseWriter, name, pass string) (user *User, isValid bool) {
	log.Println("authenticating user:", name)
	//TODO use bcrypt on the password
	if user, isValid = this.users[name]; isValid && user.Pass == pass {
		//TODO add random & unique value here
		now := time.Now()
		cookieValue := now.Format(time.RFC3339Nano)
		this.cookies[cookieValue] = user
		var cookie http.Cookie
		cookie.Path = "/"
		cookie.Name = "ZeCookie"
		cookie.Value = cookieValue
		cookieTTL := time.Second * 5
		cookie.Expires = now.Add(cookieTTL)
		cookie.MaxAge = int(cookieTTL.Seconds())
		//cookie.HttpOnly = true
		//cookie.Secure = true
		http.SetCookie(w, &cookie)
		return
	}
	return nil, false
}

// view models
type HomeVM struct {
	Header, Message string
	Items           url.Values
}

type AboutVM struct {
	Message string
}

type LoginVM struct {
	Action   string
	Username string
	Message  string
}

const views = `
{{define "home"}}
{{template "head" .}}
	<title> {{.Header}} page </title>
	<script>
		console.log("home");
	</script>
{{template "neck" .}}
	<h2> {{.Header}} </h2>
	<p> {{. }} </p>
	<ul>
	{{range $key, $val := .Items}}
		<li> {{$key}} : {{$val}} </li>
	{{end}}
	</ul>
{{template "tail" .}}
{{end}}

{{define "about"}}
{{template "head" .}}
	<title> About page </title>
	<script>
		console.log("about");
	</script>
{{template "neck" .}}
	<h2> About page </h2>
	<p> {{.Message}} </p>
{{template "tail" .}}
{{end}}

{{define "login"}}
{{template "head" .}}
	<title> Login page </title>
	<script>
		console.log("login");
	</script>
{{template "neck" .}}
	<h2> Login </h2>
	<div>
		<form action="{{.Action}}" method="POST">
			<div>Username <input name="username" type="text" value="{{.Username}}"> </input></div>
			<div>Password <input name="password" type="password" value=""> </input></div>
			<div>{{.Message}}</div>
			<div><input type="submit" value="Login"></input> </div>
		</form>
	</div>
{{template "tail" .}}
{{end}}

{{define "head"}}
<!DOCTYPE html>
<html>
	<head>
	<!-- meta attributes & script links here -->
{{end}}

{{define "neck"}}
	<!-- style links here -->
	</head>
	<body>
	<!-- html header here -->
	<div>
	<a href="/home?arg=value&so=on">Home</a>
	<a href="/protected">Protected</a>
	<a href="/about">About</a>
	<a href="/login">Login</a>
	</div>
{{end}}

{{define "tail"}}
	<!-- scripts & footer here -->
	</body>
</html>
{{end}}
`
s
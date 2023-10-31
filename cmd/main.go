package main

import (
	"os"

	g "github.com/alarbada/gomponents"
	"github.com/alarbada/gomponents/actions"
	. "github.com/alarbada/gomponents/html"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var router = actions.NewRouter()

func Page(children ...g.Node) g.Node {
	return HTML(
		Head(Raw(`
			<script src="https://unpkg.com/htmx.org@1.9.6"></script>
			<script src="https://cdn.tailwindcss.com"></script>
			<link href="https://cdn.jsdelivr.net/npm/daisyui@3.9.4/dist/full.css" rel="stylesheet" type="text/css" />
			<script src="/public/main.js"></script>
		`)),
		Body(children...),
	)
}

var index = router.GET("/").Handle(func(c *gin.Context) g.Node {
	return Div(
		Text("Hello world!"))
})

var db *sqlx.DB

func init() {
	godotenv.Load()
	dbUrl := os.Getenv("DATABASE_URL")

	var err error
	db, err = sqlx.Connect("postgres", dbUrl)
	if err != nil {
		panic(err)
	}
}

func main() {
	router.Engine().Run(":3000")
}

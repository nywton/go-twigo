func main() {
  app := echo.New()
  app.Static("/static", fsRoot: "static")

  articlesService := service.ArticlesService{}

  basicHandler := handler.BasicHandler{ArticlesService: articlesService}
  articlesHandler := handler.ArticlesHandler{ArticlesService: articlesService}

  app.GET("/", basicHandler.ShowHome)
  app.GET("/write", basicHandler.ShowWrite)
  app.GET("/articles/:id", articlesHandler.ShowArticle)
  app.GET("/articles", articlesHandler.CreateArticle)

  port := os.Getenv("PORT")
  if port == "" {
    port =  ":8000"
  }

  app.Logger.Fatal(app.Start(port))
}

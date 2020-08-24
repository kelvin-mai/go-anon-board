package server

func Init() {
	r := newRouter()
	r.LoadHTMLGlob("templates/*")
	setUpRoutes(r)
	db := DbConnect()
	defer db.Close()
	r.Run()
}

package app
import (
	"fmt"
	"log"
	"net/http"
	"github.com/api/app/handler"
	"github.com/api/app/model"
	"github.com/api/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)
// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}
// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Name,
		config.DB.Charset)
	db, err := gorm.Open(config.DB.Dialect, dbURI)


	if err != nil {
		panic(err)
		log.Fatal("Could not connect database")
	}
	fmt.Println("mysql connecting..")
	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}
// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/users", a.GetAllUsers)
	a.Post("/users", a.CreateUser)
	a.Get("/users/id", a.GetUser)
	a.Put("/users/id", a.UpdateUser)
	a.Delete("/users/id", a.DeleteUser)

}
// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}
// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}
// Handlers to manage User Data
func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.DB, w, r)
}
func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	handler.CreateUser(a.DB, w, r)
}
func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(a.DB, w, r)
}
func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	handler.UpdateUser(a.DB, w, r)
}
func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUser(a.DB, w, r)
}
// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
	fmt.Println("Listning on localhost:8080")
}

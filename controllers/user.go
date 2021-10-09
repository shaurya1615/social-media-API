package controllers
//go mod init golang
import (
  "fmt"
  "time"
  "crypto/md5"
  "encoding/json"
  "github.com/julienschmidt/httprouter"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "net/http"
  "golang/models"
)




//CREATING A CONNECTION
type UserController struct{
  session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
  return &UserController{s}
}




//CREATE, GET AND DELETE USER DETAILS
func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
  id := p.ByName("id")
  if !bson.IsObjectIdHex(id){
    w.WriteHeader(http.StatusNotFound)
  }
  oid := bson.ObjectIdHex(id)
  u := models.User{}
  if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil{
    w.WriteHeader(404)
    return
  }
  uj, err := json.Marshal(u)
  if err != nil{
    fmt.Println(err)
  }
  w.Header().Set("Content-Type","application/json")
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n" , uj)
}

func (uc UserController) CreateUser (w http.ResponseWriter, r *http.Request, _  httprouter.Params){
  u := models.User{}
  json.NewDecoder(r.Body).Decode(&u)
  u.Id = bson.NewObjectId()
  data := []byte(u.Password)
  b := md5.Sum(data) // USING MD5 ENCYPTION TO ENCRYPT AND STORE THE PASSWORD, USER PASSWORD WILL AUTOMATICALLY BE ENCRYPTED AS SOON AS IT IS ENTERED AND ONLY THE ENCRYPTED PASSWORD WILL BE STORED
  pass := fmt.Sprintf("%x", b)
  u.Password = pass
  uc.session.DB("mongo-golang").C("users").Insert(u)
  uj, err := json.Marshal(u)
  if err != nil{
    fmt.Println(err)
  }
  w.Header().Set("Content-Type","application/json")
  w.WriteHeader(http.StatusCreated)
  fmt.Fprintf(w, "%s\n",uj)
}

func (uc UserController) DeleteUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
  id := p.ByName("id")
  if !bson.IsObjectIdHex(id){
    w.WriteHeader(http.StatusNotFound)
  }
  oid := bson.ObjectIdHex(id)
  if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil{
    w.WriteHeader(404)
  }
  w.WriteHeader(http.StatusOK)
  fmt.Fprint(w,"Deleted User", oid, "\n")
}




//CREATE, GET AND DELETE POST DETAILS
func (uc UserController) GetPost (w http.ResponseWriter, r *http.Request, p httprouter.Params){
  id := p.ByName("id")
  if !bson.IsObjectIdHex(id){
    w.WriteHeader(http.StatusNotFound)
  }
  oid := bson.ObjectIdHex(id)
  u := models.Post{}
  if err := uc.session.DB("mongo-golang").C("posts").FindId(oid).One(&u); err != nil{
    w.WriteHeader(404)
    return
  }
  uj, err := json.Marshal(u)
  if err != nil{
    fmt.Println(err)
  }
  w.Header().Set("Content-Type","application/json")
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n" , uj)
}

func (uc UserController) CreatePost (w http.ResponseWriter, r *http.Request, _  httprouter.Params){
  u := models.Post{}
  json.NewDecoder(r.Body).Decode(&u)
  u.Id = bson.NewObjectId()
  time := time.Now()
  u.Time = time.String()
  uc.session.DB("mongo-golang").C("posts").Insert(u)
  uj, err := json.Marshal(u)
  if err != nil{
    fmt.Println(err)
  }
  w.Header().Set("Content-Type","application/json")
  w.WriteHeader(http.StatusCreated)
  fmt.Fprintf(w, "%s\n",uj)
}

func (uc UserController) DeletePost (w http.ResponseWriter, r *http.Request, p httprouter.Params){
  id := p.ByName("id")
  if !bson.IsObjectIdHex(id){
    w.WriteHeader(http.StatusNotFound)
  }
  oid := bson.ObjectIdHex(id)
  if err := uc.session.DB("mongo-golang").C("posts").RemoveId(oid); err != nil{
    w.WriteHeader(404)
  }
  w.WriteHeader(http.StatusOK)
  fmt.Fprint(w,"Deleted User", oid, "\n")
}




//GET ALL POSTS FOR A PARTICULAR USER
func (uc UserController) GetAllPost (w http.ResponseWriter, r *http.Request, p httprouter.Params){
  id := p.ByName("id")
  if !bson.IsObjectIdHex(id){
    w.WriteHeader(http.StatusNotFound)
  }
  u := make([]models.Post, 0, 0)
  if err := uc.session.DB("mongo-golang").C("posts").Find(bson.M{"userid" : id}).Limit(10).All(&u); err != nil{ //RESTRICTING THE SYSTEM TO ONLY DISPLAY THE FIRST 10 POSTS
    w.WriteHeader(404)
    return
  }
  uj, err := json.Marshal(u)
  if err != nil{
    fmt.Println(err)
  }
  w.Header().Set("Content-Type","application/json")
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n" , uj)
}

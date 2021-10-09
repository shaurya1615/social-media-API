package main

import (
  "github.com/julienschmidt/httprouter"
  "gopkg.in/mgo.v2"
  "net/http"
  "golang/controllers"
)

func main(){
  r := httprouter.New()
  uc := controllers.NewUserController(getsession())
  r.GET("/user/:id", uc.GetUser)
  r.POST("/user", uc.CreateUser)
  r.DELETE("/user/:id", uc.DeleteUser)
  r.GET("/post/:id", uc.GetPost)
  r.POST("/post", uc.CreatePost)
  r.DELETE("/post/:id", uc.DeletePost)
  r.GET("/posts/users/:id", uc.GetAllPost)
  http.ListenAndServe("localhost:3000", r)
}

func getsession() *mgo.Session{
  s, err := mgo.Dial("mongodb://localhost:27017")
  if err != nil{
    panic(err)
  }
  return s
}

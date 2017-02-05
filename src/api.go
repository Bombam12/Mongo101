package main

import (
	"net/http"

	 "github.com/labstack/echo"
	 "github.com/labstack/echo/middleware"
	 "gopkg.in/mgo.v2"
   "gopkg.in/mgo.v2/bson"

	 db "./helper/db"
	 "./models"
)
 var (
	 MongoSession *mgo.Session
	 UsersCollection *mgo.Collection
 )

func index(c echo.Context) error  {
  	return c.JSON(http.StatusOK, "Hello, World!")
}

func getUsers(c echo.Context) error  {
  user := new(models.User)
	result, _ := user.ReadFromDB()
  return c.JSON(http.StatusOK,result)
}
func deleteUserByID(c echo.Context) error  {
	user := new(models.User)
  id:=c.Param("id")
	user.Id = bson.ObjectIdHex(id)
	user.DeleteUserByID()
  	return c.NoContent(http.StatusOK)
}
func getUserByID(c echo.Context) error  {
	user := new(models.User)
  id:=c.Param("id")
	user.Id = bson.ObjectIdHex(id)
	result, _ := user.ReadFromDBByID()
  	return c.JSON(http.StatusOK, result)
}

func saveUser(c echo.Context) error  {
     user:=new(models.User)
     err:= c.Bind(user)

     if err != nil{
        //  myName:= User{
        //      "Pornthip",
        //      "Soonthorn",
        //   }
        // return c.JSON(http.StatusOK, myName)
       return c.NoContent(http.StatusCreated)
     }
     user.SaveToDB()
     return c.JSON(http.StatusOK, user)
}



func init()  {
	mongoSession, err := mgo.Dial("localhost:27017")
	  if err != nil {
			panic(err)
		}
		mongoSession.SetMode(mgo.Monotonic, true)
		db.MongoSession = mongoSession
		db.UsersCollection = db.MongoSession.DB("maejo").C("users")

}

func main() {
defer db.MongoSession.Close()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		  AllowOrigins: []string{"*"},
		  AllowHeaders: []string{
				echo.GET,
				echo.POST,
			},
	}))
  e.GET("/index",index)
  e.GET("/users",getUsers)
  e.GET("/users/:id",getUserByID)
  e.POST("/users",saveUser)
	e.DELETE("/users/:id",deleteUserByID)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

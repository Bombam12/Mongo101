package main

import (
	"net/http"

	 "github.com/labstack/echo"
	 "gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)
 var (
	 MongoSession *mgo.Session
	 UsersCollection *mgo.Collection
 )

type User struct{
  First string  `json:"first"`
  Last string   `json:"last"`
}

func (u *User)SaveToDB() error  { //this is call method recever
    err := UsersCollection.Insert(&u)
		if err != nil{
			return err
		}
		return nil
}

func (u *User)ReadFromDB()  ([]User, error){
	   result := []User{}
			err := UsersCollection.Find(nil).All(&result)
			if err != nil{
				return nil, err
			}
			return result, nil
		}

func index(c echo.Context) error  {
  	return c.JSON(http.StatusOK, "Hello, World!")
}

func getUsers(c echo.Context) error  {
  user := new(User)
	result, _ := user.ReadFromDB()
  return c.JSON(http.StatusOK,result)
}

func getUserByID(c echo.Context) error  {
  id:=c.Param("id")
  	return c.JSON(http.StatusOK,id)
}

func saveUser(c echo.Context) error  {
     user:=new(User)
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
	MongoSession, err := mgo.Dial("localhost:27017")
	  if err != nil {
			panic(err)
		}
		MongoSession.SetMode(mgo.Monotonic, true)
		UsersCollection = MongoSession.DB("maejo").C("users")
}

func main() {
defer MongoSession.Close()

	e := echo.New()

  e.GET("/index",index)
  e.GET("/users",getUsers)
  e.GET("/users/:id",getUserByID)
  e.POST("/users",saveUser)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

package models

 import(
   	 	  "gopkg.in/mgo.v2/bson"

       db "../helper/db"
 )

type User struct{
  Id bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
  First string  `json:"first,omitemptyomitempty" bson:"first,omitempty"`
  Last string   `json:"last,omitempty" bson:"last,omitempty"`
  Image string   `json:"image,omitempty" bson:"image,omitempty"`
  Detail string   `json:"detail,omitempty" bson:"detail,omitempty"`
}

func (u *User)SaveToDB() error  { //this is call method recever
    err := db.UsersCollection.Insert(&u)
		if err != nil{
			return err
		}
		return nil
}

func (u *User)ReadFromDB()  ([]User, error){
	   result := []User{}
			err := db.UsersCollection.Find(nil).All(&result)
			if err != nil{
				return nil, err
			}
			return result, nil
		}

    func (u *User)ReadFromDBByID()(*User, error){
      err := db.UsersCollection.Find(bson.M{"_id": u.Id}).One(&u)
      if err != nil{
        return nil, err
      }
      return u, nil
    }
    func (u *User)DeleteUserByID()(*User, error){
      err := db.UsersCollection.RemoveId(u.Id)
      if err != nil{
        return nil, err
      }
      return u, nil
    }

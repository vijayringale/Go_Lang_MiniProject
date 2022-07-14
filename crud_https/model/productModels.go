package model

import (
	"CRUD_HTTPS/entiti"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type  ProductModel struct {
	Db *mgo.Database
	Collection string
}

// func NewUserController( s *mgo.Session) *UserController{
// 	return &UserController{s}
// }

func (uc ProductModel) GetAllUser()( pd [] entiti.Product , err error){
	err =uc.Db.C(uc.Collection).Find(bson.M{}).All(&pd)
	return	
}


func (uc ProductModel) GetUser(id string)( pd  entiti.Product , err error){
	err =uc.Db.C(uc.Collection).FindId(bson.ObjectIdHex(id)).One(&pd)
	return	
}

func (uc ProductModel) UpdateUser( pd  *entiti.Product )error{
	err :=uc.Db.C(uc.Collection).UpdateId(pd.Id , &pd)
	return	err
}

func (uc ProductModel) AddUser( pd  *entiti.Product )error{
	err :=uc.Db.C(uc.Collection).Insert(&pd)
	return	err
}

func (uc ProductModel) DeleteUser( pd entiti.Product)error{
	err :=uc.Db.C(uc.Collection).Remove(pd)
	return	err
}

package main

import (
	"CRUD_HTTPS/config"
	"CRUD_HTTPS/entiti"

	"CRUD_HTTPS/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Getall",FindAll).Methods("GET")
	r.HandleFunc("/Get/{name}",FindOne).Methods("GET")
	r.HandleFunc("/add",Add).Methods("POST")
	r.HandleFunc("/Update",Update).Methods("PUT")
	r.HandleFunc("/delete/{id}",Delete).Methods("DELETE")





	fmt.Println("RR : Server is running")

	err := http.ListenAndServe(":4499",r)
	if err != nil {
		fmt.Println(err)

	}else{
		fmt.Println("Server RR Is Running On Port : 4499")
	}
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		ResponseWithError(w ,http.StatusBadRequest , err.Error())
		return
	}else{
		prM :=model.ProductModel{
			Db: db,
			Collection : "users",

		}	

		products , err2 := prM.GetAllUser()
		if err2!=nil {
			ResponseWithError(w , http.StatusBadRequest , err2.Error())
			return
		}else{
			respondWithJson(w , http.StatusOK,products)
		}
	}
}

func FindOne(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		ResponseWithError(w ,http.StatusBadRequest , err.Error())
		return
	}else{
		prM :=model.ProductModel{
			Db: db,
			Collection : "users",

		}	
		vars := mux.Vars(r)
		id:= vars["name"]

		product , err2 := prM.GetUser(id)
		fmt.Print(product , id)
		if err2!=nil {
			ResponseWithError(w , http.StatusBadRequest , err2.Error())
			return
		}else{
			respondWithJson(w , http.StatusOK,product)
		}
	}
}


func Add(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		ResponseWithError(w ,http.StatusBadRequest , err.Error())
		return
	}else{
		prM :=model.ProductModel{
			Db: db,
			Collection : "users",

		}	
		var product entiti.Product
		product.Id =bson.NewObjectId()
		err := json.NewDecoder(r.Body).Decode(&product)

		if err !=nil {
			ResponseWithError(w , http.StatusBadRequest , err.Error())
			return
		}else{
			prM.AddUser(&product)
			respondWithJson(w ,http.StatusOK , product)
		}
	}
}


func Update(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		ResponseWithError(w ,http.StatusBadRequest , err.Error())
		return
	}else{
		prM :=model.ProductModel{
			Db: db,
			Collection : "users",

		}	
		var product entiti.Product
		// product.Id =bson.NewObjectId()
		err := json.NewDecoder(r.Body).Decode(&product)

		if err !=nil {
			ResponseWithError(w , http.StatusBadRequest , err.Error())
			return
		}else{
			prM.UpdateUser(&product)
			respondWithJson(w ,http.StatusOK , product)
		}
	}
}




func Delete(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		ResponseWithError(w ,http.StatusBadRequest , err.Error())
		return
	}else{
		prM :=model.ProductModel{
			Db: db,
			Collection : "users",

		}	
		vars := mux.Vars(r)
		id:= vars["id"]
		deletingValue,_ := prM.GetUser(id)
		err2 := prM.DeleteUser(deletingValue)
		if err2!=nil {
			ResponseWithError(w , http.StatusBadRequest , err2.Error())
			return
		}else{
			respondWithJson(w , http.StatusOK,entiti.Product{})
		}
	}
}



func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
package main

import(
  "log"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/gorilla/mux"
)

//create the structure for Server data type
type Server struct {
  ID     string `json:"id, omitempty"`
  Name   string `json:"name, omitempty"`
  Cores  string `json:"cores, omitempty"`
  Memory string `json:"memory, omitempty"`
  Disk   string `json:"disk, omitempty"`
}

//create an instance
var servers []Server

//Welcome message
func ApiRunning(){
  log.Println("The API is running.")
}

//Test to know if the API is ok
func ApiTest(w http.ResponseWriter, r *http.Request) {
  log.Println("Executing test")

  response := "The API is Running. Everything is OK."
  out, _ := json.Marshal(response)
  w.Write(out)
}

//show all servers
func GetAllServers(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  out, _ := json.Marshal(servers)
  w.Write(out)
  log.Println("All servers have been requested.")
}


//show the requested server
func GetOneServer(w http.ResponseWriter, r *http.Request){
  //get id
  params := mux.Vars(r)

  //search the server with id
  for _, item := range servers{
    if item.ID == params["id"]{
      w.Header().Set("Content-Type", "application/json")
      out, _ := json.Marshal(item)
      w.Write(out)
      log.Println("The server", item,"has been requested.")
      return
    }
  }

  response := "The server requested does not exist."
  out, _ := json.Marshal(response)
  w.Write(out)
  log.Println("The server requested does not exist.")
}

//create a new server
func CreateServer(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  //create an instance and get the input data
  var serv Server
  in, _ := ioutil.ReadAll(r.Body)
  json.Unmarshal(in, &serv)

  //save the server if the ID is not empty
  if serv.ID != "" {
    servers = append(servers, serv)
    out, _ := json.Marshal(serv)
    w.Write(out)
    log.Println("The server", serv, "has been created.")
  }else {
    log.Println("Can't create the server. Don't accept empty ID.")
  }
}

//edit all data set for one server
func EditServer(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  //create an instance and get the input data
  var serv Server
  in, _ := ioutil.ReadAll(r.Body)
  json.Unmarshal(in, &serv)

  //search the server and save the new data set
  for index, item := range servers{
    if item.ID == serv.ID {
      servers[index] = serv
      out, _ := json.Marshal(servers[index])
      w.Write(out)
      log.Println("The edition has been applied. New data:", servers[index])
    }
  }
}

//edit some resource
func EditResource(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r) //get id as parameter

  var serv Server //get resource
  in, _ := ioutil.ReadAll(r.Body)
  json.Unmarshal(in, &serv)

  //search which resorce has changed and save new data
  for index, item := range servers{
    if item.ID == params["id"] {
      if serv.Name != ""{
        item.Name = serv.Name
      }
      if serv.Cores != ""{
        item.Cores = serv.Cores
      }
      if serv.Memory != ""{
        item.Memory = serv.Memory
      }
      if serv.Disk != ""{
        item.Disk = serv.Disk
      }
      servers[index] = item
      out, _ := json.Marshal(servers[index])
      w.Write(out)
      log.Println("The patch has been applied. New data:", servers[index])
    }
  }
}

//delete the selected server
func DeleteServer(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  //get the ID
  params := mux.Vars(r)
  for index, item := range servers{
    if item.ID == params["id"]{
      log.Println("The server", item, "will be deleted.")
      servers = append(servers[:index], servers[index+1:]...)
      out, _ := json.Marshal(servers)
      w.Write(out)
      return
    }
  }
}

func main() {
  router := mux.NewRouter() //open a new router

  //welcome message
  ApiRunning()

  //Examples
  servers = append(servers, Server{ID: "1", Name: "vxhco-23", Cores: "10", Memory: "32GB", Disk: "100GB"})
  servers = append(servers, Server{ID: "2", Name: "vxadi-02", Cores: "2", Memory: "4GB", Disk: "50GB"})
  servers = append(servers, Server{ID: "3", Name: "fnadh-15", Cores: "8", Memory: "16GB", Disk: "150GB"})

  //declare the function realted to each method
  router.HandleFunc("/test", ApiTest).Methods("GET")
  router.HandleFunc("/servers", GetAllServers).Methods("GET")
  router.HandleFunc("/servers/{id}", GetOneServer).Methods("GET")
  router.HandleFunc("/servers/{id}", CreateServer).Methods("POST")
  router.HandleFunc("/servers/{id}", EditServer).Methods("PUT")
  router.HandleFunc("/servers/{id}", EditResource).Methods("PATCH")
  router.HandleFunc("/servers/{id}", DeleteServer).Methods("DELETE")

  //open port 8080
  log.Fatal(http.ListenAndServe(":8080", router))
}

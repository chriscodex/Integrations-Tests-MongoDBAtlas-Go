# Integration Tests to MongoDBAtlas with Go ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ChrisCodeX/CRUD-MongoDBAtlas-Go) ![](https://camo.githubusercontent.com/3084f133857f6d0a29d410e59ba39f6906b0f2e32b24082d1e95710196984db6/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f2d4d6f6e676f44422d3444423333443f7374796c653d666c6174266c6f676f3d6d6f6e676f6462266c6f676f436f6c6f723d464646464646)
This repository can help you to integrate a MongoDBAtlas cluster database into your Go application and test the operations of creating, reading, updating, and deleting documents from a collection.  

![ITMDBG](https://user-images.githubusercontent.com/106860308/178126005-e03589d7-08ea-4b38-8897-86ad38b85f19.png)

---

### Pre-Requirements 📋  
Have access to a MongoDBAtlas cluster.  
In this link you have a guide on how to create it: https://www.mongodb.com/docs/atlas/tutorial/create-new-cluster/  

---  

### Instalation 🔧  
- Install the MongoDB driver as a dependency  
```
go get go.mongodb.org/mongo-driver/mongo  
```
- Inside the database folder, modify the `mongodb.con.go` file. Add the uri of your cluster and the name of your database  
```
var (
	// Add the uri of your cluster of MongoDBAtlas
	// https://www.mongodb.com/docs/drivers/go/current/quick-start/#connect-to-your-cluster
	uri = ""
	// Add the name of your database
	database = ""
)
```

---

### Running Tests ⚙️
To run the test, type this command:
```
go test
```
- To verify the connection, run the test inside `/database` directory.
- To test CRUD operations, run the test inside `/service/user.service` directory.  
By default it will create the "users" collection, if you want to change the collection, modify the `user.repository.go` file located inside `/repositories/user.repository/`  
```
// Name of collection
var collection = database.GetCollection("users")
```

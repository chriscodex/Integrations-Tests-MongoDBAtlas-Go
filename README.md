# Integration Tests to MongoDBAtlas with Go
This repository can help you integrate a MongoDBAtlas cluster database into your Go application and test the operations of creating, reading, updating, and deleting documents from a collection.  

---

### Pre-requirements 📋  
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

### Running the test ⚙️
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

---

### Built with 🛠️

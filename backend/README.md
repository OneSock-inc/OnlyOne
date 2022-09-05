# User instruction
(you just want to use the service you don't care how it works)  
You need to have a working [Go](https://go.dev/learn/) installation to run this project.

You can check it with 
```bash
$ go version
#go version go1.19 darwin/arm64
```
You should see something similar. This project is done with Go 1.19 and modules enabled.
### Set up the database
We are using cloud firestore as our database. You'll need to setup a google cloud firestore project (not a google firebase firestore project).
For that just follow this [article](https://cloud.google.com/firestore/docs/quickstart-servers#set_up_your_project). 
The `service-account.json` is expected to be in the home of whichever user run the server : `~/service-account.json`.
You can configure the database connection in the `db/db.go` at the top of the file. Specificaly you **MUST** change the project id (`var projectID string = "your_project_id_here"`) variable in `db/db.go`.
_You can find your project id in the firebase console._
If you want to place the `service-account.json` somewhere else you can change the `createClient` function in `db/db.go` to use the `option.WithCredentialsFile("path/to/service-account.json")` option.


## Running the project
Go to the onlyOne/backend folder (you should see the `main.go` file there) and run
```bash
$ go mod tidy -v
#might take some time (downloads all the dependancies)
go run .
```
And you can now make api call to the backend on `localhost:8000`. 
API is documented in the [notion](https://www.notion.so/webblitchy/Sp-cification-API-049c86e231324f048c2d2569b49a30ac). You might want to download [Postman](https://www.postman.com/) to make api calls.



# Developper information

## Project's description 
In this folder you can see the backend of the onlyOne project.
The project is a reponsive website where someone can signup, login and propose lonely sock to match. We try to match two lonely sock together to form a new happy pair. 
In this folder you can see the backend server code. it's written in Go/Golang using the [gin](https://gin-gonic.com/) framework.

## Structure of the backend's code

The backend is divided in two parts (go modules):
 1. API
 2. DB  


The api module is composed of the code setting up the routes and the handlers for the requests/response. The JWT middleware used for the authentification is also defined there.

Most of these handlers just do a call to the DB module where the heavy lifting (db transaction, commit rollback in case of error) is done. Then the handlers return the response to the client or the error .
The databases is totaly abstracted from the API viewpoint

The DB module is where all the calls to the database are made. We use the firestore [library](https://pkg.go.dev/google.golang.org/cloud/firestore) . The matching algorithme is implemented in this module currently this is a [knn](https://en.wikipedia.org/wiki/K-nearest_neighbors_algorithm) search in a [kdtree](https://github.com/sjwhitworth/golearn/blob/master/kdtree/kdtree.go). We also take care of the error handling related to the db there, for exemple if the data send by a user to a route / response handler is ill formed or simply missing we detect it there before saving it to the database. 


## Running the tests
The tests are run on a emulated database, in order to run them download the firebase CLI [here](https://firebaseopensource.com/projects/firebase/firebase-tools/).
Then add the following files in the `backend` folder:
- `firebase.json`
```json
{
  "emulators": {
    "firestore": {
      "port": 8080
    },
    "ui": {
      "enabled": true
    }
  }
}
```
- `.firebaserc`
```json
{
  "projects": {
    "default": "put-your-project-id-here"
  }
}
```
_You can find your project id in the firebase console._


 and run the following command while beeing in the `backend` folder
```bash
$ go test ./...
# ?       backend [no test files]
# ok      backend/api     1.408s
# ok      backend/db      1.040s
# ?       backend/utils   [no test files]
```
You should see something similar to the output above. If you want to see the test coverage you can run
```bash
$ go test -cover ./...
# ?       backend [no test files]
# ok      backend/api     1.343s  coverage: 77.6% of statements
# ok      backend/db      0.938s  coverage: 75.6% of statements
# ?       backend/utils   [no test files]
```

# GoLang Seminary
### Simple REST API with the topic of a Vet that includes the following topics:
- CRUD of animals
- Utilization of SQLite as persistence
- Usage of Gin Framework
- Employment of Yaml      
  

## To run the application follow these straightforward steps:
#### **Prerequisites:** [golang](https://golang.org/) installed
1. Clone the repository in your $GOPATH/src. Create the *src* folder if necessary  
(in Windows, the $GOPATH by default is C:/Users/$YOURUSER/go)
2. Open it in your favorite IDE  
(I use [Visual Studio Code](https://code.visualstudio.com/) but you can use other options: [Atom](https://atom.io/), [Sublime Text](https://www.sublimetext.com/3) or others)
3. Open the console and write the command:
~~~
go run cmd/vet/vetsrv.go -config ./config/config.yaml
~~~
4. (Note) I have tested it on Windows 10 and, after running, a Firewall alert opens. Just click in **Accept**. It´s not a virus, trust me..
5. Congratulations, you can start testing the API with [Postman](https://mongusteam.postman.co/home) or a similar software. The posible operations are the following:
- Get all animals: 
    - **URL:** *http://localhost:8080/animals*
    - **Method:** GET
    - Doesn't require body
- Get one animal:
    -  **URL:** *http://localhost:8080/animal*
    -  **Method:** GET
    -  Doesn't require body
- Add one animal: 
    - **URL:** *http://localhost:8080/animal*
    - **Method:** POST
    - Requires body. Example: 
      ~~~
      { "Name": "Risonante II", "Age": 21 }
      ~~~
- Update one animal: 
    - **URL:** *http://localhost:8080/animal/{id}*
    - **Method:** UPDATE
    - Requires body. Example: 
      ~~~
      { "Name": "Capitán Coscacho", "Age": 7 }
      ~~~
- Delete one animal: 
    - **URL:** *http://localhost:8080/animal/{id}*
    - **Method:** DELETE
    - Doesn't require body.

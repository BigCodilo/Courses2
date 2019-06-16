1. Testing functions located in folder "testAPI" (package "main") (used standart test library).
2. `go test -v` will start testing
3. You can run "testAPI/testapi.go" after run server, this package will send 4 requestes to server to
  /add /delete/ /update /persons -> _POST, POST, POST, GET_, in first 3 times expected strings answer, and in 4 time list of persons 
4. Logger in package "logger"
5. Struct of database table of persons: 
   

    Person{
        	ID           
        	FirstName  
        	LastName     
        	Email       
        	Gender       
        	GenderIota   
        	RegisterDate 
        	Loan         
    }
    
    
 


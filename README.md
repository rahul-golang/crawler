# crawler App with no of words in APP 

 ## Problem Statment

        Build web UI per below. It has a text box to enter a URL. After URL submission, it should:
        - Grab the content of URL
        - Strips all HTML tag thus only the real-content left
        - Get the count of all words and presented in tabular manner as below




## Project Sructure

        .
        ├── bin                          <--  after "make run" command it will generate binary here
        │   └── url-word-count
        ├── handlers                     <--  Project handlers
        │   └── app_handler.go
        ├── main.go                      <--  Entrypoint
        ├── Makefile
        ├── README.md
        ├── services                     <--  Project Service buisness logic to find word count here
        │   └── app_service.go
        └── templates                    <--  Project UI Html Templates are here
            └── layout.html
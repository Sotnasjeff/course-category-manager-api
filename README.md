# Project Title: Course-Category-Manager-Api

# Description:
 - This is a project in order to apply the knowledges I've got it about GRPC, persisting data in a database and using unary and binary communication.

 # Prerequisites to run locally:
 - Make the following steps to do it:
    - Install Go in your machine: https://go.dev/doc/install 
    - Install Protobuf Compiler: https://grpc.io/docs/protoc-installation/
    - Install Protoc-gen-go and protoc-gen-go-grpc GO plugins: https://grpc.io/docs/languages/go/quickstart/
    - Install Sqlite3: https://www.tutorialspoint.com/sqlite/sqlite_installation.htm
    - Install Evans Client (GRPC's client to make request to server): https://github.com/ktr0731/evans
    - Run the commands created on Makefile:
        - First, type in your terminal: make proto, to remove the protofiles and generate a new protofile
        - In your terminal type: Go Mod Tidy to install the dependencies not downloaded.
        - After updated the project, type in terminal: Make Server
        - After running the server, open a new tab in terminal and type: Make Evans
        - In Evans, you can make the call for the methods to test it and send request.

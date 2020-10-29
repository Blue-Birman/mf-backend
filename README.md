# Millenial Fashion

### Tech

Millenial Fashion uses a number of open source projects to work properly:

* [Go](https://golang.org) - Open source programming language that makes it easy to build simple, reliable, and efficient software
* [Gin](https://gin-gonic.com/docs/) - Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API, but with performance up to 40 times faster than Martini. If you need smashing performance, get yourself some Gin.

### Installation

Install the dependencies and devDependencies and start the server.

#### With Cli:
```sh
$ mysql_username=YOUR_USERNAME mysql_password=YOUR_PASSWORD mysql_host=localhost mysql_schema=YOUR_DB_NAME go run main.go
```

#### With VSCode Debugger
1. Create .vscode/launch.json file 
```
    {
      // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
      "version": "0.2.0",
      "configurations": [
        {
          "name": "Launch",
          "type": "go",
          "request": "launch",
          "mode": "auto",
          "program": "${workspaceFolder}",
          "envFile": "${workspaceFolder}/.env",
          "args": []
        }
      ]
    }
```
2. Create .env file in project 
```
    mysql_username=YOUR_USERNAME
    mysql_password=YOUR_PASSWORD
    mysql_host=localhost
    mysql_schema=YOUR_DB_NAME
```
# Group Members

> 1118007 - Daniel Christianto  
> 1118011 - Hanjaya Suryalim  
> 1118021 - Michelle Natasha Irawan  
> 1118023 - Daniel Alexander  
> 1118036 - Raffi Verrel Alessandro


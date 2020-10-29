# mf-backend

## To Run:

### With CLI:
`mysql_username=YOUR_USERNAME mysql_password=YOUR_PASSWORD mysql_host=localhost mysql_schema=YOUR_DB_NAME go run main.go`

### With VSCode Debugger:
1. Create .vscode/launch.json file
  `
    {
      // Use IntelliSense to learn about possible attributes.
      // Hover to view descriptions of existing attributes.
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
  `
 2. Create .env file in project root
    mysql_username=YOUR_USERNAME
    mysql_password=YOUR_PASSWORD
    mysql_host=localhost
    mysql_schema=YOUR_DB_NAME
3. Press F5 to run your VSCode debugger

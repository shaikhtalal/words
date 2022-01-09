# Words 

## To use makefile install make on your machine (e.g In Ubuntu)

```sh
 sudo apt install make
```
## For developer help (Use makefile)
```sh
  make help (list the command)
```
```sh
Available targets are:

    run-server                         Run the backend service
```

## Endpoint for counting words
```sh
  POST :  http://localhost:9112/words
```

```sh
#request input
{
    "text": ""
    "limit":0
}
```

## Run the backend without Makefile
```sh
go run main.go serve
```

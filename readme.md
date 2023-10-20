### Ejecutando script JS desde Golang

```bash
docker build -t rungo .
docker run rungo

docker ps
docker exec -it ELHASH bash

go run .
```


### Enviando codigo desde el cliente al servidor

1. Iniciamos el servidor con `go run .`
2. Vamos a __postman__ y ponemos via __form-data__ lo siguiente:
```go
package main
import (
	"fmt"
)
func main() {
	fmt.Println("Hola mundo!!")
}
```
3. No importa el verbo http, mandale hacia `localhost:8080`

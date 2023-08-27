### Installation du swagger

- Téléchargeons Swag pour Go :

```
go install github.com/swaggo/swag/cmd/swag@latest
```

ou la version binaire (version 1.8.12)

```
wget -c https://github.com/swaggo/swag/releases/download/v1.8.12/swag_1.8.12_Linux_x86_64.tar.gz

sudo mv swag /usr/local/bin/

swag -v
```

- Exécutons **Swag** à la racine de notre projet Go

```
swag init
```

- Téléchargeons **gin-swagger** 

```
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

- Intégrons **gin-swagger** dans notre fichier main

```
import (
    "github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"hello-world/docs" // docs est généré par Swag CLI, nous devons l'importer.
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
}
```

### Installation de Mongo

```
docker run -d --name mongodb -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password -p 27017:27017 mongo:4.4.24
```

### Installation du driver Mongo (version 1.12.1)

```
go get go.mongodb.org/mongo-driver/mongo
```
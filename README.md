#### Installation du framework Gin

Initialisation d'un projet **hello-world**

```
mkdir ~/hello-world
cd ~/hello-world
go mod init hello-world
```

Pour installer le package Gin pour notre projet **hello-world**, nous exécutons la commande :

```
go get -u github.com/gin-gonic/gin
```

Pour importer Gin dans notre code :

```
import "github.com/gin-gonic/gin"
```

(Facultatif) nous pouvons importer **net/http**. Ceci est requis par exemple si nous utilisons des constantes telles que **http.StatusOK**

```
import "net/http"
```

Pour lister toutes les dépendances à notre projet **hello-world** 

```
go list -m all
```

Pour supprimer les dépendances inutilisées, nous pouvons utiliser la commande 

```
go mod tidy
```

Tidy s'assure que **go.mod** correspond au code source du module. Il ajoute tous les modules manquants nécessaires pour construire les packages et les dépendances du module actuel, et il supprime les modules inutilisés qui ne fournissent aucun package pertinent. Il ajoute également toutes les entrées manquantes à go.sum et supprime celles qui ne sont pas nécessaires.

<br>

Il est utile de stocker les modules ou les packages tiers dont dépend notre projet et de les placer dans un dossier (vendor), afin qu'ils puissent être archivés dans le contrôle de version. Heureusement, les modules Go prennent en charge vendor :

```
go mod vendor
```

Vendor réinitialise le répertoire du fournisseur du module main pour inclure tous les packages nécessaires pour construire et tester tous les packages du module main. Il n'inclut pas le code de test pour les packages de vendor.
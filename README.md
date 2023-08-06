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

<br>

Nous pouvons utiliser la commande **go mod graph** pour afficher la liste des modules dans le fichier go.mod :

```
go mod graph | sed -Ee 's/@[^[:blank:]]+//g' | sort | uniq > unver.txt
```

Nous créeons le fichier **graph.dot** avec pour contenu :

```
digraph {
    graph [overlap=false, size=14];
    root="hello-world";
    node [ shape = plaintext, fontname = "Helvetica", fontsize=24];
    "hello-world" [style = filled, fillcolor = "#E94762"];
```

Nous allons injecter la sortie de **unvert.txt** dans le fichier **graph.dot** avec les commandes suivantes :

```
cat unver.txt | awk '{print "\""$1"\" -> \""$2"\""};' >> graph.dot
echo "}" >> graph.dot
```

Il faudrait enlever toutes les lignes (vers la fin du fichier) commençant par **hello-world** dans le fichier **graph.dot**.
<br>
Nous pouvons maintenant rendre les résultats avec l'outil **Graphviz**. Cet outil peut être installé avec la commande suivantes sous ubuntu20.04

```
sudo apt-get install graphviz
```

Une fois **Graphviz** installé, exécutons la commande suivante pour convertir le fichier **graph.dot** au format **.svg** :

```
sfdp -Tsvg -o graph.svg graph.dot
```

Un fichier **graph.svg** sera généré. Ouvrons le fichier avec la commande suivante :

```
google-chrome ./graph.svg
```

#### Stratégie de versionning

- **main** : Cette branche correspond au code de production courant. Nous ne pouvons pas y pousser du code directement, sauf pour les correctifs. Les tags Git peuvent être utilisées pour baliser tous les commits de la branche **main** avec un numéro de version (par exemple, pour utiliser la convention de version sémantique, https://semver.org/ , qui comporte trois parties : majeure, mineure et patch, donc une balise avec la version 1.2.3 a 1 comme version majeure, 2 comme version mineure et 3 comme version de correctif).
<br>

- **preprod** : il s'agit d'une branche de publication et d'un miroir de la production. Il peut être utilisé pour tester toutes les nouvelles fonctionnalités développées sur la branche **develop** avant qu'elles ne soient fusionnées avec la branche **main**.
<br>

- **develop** : il s'agit de la branche d'intégration de développement, qui contient le dernier code de développement intégré.
<br>

- **feature/X** : il s'agit d'une branche de fonctionnalité individuelle en cours de développement. Chaque nouvelle fonctionnalité réside dans sa propre branche et est généralement créée pour la dernière branche de développement.
<br>

- **hotfix/X** : lorsque nous avons besoin de résoudre quelque chose dans le code de production, nous pouvons utiliser la branche **hotfix** et ouvrir une pull request pour la branche **main**. Cette branche est basée sur la branche **main**.
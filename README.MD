# Présentation

Ce projet est une implémentation de l'algorithme A* en Go. Le but est de comparer les performances de l'algorithme en Go et en C++.

L'implémentation en C++ n'est pas asynchrone contrairement à celle en Go.

[Explication détaillée de l'algorithme](https://fr.wikipedia.org/wiki/Algorithme_A*)

[Interprétation visuelle](https://www.angeluriot.com/maze_solver/)

## Run application
Pour faciliter l'utilisation de l'application, un Makefile est présent à la racine du projet. 

L'application prend en paramètre le chemin vers l'image à analyser avec `img`.


```shell
$ make run img=ressources/first_level.png
```

## Screenshots

[couleur_selon_g](https://i.imgur.com/j8gsEYY.png)

[bidirectionnel](https://i.imgur.com/TmK49hq.png)

## Améliorations
### A* Bidirectionnel
Nous avons pu faire quelques tests sur l'algorithme A* bidirectionnel. Il n'est pas fini et reste buggé.
Cependant, nous avons quand même eux quelques résultats :
 - Sur des très petites images, le lancement de A* en bidirectionnel n'a parfois pas de gain de performance comparé à A* classique. 
   - En effet, selon le scheduler du système, il se peut qu'une routine finisse avant même que l'autre ai commencé.
 - Sur des images plus grandes, on observe un gain allant jusqu'à 70% dans certains cas.

Il serait intéressant de se baser sur ces [travaux](https://core.ac.uk/download/pdf/19184689.pdf) pour finaliser / améliorer l'algorithme A* bidirectionnel.

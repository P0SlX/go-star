# Présentation

Ce projet est une implémentation de l'algorithme A* en go, tout en le comparant avec une implémentation en C++.
Le but est de comparer les performances de l'algorithme en go et en C++. Tout en utilisant la concurrence de go.
Explication détaillé de l'agorithme : https://fr.wikipedia.org/wiki/Algorithme_A*
Intépreation visuel : https://www.angeluriot.com/maze_solver/

## Run application
Pour faciliter l'utilisation de l'application, un Makefile est présent à la racine du projet. 
L'application prend en paramètre le chemin vers l'image à analyser avec **img**.


```shell
$ make run img=ressources/first_level.png
```
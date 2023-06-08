<div align="center">

<br/>

**Gabriel Botelho**

<br/><br/>

**Projet Final - Veille Technologique**

**Travail hybride**

<br/><br/>

**Travail remis à**

**Nicolas Bourré**

<br/><br/>

**Collège Shawinigan**

</div>
<br/><br/>

# Table des matières
1. [Introduction](#Introduction)
2. [Développement](#Développement)

    2.1. [Historique](#Historique)

    2.2. [Méthodologie](#Méthodologie)

    2.3. [Résultats](#Résultats)

    2.4. [Débat](#Débat)

3. [Conclusion](#Conclusion)
4. [Médiagraphie](#Médiagraphie)

## Introduction

En tant que programmeurs, nous savons que la technologie évolue constamment. En effet, nous faisons partie des premiers groupes touchés par cette révolution, principalement parce qu'elle dicte nos outils et nos futurs projets.

Il est de notre devoir d'être capables d'évoluer dans ce domaine et de nous approprier de nouvelles technologies. Bien que nos cours nous exposent à un grand nombre de langages de programmation recherchés de nos jours, ils ne peuvent pas tout nous montrer. Le web, étant un sujet très populaire en 2023, fait partie intégrante de notre réalité. Le langage Go, simple à apprendre, sera le sujet de cette recherche.

Pendant les deux prochaines semaines, j'essaierai de me familiariser avec Go en créant un CRUDL complet, comprenant un système de connexion, le tout relié à une base de données MySQL qui servira également de serveur pour le site web.

![image](https://github.com/Galabott/Veille-Techno/assets/94148198/5f95551a-a6ff-4f7b-b8f2-c13d18f0e178)

## Développement

### Besoins

La recherche portera sur Go et aura pour but de créer un CRUDL simple avec un système de connexion, une liaison avec une base de données MySQL, et une interface utilisant simplement HTML, afin de démontrer les principes sans recourir à une interface console.

Ce projet me permettra d'expérimenter un langage qui n'a pas été abordé en classe et de réaliser un projet CRUDL, facilitant ainsi la comparaison avec Go et me permettant de me forger une opinion sur le sujet.

### Historique

Go est un langage de programmation développé par Google. Créé en 2007, il est devenu public en 2009 et a gagné en popularité depuis le lancement de la version 1 en 2012. Bien qu'il soit récent et conçu pour remédier aux lacunes des anciens langages, il s'inspire fortement du Pascal et du C.

### Méthodologie

Après avoir tiré des enseignements de mes projets précédents présentés en classe et grâce à mes recherches sur Go, j'ai développé une série d'étapes simples à suivre pour mettre en place un CRUDL simple.

1. J'ai commencé par effectuer des recherches sur les possibilités de Go, son origine et ses utilisations.
2. Ensuite, j'ai consulté plusieurs tutoriels, documentations et autres ressources pour me familiariser avec Go.
3. J'ai commencé par créer un système de connexion simple avec un stockage local dans le programme.
4. J'ai expérimenté les connexions à une base de données en effectuant des requêtes "Select" et "Insert" simples.
5. Ensuite, j'ai combiné les deux pour créer un système de connexion utilisant une base de données et des cookies.
6. Enfin, j'ai achevé le projet en créant un CRUDL simple.

### Résultats

Au final, j'ai réussi à réaliser mon projet tel que prévu.

L'interface en HTML est simple et rudimentaire, mais suffisante pour démontrer le fonctionnement du projet. Elle comprend un système de connexion et de création de compte, ainsi qu'un CRUDL simple. Le CRUDL permet d'ajouter, de modifier et de supprimer des balises affichées dans une liste, et les fonctionnalités du CRUDL sont gérées par un formulaire HTML (Créer et Modifier) et des liens (Modifier et Supprimer).

Grâce au système de connexion, il est impossible d'accéder au CRUDL sans être connecté.

La base de données a été créée avec MySQL et comprend deux tables : une table "Users" et une table "Tags".

Go fait le lien entre les deux. Il gère les interactions des utilisateurs et des balises entre HTML et SQL. Go sert également à gérer le serveur en permettant de spécifier une adresse et un port, ainsi qu'à gérer les routes utilisées.

L'utilisation de packages m'a permis de mettre en place un système de connexion utilisant des cookies. J'ai exploré la possibilité d'utiliser un stockage local, mais l'utilisation de cookies m'a semblé plus simple en Go, et l'utilisation de packages provenant directement de GitHub m'a permis de disposer d'un système de connexion simple et efficace en ajoutant simplement quelques lignes de code.

Ainsi, j'ai dépassé mes attentes en ajoutant le système de connexion tout en incluant un CRUDL fonctionnel qui utilise une base de données.

### Débat

Angular vs Go

Bien que Go soit un langage simple à apprendre et laisse davantage de liberté au développeur, je préfère de loin Angular. Fort de mon expérience sur mon dernier projet avec Angular, ces projets sont plus volumineux et comportent davantage de fichiers pour faire fonctionner une application web. Malgré ce détail, Angular ne nécessite pas d'autres langages et peut être utilisé indépendamment. Go, en revanche, est beaucoup plus simple à utiliser, mais adopte une philosophie similaire à celle de "Facile à apprendre, difficile à maîtriser". J'ai pu faire fonctionner une application complète avec un seul fichier. Ce fichier permettait de faire fonctionner le serveur, les templates, les interactions avec la base de données et l'interface. Malgré cette simplicité, Go ne peut fonctionner que par lui-même. En effet, pour recharger une liste sans recharger une page entière, Go nécessite l'utilisation de JavaScript, tandis qu'Angular le permet sans recourir à d'autres langages pour gérer son interactivité.

## Conclusion

En conclusion, le projet utilisant Go a été un succès. Bien que le résultat final ne soit en aucun cas un produit final et qu'il y ait encore beaucoup de place pour amélioration, j'ai réussi à apprendre Go, à gérer l'interaction entre HTML et la base de données, à créer un CRUDL et un système de connexion utilisant des cookies.

Après ce projet, vais-je utiliser Go pour un projet personnel ? Bien que Go soit simple à utiliser et que son apprentissage m'intéresse, je me sens beaucoup plus à l'aise pour réaliser un projet web avec Angular, car je connais les possibilités d'Angular et les limitations de Go me font pencher du côté d'Angular.

## Médiagraphie

Setapp. (n.d.). How to Use Go with MySQL: A Step-by-Step Guide. Récupéré de https://setapp.com/how-to/use-go-with-mysql

w3schools.com. (n.d.). Go Tutorial. Récupéré de https://www.w3schools.com/go/

Go.dev. (n.d.). Go Documentation. Récupéré de https://go.dev/doc/

Sathish, K. (2019, March 19). Go Tutorial for Beginners #3 - Functions in Go. [Video]. YouTube. Récupéré de https://www.youtube.com/watch?v=BOxBcuvASag&list=PLve39GJ2D71yyECswi0lVaBm_gbnDRR9v&index=3

Traversy Media. (2017, May 4). GoLang Crash Course #1: Introduction and Installation. [Video]. YouTube. Récupéré de https://www.youtube.com/watch?v=446E-r0rXHI

devhints.io. (n.d.). Go by Example. Récupéré de https://devhints.io/go

Trio. (2021, February 5). What Is Golang Used For? Récupéré de https://www.trio.dev/blog/what-is-golang-used-for

Mindinventory. (2021, March 24). Pros and Cons of Programming in Golang. Récupéré de https://www.mindinventory.com/blog/pros-and-cons-programming-in-golang/

Basic Go. (n.d.). Structs Explained. Récupéré de https://yourbasic.org/golang/structs-explained/#:~:text=with%20dot%20notation.-,2%20ways%20to%20create%20and%20initialize%20a%20new%20struct,to%20the%20newly%20created%20struct.&text=You%20can%20also%20create%20and%20initialize%20a%20struct%20with%20a%20struct%20literal.&text=An%20element%20list%20that%20contains,element%20for%20each%20struct%20field.

Golang By Example. (n.d.). SHA256 Hashes. Récupéré de https://gobyexample.com/sha256-hashes

Golang By Example. (n.d.). How to Set a Cookie in HTTP Using Golang. Récupéré de https://golangbyexample.com/set-cookie-http-golang/

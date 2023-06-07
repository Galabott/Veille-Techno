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

    2.2. [Historique](#Historique)

    2.3. [Méthodologie](#Méthodologie)

    2.4. [Résultats](#Résultats)

    2.5. [Débat](#Débat)

3. [Conclusion](#Conclusion)
4. [Médiagraphie](#Médiagraphie)

## Introduction

En tant que programmeur, nous savons que la technologie est quelque chose qui évolue constamment. En effet, nous en sommes un des premiers groupes touchés par cette révolution, principalement car elle dicte nos outils, et nos futurs projets.

Il est de nôtre devoir d'être capable d'évoluer dans ce domaine, et d'être capable de s'approprier de nouvelles technologies. Malgré que nos cours nous parlent d'un grand nombre de langages de programmations recherchés de nos jours, ils ne peuvent pas tout nous montrer.
Le web, étant une matière très populaire en 2023, est une technologie qui n'est pas exclu de notre situation. Le GO est un langage qui est simple à apprendre, et il sera le sujet de cette recherche.

Pendant ces deux prochaines semaines, j'essaierai de m'apprivoiser sur le GO en essayant de créer un CRUDL complet, avec un système de connexion, et le tout relié à une base de données avec MySql, qui agira aussi comme serveur pour le site web. 

## Développement

### Besoins

La recherche portera sur le GO, et elle aura comme simple but de créer un CRUDL simple avec un système de connexion, une liaison avec une base de données MySql, puis une interface à l'aide d'une utilisation simple de HTML, tout simplement pour démontrer les principes et ne pas avoir une interface console.

Ce projet me permettera d'expérimenter un langage qui n'a pas été vu en classe, et d'en faire un projet CRUDL, ce qui facilitera la comparaison avec le GO et me permettera de me faire une opinion sur le sujet.

### Historique

Le GO est un langage de programmation développé par Google. Créé en 2007, ce n'est qu'en 2009 qu'il est devenu publique, et ce n'est qu'en 2012 qu'il a commencé à prendre de la popularité depuis le lancement de la version 1 dans la même année. Malgré qu'il est récent et vise à réparer les failles d'anciens langages pour en créer un meilleur, il est fortement inspiré du Pascal et du C.

### Méthodologie

Après mon expérience précédente sur mes anciens projets présentés en classe et la recherche qui m'a permis d'en apprendre plus sur le GO, j'ai ensuite développé une série d'étapes simples à suivre pour être en mesure de procurer un CRUDL simple.

1. J'ai débuté par faire ma recherche. Une recherche sur les possiblités du GO, ses origines, et ses utilités.
2. J'ai continué en recherchant plusieurs tutoriels, de la documentation, et autres pour m'accoutumer au GOlang.
3. J'ai débuté par faire un système de connexion simple avec un storage local dans le programme.
4. J'ai expérimenté avec des connexions en base de données, avec un "Select" et un "Insert" simples.
5. J'ai ensuite lié les deux ensembles pour créer un système de connexion qui utilise une base de données, et des cookies.
6. J'ai ensuite complété le projet en créant un CRUDL simple.

### Résultats

Au final, j'ai été capable de créer mon projet comme j'avais planifié.

L'interface en HTML est simple et très rudimentaire, mais assez pour donner une interface simple à utiliser pour démontrer le fonctionnement du projet. Elle contient un système de connexion et de création de compte, et un CRUDL simple. Le CRUDL permet d'ajouter, de modifier et de supprimer des Tags qui sont affichés dans une "liste", et les fonctions du CRUDL sont gérées par un formulaire HTML (Créer et Modifier) et par des liens (Modifier et Supprimer).

Avec le système de connexion, il est impossible d'accéder au CRUDL sans être connecté.

La base de données a été faite avec MySql, et a deux tables: une table "Users", et une table "Tags".

Le GO fait l'intermédiaire entre les deux. Il gère les interactions des utilisateurs et des tags entre l'HTML et le SQL.
Le GO sert aussi à gérer le serveur, permettant de mettre une adresse et un port, et de gérer les routes utilisées.

L'utilisation des packages m'a permis d'avoir un système de connexion qui utilise les cookies.
J'avais tenté de voir si l'utilisation du stockage local était possible, mais l'utilisation des cookies me semblait le plus simple en GO, et l'utilisation de packages provenant directement de Github m'a permis d'avoir un système de connexion qui était simple et efficace en rajoutant seulement quelques lignes simples dans le code source.

Donc, selon mes attentes, je les ai dépassées légèrement en faisant l'ajout du système de connexion, tout en incluant un CRUDL fonctionnel qui utilise une base de données.

### Débat

Angular vs GO

Malgré que le GO est un langage qui est simple à apprendre et qu'il laisse beaucoup plus d'espace au développeur, je préfère absolument l'Angular.
Basé sur l'expérience de mon dernier projet avec Angular, ces projets sont plus volumineux, et contiennt beaucoup plus de fichiers pour faire fonctionner une application web, mais malgré ce détail, l'Angular ne demande pas d'autres langages et peut être utilisé indépendamment.
Le GO, lui est beaucoup plus simple d'utilisation, mais il pousse une philosophie semblable à "Facile à apprendre, difficile à maitriser". J'ai pu faire fonctionner une application en entier avec un fichier. Il permettait de faire fonctionner le serveur, les templates, de faire les intéractions avec la base de données et l'interface. Malgré cette simplicité, le GO ne peut fonctionner que par lui même. En effet, pour faire recharger une liste sans faire recharger une page complète, le GO demande une utilisation de Javascript, alors que l'Angular le permet sans utiliser d'autres langages pour gérer son intéractivité.

## Conclusion

Pour conclure, le projet en utilisant le GO a été un succès. Malgré que le résultat final n'est en aucun cas un produit final, et qu'il y a place à beaucoup d'amélioration, j'ai été capable d'apprendre le GO, de gérer l'intéraction du HTML, de gérer l'intéraction en base de données, de lier les deux, de faire un CRUDL et un système de connexion qui utilise des cookies.

Après ce projet, utiliserais-je le GO pour un projet personnel? Malgré que le GO est simple d'utilisation, et que son apprentissage m'intéresse, je me sens beaucoup plus à l'aise de faire un projet web en Angular, car je connais les possibilités de l'Angular et les restrictions du GO me font pencher du côté de l'Angular.


## Médiagraphie


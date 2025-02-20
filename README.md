# 🚀 Groupie Tracker API Rick et Morty

## 📌 Description du Projet

Ce projet est un site web exploitant l'API Rick and Morty pour récupérer et afficher des informations sur les personnages de la série, des lieux ou bien encore des episodes. L'objectif est de fournir une interface interactive permettant la recherche, l'affichage détaillé et une gestion des favoris.

## 🔧 Fonctionnalités principales

- 📜 Affichage des personnages à partir de l'API Rick et Morty

- 🔍 Recherche des personnages, des lieux ou des épisodes

- 🚬 Filtrage des personnages

- 📄 Affichage détaillé d'un personnage spécifique

- 📝 Inscription et connexion des utilisateurs

- ⭐ Gestion des favoris

- 🎨 Navigation fluide et ergonomique

## 🛠 Installation

### 📋 Prérequis

Assurez-vous d'avoir installé :

- Golang

- Un éditeur de code (ex: VS Code)

- Un navigateur web récent

### 📥 Étapes d'installation

1. Cloner le projet :

```sh
https://github.com/Shiokaa/groupie-tracker.git
cd .\src\
```

2. Installer les dépendances backend :

```sh
go mod tidy
```

3. Lancer le serveur :

```sh
go run .\main\main.go
```

## ➡️ Lancement du projet

- Aller sur le navigateur de votre choix, puis tapez dans la barre de navigation l'url suivant : http://localhost:3000/home

## 🛣️ Routes Implémentées

- **/home** : Page d'accueil
- **/collection** : Liste des personnages
- **/register** : Page d'inscription
- **/login** : Page de connexion
- **/favorite** : Liste des favoris
- **/research** : Affiche la recherche
- **/about** : A propos pour plus de détails

## 📡 API Utilisée et Endpoints Exploités

L'API Rick and Morty permet d'obtenir des informations sur les personnages, épisodes et lieux. Voici les endpoints utilisés :

1. Liste des personnages : *GET https://rickandmortyapi.com/api/character*
2. Liste des lieux : *GET https://rickandmortyapi.com/api/location*
3. Liste des episodes : *GET https://rickandmortyapi.com/api/episode*

## 🏗 Technologies Utilisées

- Backend : Golang
- Frontend : HTML / CSS / JavaScript
- Base de données : Fichier JSON

## 🎯 Conclusion

Ce projet intègre une API externe pour offrir une expérience utilisateur immersive dans l'univers de Rick et Morty. Il combine authentification, gestion des favoris et affichage dynamique pour une interaction optimale.

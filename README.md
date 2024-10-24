# TP-Final-SQL

Notre lien Github : https://github.com/PaulDecat/TP-Final-SQL.git

# Gestion des Employés

Ce projet est une application web pour gérer les employés, les postes et les départements d'une organisation. Il est développé en Go, en utilisant SQLite pour la gestion de la base de données.

## Fonctionnalités

- **Affichage des employés** : Affiche tous les employés avec possibilité de recherche par nom ou prénom.

- **Ajout d'un employé** : Permet d'ajouter un nouvel employé avec des informations détaillées.

- **Suppression d'un employé** : Permet de supprimer un employé de la base de données.
  INFO !! : Ne soyez pas surppris que dans notre table, employeId commence par 2, c'est parcequ'on a fais des test en supprimant le premier employé.

- **Liste des postes** : Affiche tous les postes disponibles dans l'organisation.

- **Ajout de postes** : Permet d'ajouter de nouveaux postes.

- **Liste des départements** : Affiche tous les départements de l'organisation.

- **Ajout de départements** : Permet d'ajouter de nouveaux départements.

## Technologies Utilisées

- **Langage** : Go
- **Base de données** : SQLite
- **Template** : HTML avec la bibliothèque `html/template`

## Le main.go

Dans le main.go nous avons le chemain des différentes route des fonctions.

## data.db

Dans data.db nous avons tous les données de l'entreprise

## data.sql

Dans data.sql nous avons les différentes requêtes pour créer nos tables

## Le server.go

Dans le server.go nous avons toutes les fonctions pour faire fonctionner notre application

## Dossier templates

Dossier contenant les fichiers HTML pour l'interface utilisateur.

## Dossier static

Dossier pour les fichiers CSS et autres ressources statiques.

## Installation

1. Clonez le dépôt :

   git clone https://github.com/votre-utilisateur/nom-du-repo.git
   cd nom-du-repo

2. Installez les dépendances nécessaires :

go get github.com/mattn/go-sqlite3

## Lancez le projet

go run server.go main.go

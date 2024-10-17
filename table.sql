CREATE TABLE employes (
    employeId INTEGER PRIMARY KEY NOT NULL,
    nom VARCHAR(255) NOT NULL,
    prenom VARCHAR(255) NOT NULL,
    sexe VARCHAR(255) NOT NULL,
    dateDeNaissance DATE NOT NULL,
    posteId INTEGER NOT NULL,
    telephone VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    superieur INTEGER NOT NULL,
    salaire INTEGER NOT NULL,
    FOREIGN KEY (posteId) REFERENCES postes(posteId),
    FOREIGN KEY (superieur) REFERENCES employes(employeId)
);

CREATE TABLE postes (
    posteId INTEGER PRIMARY KEY NOT NULL,
    nomPoste VARCHAR(255) NOT NULL,
    departementId INTEGER NOT NULL,
    FOREIGN KEY (departementId) REFERENCES departements(departementId),
);

CREATE TABLE departements (
    departementId INTEGER PRIMARY KEY NOT NULL,
    nomDepartement VARCHAR(255) NOT NULL,
    directeurDuDepartement VARCHAR(255) NOT NULL, 
    FOREIGN KEY (directeurDuDepartement) REFERENCES employes(employeId)
    
)


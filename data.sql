
CREATE TABLE IF NOT EXISTS departements (
    departementId INTEGER PRIMARY KEY NOT NULL,
    nomDepartement VARCHAR(255) NOT NULL,
    directeurDuDepartement VARCHAR(255) NOT NULL, 
    FOREIGN KEY (directeurDuDepartement) REFERENCES employes(employeId)
);


CREATE TABLE IF NOT EXISTS postes (
    posteId INTEGER PRIMARY KEY,
    nom VARCHAR(150) NOT NULL,
    departementId INTEGER,
    FOREIGN KEY (departementId) REFERENCES departements(departementId)
);


CREATE TABLE IF NOT EXISTS employes (
    employeId INTEGER PRIMARY KEY,
    nom VARCHAR(150) NOT NULL,
    prenom VARCHAR(150) NOT NULL,
    sexe VARCHAR(1) NOT NULL,
    dateDeNaissance DATE NOT NULL,
    posteId INTEGER,
    telephone VARCHAR(10) NOT NULL,
    email VARCHAR(150) NOT NULL,
    superieur INTEGER NOT NULL,
    salaire INTEGER NOT NULL,
    FOREIGN KEY (posteId) REFERENCES postes(posteId)
);



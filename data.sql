
CREATE TABLE IF NOT EXISTS departements (
    departementId INTEGER PRIMARY KEY,
    nom VARCHAR(150) NOT NULL
);


CREATE TABLE IF NOT EXISTS postes (
    posteId INTEGER PRIMARY KEY,
    nom VARCHAR(150) NOT NULL,
    departementId INTEGER,
    FOREIGN KEY (departementId) REFERENCES departements(departementId)
);


CREATE TABLE IF NOT EXISTS employes (
    employeId INTEGER PRIMARY KEY,
    nom VARCHAR(150)T NOT NULL,
    prenom VARCHAR(150) NOT NULL,
    sexe VARCHAR(1) NOT NULL,
    dateDeNaissance DATE NOT NULL,
    dateEmbauche DATE NOT NULL,
    posteId INTEGER,
    telephone VARCHAR(10) NOT NULL,
    email VARCHAR(150)EXT NOT NULL,
    superieur INTEGER,
    salaire INTEGER NOT NULL,
    FOREIGN KEY (posteId) REFERENCES postes(posteId)
);



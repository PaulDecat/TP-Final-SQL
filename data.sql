
CREATE TABLE IF NOT EXISTS departements (
    departementId INTEGER PRIMARY KEY,
    nom TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS postes (
    posteId INTEGER PRIMARY KEY,
    nom TEXT NOT NULL,
    departementId INTEGER,
    FOREIGN KEY (departementId) REFERENCES departements(departementId)
);


CREATE TABLE IF NOT EXISTS employes (
    employeId INTEGER PRIMARY KEY,
    nom TEXT NOT NULL,
    prenom TEXT NOT NULL,
    sexe TEXT NOT NULL,
    dateDeNaissance DATE NOT NULL,
    dateEmbauche DATE NOT NULL,
    posteId INTEGER,
    telephone TEXT NOT NULL,
    email TEXT NOT NULL,
    superieur INTEGER,
    salaire REAL NOT NULL,
    FOREIGN KEY (posteId) REFERENCES postes(posteId)
);



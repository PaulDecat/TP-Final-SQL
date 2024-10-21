INSERT INTO employes (employeId, nom, prenom, sexe, dateDeNaissance, posteId, telephone, email, superieur, salaire) VALUES
(1, 'Dupont', 'Jean', 'M', '1975-05-12', 1, '0102030405', 'jean.dupont@example.com', 1, 7000),
(2, 'Martin', 'Alice', 'F', '1988-02-23', 4, '0602030405', 'alice.martin@example.com', 1, 2600),
(3, 'Lefevre', 'Paul', 'M', '1985-11-13', 7, '0601020304', 'paul.lefevre@example.com', 1, 4200),
(4, 'Durand', 'Sophie', 'F', '1990-07-18', 2, '0603040506', 'sophie.durand@example.com', 1, 2200),
(5, 'Bernard', 'Marc', 'M', '1992-09-22', 5, '0604050607', 'marc.bernard@example.com', 2, 2800),
(6, 'Petit', 'Marie', 'F', '1983-10-03', 3, '0605060708', 'marie.petit@example.com', 1, 2500),
(7, 'Moreau', 'Pierre', 'M', '1995-03-15', 6, '0606070809', 'pierre.moreau@example.com', 2, 5000),
(8, 'Girard', 'Luc', 'M', '1988-12-30', 13, '0701020304', 'luc.girard@example.com', 5, 3700),
(9, 'Roux', 'Laura', 'F', '1991-06-12', 14, '0702030405', 'laura.roux@example.com', 5, 2300),
(10, 'Lemoine', 'Antoine', 'M', '1982-04-20', 8, '0703040506', 'antoine.lemoine@example.com', 3, 3900),
(11, 'Faure', 'Isabelle', 'F', '1994-08-29', 10, '0704050607', 'isabelle.faure@example.com', 4, 3100),
(12, 'Blanc', 'Lucie', 'F', '1996-07-08', 11, '0705060708', 'lucie.blanc@example.com', 4, 4000),
(13, 'Gauthier', 'Hugo', 'M', '1986-11-19', 15, '0706070809', 'hugo.gauthier@example.com', 5, 1800),
(14, 'Masson', 'Julien', 'M', '1993-02-05', 9, '0801020304', 'julien.masson@example.com', 3, 2000),
(15, 'Caron', 'Clara', 'F', '1998-09-11', 12, '0802030405', 'clara.caron@example.com', 4, 2400),
(16, 'Leclerc', 'Camille', 'F', '1984-01-15', 9, '0803040506', 'camille.leclerc@example.com', 3, 1800),
(17, 'Simon', 'Juliette', 'F', '1990-05-03', 7, '0804050607', 'juliette.simon@example.com', 3, 2300),
(18, 'Perrin', 'Nicolas', 'M', '1992-04-25', 6, '0805060708', 'nicolas.perrin@example.com', 2, 5000),
(19, 'Andre', 'Caroline', 'F', '1989-12-10', 12, '0806070809', 'caroline.andre@example.com', 4, 2400),
(20, 'Rolland', 'Thomas', 'M', '1994-08-02', 14, '0901020304', 'thomas.rolland@example.com', 5, 3700);



INSERT INTO postes (posteId, nomPoste, departementId) VALUES
(1, 'Directeur des Ressources Humaines', 1),
(2, 'Chargé de recrutement', 1),
(3, 'Gestionnaire paie', 1),
(4, 'Développeur Full Stack', 2),
(5, 'Administrateur Réseau', 2),
(6, 'Ingénieur Sécurité', 2),
(7, 'Comptable', 3),
(8, 'Chef Comptable', 3),
(9, 'Auditeur Financier', 3),
(10, 'Responsable Marketing', 4),
(11, 'Chef de Projet Marketing', 4),
(12, 'Analyste de marché', 4),
(13, 'Responsable Logistique', 5),
(14, 'Technicien Logistique', 5),
(15, 'Gestionnaire des stocks', 5);


INSERT INTO departements (departementId, nomDepartement, directeurDuDepartement) VALUES
(1, 'Ressources Humaines', 1),
(2, 'Informatique', 2),
(3, 'Comptabilité', 3),
(4, 'Marketing', 4),
(5, 'Logistique', 5);

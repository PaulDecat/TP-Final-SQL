const express = require('express'); 
const bodyParser = require('body-parser');
const sqlite3 = require('sqlite3').verbose();
const path = require('path');

const app = express();
const PORT = 3000;

// Middleware pour parser les données du formulaire
app.use(bodyParser.urlencoded({ extended: true }));

let db = new sqlite3.Database('./entreprise.db', (err) => {
  if (err) {
    console.error(err.message);
    return;
  }
  console.log('Connecté à la base de données SQLite.');
});

db.run(`CREATE TABLE IF NOT EXISTS employes (
   employeId INTEGER PRIMARY KEY NOT NULL,
   nom VARCHAR(255) NOT NULL,
   prenom VARCHAR(255) NOT NULL,
   sexe VARCHAR(1) NOT NULL,
   dateDeNaissance DATE NOT NULL,
   posteId INTEGER NOT NULL,
   telephone VARCHAR(10) NOT NULL,
   email VARCHAR(255) NOT NULL,
   superieur INTEGER NOT NULL,
   salaire INTEGER NOT NULL,
   FOREIGN KEY (posteId) REFERENCES postes(posteId),
   FOREIGN KEY (superieur) REFERENCES employes(employeId)
)`);

// Route pour afficher le formulaire
app.get('/', (req, res) => {
  res.sendFile(path.join(__dirname, 'formulaire.html'));
});

// Route pour gérer la soumission du formulaire
app.post('/ajouter-employe', (req, res) => {
  const { nom, prenom, sexe, dateDeNaissance, telephone, email, salaire, posteId, superieur } = req.body;

  // Insérer les données dans la base de données
  db.run(`INSERT INTO employes(nom, prenom, sexe, dateDeNaissance, telephone, email, salaire, posteId, superieur) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`, [nom, prenom, sexe, dateDeNaissance, telephone, email, salaire, posteId, superieur], function(err) {
    if (err) {
      console.error(err.message);
      res.send('Erreur lors de l\'ajout de l\'employé.');
      return;
    }
    res.send('Employé ajouté avec succès.');
  });
});

// Démarrer le serveur
app.listen(PORT, () => {
  console.log(`Serveur en cours d'exécution sur http://localhost:${PORT}`);
});

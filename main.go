package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./entreprise.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connexion à la base de données réussie!")
}

func getEmployes() ([]Employe, error) {
	rows, err := db.Query("SELECT employeId, nom, prenom, sexe, telephone, email FROM employes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employes []Employe
	for rows.Next() {
		var e Employe
		if err := rows.Scan(&e.EmployeId, &e.Nom, &e.Prenom, &e.Sexe, &e.Telephone, &e.Email); err != nil {
			return nil, err
		}
		employes = append(employes, e)
	}
	return employes, nil
}

type Employe struct {
	EmployeId int
	Nom       string
	Prenom    string
	Sexe      string
	Telephone string
	Email     string
}

func getPostes() ([]Poste, error) {
	rows, err := db.Query("SELECT posteId, nomPoste FROM postes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postes []Poste
	for rows.Next() {
		var p Poste
		if err := rows.Scan(&p.PosteId, &p.NomPoste); err != nil {
			return nil, err
		}
		postes = append(postes, p)
	}
	return postes, nil
}

type Poste struct {
	PosteId  int
	NomPoste string
}

func getDepartements() ([]Departement, error) {
	rows, err := db.Query("SELECT departementId, nomDepartement FROM departements")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departements []Departement
	for rows.Next() {
		var d Departement
		if err := rows.Scan(&d.DepartementId, &d.NomDepartement); err != nil {
			return nil, err
		}
		departements = append(departements, d)
	}
	return departements, nil
}

type Departement struct {
	DepartementId  int
	NomDepartement string
}

func employesHandler(w http.ResponseWriter, r *http.Request) {
	employes, err := getEmployes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/showemployes.html"))
	err = tmpl.Execute(w, employes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postesHandler(w http.ResponseWriter, r *http.Request) {
	postes, err := getPostes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/showpostes.html"))
	err = tmpl.Execute(w, postes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func departementsHandler(w http.ResponseWriter, r *http.Request) {
	departements, err := getDepartements()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/showdepartements.html"))
	err = tmpl.Execute(w, departements)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	initDB()

	http.HandleFunc("/employes", employesHandler)
	http.HandleFunc("/postes", postesHandler)
	http.HandleFunc("/departements", departementsHandler)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

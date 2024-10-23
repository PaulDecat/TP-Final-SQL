package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "html/template"
    "io/ioutil"
    "log"
    "net/http"
    "time"

)

var db *sql.DB


type Employee struct {
    EmployeId      int             // ID de l'employé
    Nom            string          // Nom de l'employé
    Prenom         string          // Prénom de l'employé
    Sexe           string          // Sexe de l'employé
    DateDeNaissance  time.Time        // Date de naissance
    PosteId        int             // ID du poste
    Telephone      string          // Numéro de téléphone
    Email          string          // Email
    Superieur      int
    Salaire        int             // Salaire
}

type Poste struct {
    PosteId      int    
    NomPoste     string 
    DepartementId int   
}

type Departement struct {
    DepartementId      int    
    NomDepartement     string 
    DirecteurDuDepartement string 
}

func initDB() {
    var err error
    db, err = sql.Open("sqlite3", "./data.db")
    if err != nil {
        log.Fatal(err)
    }

    sqlStmt, err := ioutil.ReadFile("data.sql")
    if err != nil {
        log.Fatal(err)
    }
    _, err = db.Exec(string(sqlStmt))
    if err != nil {
        log.Fatal(err)
    }
}

func closeDB() {
    db.Close()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    
    employees, err := getAllEmployees()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    
    postes, err := getAllPostes()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

   
    departements, err := getAllDepartements()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

 
    data := struct {
        Employees    []Employee
        Postes       []Poste
        Departements []Departement
    }{
        Employees:    employees,
        Postes:       postes,
        Departements: departements,
    }

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, data)
}

func getAllEmployees() ([]Employee, error) {
    rows, err := db.Query("SELECT employeId, nom, prenom, sexe, dateDeNaissance, posteId, telephone, email, superieur, salaire FROM employes")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var employees []Employee
    for rows.Next() {
        var e Employee
        err = rows.Scan(&e.EmployeId, &e.Nom, &e.Prenom, &e.Sexe, &e.DateDeNaissance, &e.PosteId, &e.Telephone, &e.Email, &e.Superieur, &e.Salaire)
        if err != nil {
            return nil, err
        }
        employees = append(employees, e)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return employees, nil
}




func getAllPostes() ([]Poste, error) {
    rows, err := db.Query("SELECT * FROM postes")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var postes []Poste
    for rows.Next() {
        var p Poste
        err = rows.Scan(&p.PosteId, &p.NomPoste, &p.DepartementId)
        if err != nil {
            return nil, err
        }
        postes = append(postes, p)
    }
    return postes, nil
}


func getAllDepartements() ([]Departement, error) {
    rows, err := db.Query("SELECT * FROM departements")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var departements []Departement
    for rows.Next() {
        var d Departement
        err = rows.Scan(&d.DepartementId, &d.NomDepartement, &d.DirecteurDuDepartement)
        if err != nil {
            return nil, err
        }
        departements = append(departements, d)
    }
    return departements, nil
}


func ajouterEmployeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        
        nom := r.FormValue("nom")
        prenom := r.FormValue("prenom")
        sexe := r.FormValue("sexe")
        dateDeNaissance := r.FormValue("dateDeNaissance")
        posteId := r.FormValue("posteId")
        telephone := r.FormValue("telephone")
        email := r.FormValue("email")
        superieur := r.FormValue("superieur")
        salaire := r.FormValue("salaire")

       
        _, err := db.Exec("INSERT INTO employes (nom, prenom, sexe, dateDeNaissance, posteId, telephone, email, superieur, salaire) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
            nom, prenom, sexe, dateDeNaissance, posteId, telephone, email, superieur, salaire)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    tmpl, err := template.ParseFiles("templates/ajouter_employe.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}




func supprimerEmployeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        employeId := r.FormValue("employeId")
        _, err := db.Exec("DELETE FROM employes WHERE employeId = ?", employeId)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    tmpl, err := template.ParseFiles("templates/supprimer_employe.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

func listerPostesHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT * FROM postes")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var postes []Poste

    for rows.Next() {
        var p Poste
        err = rows.Scan(&p.PosteId, &p.NomPoste, &p.DepartementId)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        postes = append(postes, p)
    }

    tmpl, err := template.ParseFiles("templates/lister_postes.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, postes)
}

func listerDepartementsHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT * FROM departements")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var departements []Departement

    for rows.Next() {
        var d Departement
        err = rows.Scan(&d.DepartementId, &d.NomDepartement, &d.DirecteurDuDepartement)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        departements = append(departements, d)
    }

    tmpl, err := template.ParseFiles("templates/lister_departements.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, departements)
}

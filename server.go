package main

import (
    "database/sql"
    "html/template"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "time"

    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Employee struct {
    EmployeId       int
    Nom             string
    Prenom          string
    Sexe            string
    DateDeNaissance time.Time
    PosteId         int
    Telephone       string
    Email           string
    Superieur       *int
    Salaire         int
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
    searchTerm := r.URL.Query().Get("search")
    var employees []Employee
    var err error

    if searchTerm != "" {
        employees, err = searchEmployees(searchTerm)
    } else {
        employees, err = getAllEmployees()
    }

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

func searchEmployees(searchTerm string) ([]Employee, error) {
    rows, err := db.Query("SELECT * FROM employes WHERE nom LIKE ? OR prenom LIKE ? OR employeId LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%")

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var employees []Employee
    for rows.Next() {
        var emp Employee
        err = rows.Scan(&emp.EmployeId, &emp.Nom, &emp.Prenom, &emp.Sexe, &emp.DateDeNaissance, &emp.PosteId, &emp.Telephone, &emp.Email, &emp.Superieur, &emp.Salaire)
        if err != nil {
            return nil, err
        }
        employees = append(employees, emp)
    }
    return employees, nil
}

func getAllEmployees() ([]Employee, error) {
    rows, err := db.Query("SELECT * FROM employes")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var employees []Employee
    for rows.Next() {
        var emp Employee
        err = rows.Scan(&emp.EmployeId, &emp.Nom, &emp.Prenom, &emp.Sexe, &emp.DateDeNaissance, &emp.PosteId, &emp.Telephone, &emp.Email, &emp.Superieur, &emp.Salaire)
        if err != nil {
            return nil, err
        }
        employees = append(employees, emp)
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
        dateDeNaissanceStr := r.FormValue("dateDeNaissance")

        dateDeNaissance, err := time.Parse("2006-01-02", dateDeNaissanceStr)
        if err != nil {
            http.Error(w, "Invalid date format", http.StatusBadRequest)
            return
        }

        posteId, err := strconv.Atoi(r.FormValue("posteId"))
        if err != nil {
            http.Error(w, "Invalid posteId", http.StatusBadRequest)
            return
        }

        telephone := r.FormValue("telephone")
        email := r.FormValue("email")

        var superieur *int
        if superieurStr := r.FormValue("superieur"); superieurStr != "" {
            sup, err := strconv.Atoi(superieurStr)
            if err != nil {
                http.Error(w, "Invalid superieur", http.StatusBadRequest)
                return
            }
            superieur = &sup
        }

        salaire, err := strconv.Atoi(r.FormValue("salaire"))
        if err != nil {
            http.Error(w, "Invalid salaire", http.StatusBadRequest)
            return
        }

        _, err = db.Exec("INSERT INTO employes (nom, prenom, sexe, dateDeNaissance, posteId, telephone, email, superieur, salaire) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
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

func listerEmployesHandler(w http.ResponseWriter, r *http.Request) {
    searchTerm := r.URL.Query().Get("search")
    var rows *sql.Rows
    var err error

    if searchTerm != "" {
        rows, err = db.Query("SELECT * FROM employes WHERE nom LIKE ? OR prenom LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%")
    } else {
        rows, err = db.Query("SELECT * FROM employes")
    }

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var employees []Employee
    for rows.Next() {
        var emp Employee
        err = rows.Scan(&emp.EmployeId, &emp.Nom, &emp.Prenom, &emp.Sexe, &emp.DateDeNaissance, &emp.PosteId, &emp.Telephone, &emp.Email, &emp.Superieur, &emp.Salaire)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        employees = append(employees, emp)
    }

    data := struct {
        Employees []Employee
    }{
        Employees: employees,
    }

    tmpl, err := template.ParseFiles("templates/lister_employes.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, data)
}

func ajouterDepartementHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        nomDepartement := r.FormValue("nomDepartement")
        directeurDuDepartement := r.FormValue("directeurDuDepartement")

        _, err := db.Exec("INSERT INTO departements (nomDepartement, directeurDuDepartement) VALUES (?, ?)",
            nomDepartement, directeurDuDepartement)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        
        http.Redirect(w, r, "/departements", http.StatusSeeOther)
        return
    }

    tmpl, err := template.ParseFiles("templates/ajouter_departement.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

func ajouterPosteHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        nom := r.FormValue("nom")
        departementId, err := strconv.Atoi(r.FormValue("departementId"))
        if err != nil {
            http.Error(w, "Invalid departementId", http.StatusBadRequest)
            return
        }

        _, err = db.Exec("INSERT INTO postes (nom, departementId) VALUES (?, ?)",
            nom, departementId)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/postes", http.StatusSeeOther)
        return
    }

    tmpl, err := template.ParseFiles("templates/ajouter_poste.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

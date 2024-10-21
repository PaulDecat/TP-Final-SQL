package main

import (
    "log"
    "net/http"
)

func main() {
 
    initDB()
    defer closeDB()

  
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/ajouter", ajouterEmployeHandler)
    http.HandleFunc("/modifier", modifierEmployeHandler)
    http.HandleFunc("/supprimer", supprimerEmployeHandler)
    http.HandleFunc("/postes", listerPostesHandler)
    http.HandleFunc("/departements", listerDepartementsHandler)

    
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Println("Serveur démarré sur le port 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
    }
}

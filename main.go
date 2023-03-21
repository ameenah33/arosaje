package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Struct pour stocker les données de Categorie
type Categorie struct {
	NumeroCategorie    int    `json:"NumeroCategorie"`
	TarifNormal       int    `json:"TarifNormal"`
	TarifSpecial  	   int    `json:"TarifSpecial"`
	TarifSpe  	       int    `json:"TarifSpe"`
	Classe			   string	  `json:"Classe"`
}

// Struct pour stocker les données de Client
type Client struct {
	Nom  	string    `json:"nom"`
	Prenom  string    `json:"prenom"`
	DDN		time.Time `json:"ddn"`
	Telephone string  `json:"telephone"`
}

// Struct pour stocker les données de Facture
type Facture struct {
	DateFacture int    `json:"dateFacture"`
	Montant     string `json:"montant"`
}


// Struct pour stocker les données de Reservation
type Reservation struct {
	NumeroReservation int       `json:"numeroReservation"`
	DateEntree        time.Time `json:"date_entree"`
	DateSortie        time.Time `json:"date_sortie"`
	DateReservation   time.Time `json:"date_reservation"`
	Nuite             int       `json:"nuite"`
	ClientIDClient    int       `json:"Client_idClient"`
	FactureIDFacture  int       `json:"Facture_idFacture"`
}

// Struct pour stocker les données de Service
type Service struct {
	PetitDejeuner    int    `json:"petit_dejeuner"`
	Phone		   	  int    `json:"phone"`
	Bar    			  int    `json:"bar"`
	TarifDejeuner                int `json:"tarif_dejeuner"`
	TarifPhone                   int `json:"tarif_phone"`
	TarifBar                     int `json:"tarif_bar"`
	ReservationNumeroReservation int `json:"Reservation_numeroReservation"`
}

// Struct pour stocker les données de Niveau
type Niveau struct {
	NumeroNiveau  int `json:"numeroNiveau"`
	NbChambres    int `json:"nb_chambres"`
	HotelIDHotel  int `json:"Hotel_idHotel"`
}

// Struct pour stocker les données de Chambre
type Chambre struct {
	NumeroChambre              int    `json:"numeroChambre"`
	Etat                       string `json:"etat"`
	NiveauNumeroNiveau         int    `json:"Niveau_numeroNiveau"`
	ReservationNumeroReservation int  `json:"Reservation_numeroReservation"`
	CategorieNumeroCategorie   int    `json:"Categorie_numeroCategorie"`
}

// Struct pour stocker les données de Hotel
type Hotel struct {
	NomHotel    string    `json:"NomHotel"`
	NbNiveaux	   int    `json:"NbNiveaux"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "12345678"
	dbName := "gestionHotel"
	dbHost := "localhost" // or your ProxySQL address
	dbPort := "3306"

	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{

		                // Catégories routes

				v1.POST("/ajout-categorie", createCategorie)
				v1.GET("/liste-categories", getAllCategories)
				v1.GET("/categorie-details/:id", getCategorieByID)
				v1.PUT("/modifier-categorie/:id", updateCategorie)
				v1.DELETE("/supprimer-categorie/:id", deleteCategorie)
				
				// Hôtels routes
		
				v1.POST("/ajout-hotel", createHotel)
				v1.GET("/liste-hotels", getAllHotels)
				v1.GET("/hotel-details/:id", getHotelByID)
				v1.PUT("/modifier-hotel/:id", updateHotel)
				v1.DELETE("/supprimer-hotel/:id", deleteHotel)
				
				// Services routes
		
				v1.POST("/ajout-service", createService)
				v1.GET("/liste-services", getAllServices)
				v1.GET("/service-details/:id", getServiceByID)
				v1.PUT("/modifier-service/:id", updateService)
				v1.DELETE("/supprimer-service/:id", deleteService)
		
				// Clients routes
		
				v1.POST("/ajout-client", createClient)
				v1.GET("/liste-clients", getAllClients)
				v1.GET("/client-details/:id", getClientByID)
				v1.PUT("/modifier-client/:id", updateClient)
				v1.DELETE("/supprimer-client/:id", deleteClient)
		
				// Factures routes
		
				v1.POST("/ajout-facture", createFacture)
				v1.GET("/liste-factures", getAllFactures)
				v1.GET("/facture-details/:id", getFactureByID)
				v1.PUT("/modifier-facture/:id", updateFacture)
				v1.DELETE("/supprimer-facture/:id", deleteFacture)
		
				// Réservations routes
		
				v1.POST("/ajout-reservation", createReservation)
				v1.GET("/liste-reservations", getAllReservations)
				v1.GET("/reservation-details/:id", getReservationByID)
				v1.PUT("/modifier-reservation/:id", updateReservation)
				v1.DELETE("/supprimer-reservation/:id", deleteReservation)

				// Chambres routes
		
				v1.POST("/ajout-chambre", createChambre)
				v1.GET("/liste-chambres", getAllChambres)
				v1.GET("/chambre-details/:id", getChambreByID)
				v1.PUT("/modifier-chambre/:id", updateChambre)
				v1.DELETE("/supprimer-chambre/:id", deleteChambre)
		
				// Niveaux routes
		
				v1.POST("/ajout-niveau", createNiveau)
				v1.GET("/liste-niveaux", getAllNiveaux)
				v1.GET("/niveau-details/:id", getNiveauByID)
				v1.PUT("/modifier-niveau/:id", updateNiveau)
				v1.DELETE("/supprimer-niveau/:id", deleteNiveau)
		
			
		
	}

	router.Run(":8080")
}
// Fonction Categories 
func getAllCategories(c *gin.Context) {
	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT numeroCategorie, tarif_normal, tarif_special, tarif_spe, classe FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	categories := make([]Categorie, 0)
	for rows.Next() {
		var categorie Categorie
		err := rows.Scan(&categorie.NumeroCategorie, &categorie.TarifNormal, &categorie.TarifSpecial, &categorie.TarifSpe, &categorie.Classe)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categories = append(categories, categorie)
	}

	c.JSON(http.StatusOK, categories)
}

func createCategorie(c *gin.Context) {
	var categorie Categorie
	err := json.NewDecoder(c.Request.Body).Decode(&categorie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO categories(tarif_normal, tarif_special, tarif_spe, classe) VALUES(?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(categorie.TarifNormal, categorie.TarifSpecial, categorie.TarifSpe, categorie.Classe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Catégorie créée avec succès"})
}

func getCategorieByID(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	row := db.QueryRow("SELECT numeroCategorie, tarif_normal, tarif_special, tarif_spe, classe FROM categories WHERE numeroCategorie = ?", id)

	var categorie Categorie
	err := row.Scan(&categorie.NumeroCategorie, &categorie.TarifNormal, &categorie.TarifSpecial, &categorie.TarifSpe, &categorie.Classe)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Catégorie introuvable"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categorie)
}

func updateCategorie(c *gin.Context) {
	id := c.Param("id")

	var categorie Categorie
	err := json.NewDecoder(c.Request.Body).Decode(&categorie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE categories SET tarif_normal = ?, tarif_special = ?, tarif_spe = ?, classe = ? WHERE numeroCategorie = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(categorie.TarifNormal, categorie.TarifSpecial, categorie.TarifSpe, categorie.Classe, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Catégorie mise à jour avec succès"})
}

func deleteCategorie(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM categories WHERE numeroCategorie = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Catégorie supprimée avec succès"})
}

// Fonction Hotel

func createHotel(c *gin.Context) {
	var hotel Hotel
	err := json.NewDecoder(c.Request.Body).Decode(&hotel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Hotel(nomHotel, nb_niveaux) VALUES(?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(hotel.NomHotel, hotel.NbNiveaux)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Hôtel créé avec succès"})
}

func getAllHotels(c *gin.Context) {
	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT id, nomHotel, nb_niveaux FROM Hotel")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	hotels := make([]Hotel, 0)
	for rows.Next() {
		var hotel Hotel
		var id int
		err := rows.Scan(&id, &hotel.NomHotel, &hotel.NbNiveaux)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		hotels = append(hotels, hotel)
	}

	c.JSON(http.StatusOK, hotels)
}

func getHotelByID(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	row := db.QueryRow("SELECT nomHotel, nb_niveaux FROM Hotel WHERE id = ?", id)

	var hotel Hotel
	err := row.Scan(&hotel.NomHotel, &hotel.NbNiveaux)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Hôtel introuvable"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotel)
}

func updateHotel(c *gin.Context) {
	id := c.Param("id")

	var hotel Hotel
	err := json.NewDecoder(c.Request.Body).Decode(&hotel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Hotel SET nomHotel = ?, nb_niveaux = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(hotel.NomHotel, hotel.NbNiveaux, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hôtel mis à jour avec succès"})
}

func deleteHotel(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Hotel WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hôtel supprimé avec succès"})
}

// Fonction service 
func createService(c *gin.Context) {
	var service Service
	err := json.NewDecoder(c.Request.Body).Decode(&service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Service(petit_dejeuner, phone, bar, tarif_dejeuner, tarif_phone, tarif_bar, Reservation_numeroReservation) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(service.PetitDejeuner, service.Phone, service.Bar, service.TarifDejeuner, service.TarifPhone, service.TarifBar, service.ReservationNumeroReservation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": " Service créé avec succès "})
}

func getAllServices(c *gin.Context) {
	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT id, petit_dejeuner, phone, bar, tarif_dejeuner, tarif_phone, tarif_bar, Reservation_numeroReservation FROM Service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	services := make([]Service, 0)
	for rows.Next() {
		var service Service
		var id int
		err := rows.Scan(&id, &service.PetitDejeuner, &service.Phone, &service.Bar, &service.TarifDejeuner, &service.TarifPhone, &service.TarifBar, &service.ReservationNumeroReservation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		services = append(services, service)
	}

	c.JSON(http.StatusOK, services)
}

func getServiceByID(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	row := db.QueryRow("SELECT petit_dejeuner, phone, bar, tarif_dejeuner, tarif_phone, tarif_bar, Reservation_numeroReservation FROM Service WHERE id = ?", id)

	var service Service
	err := row.Scan(&service.PetitDejeuner, &service.Phone, &service.Bar, &service.TarifDejeuner, &service.TarifPhone, &service.TarifBar, &service.ReservationNumeroReservation)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Service non trouvé"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, service)
}

func updateService(c *gin.Context) {
	id := c.Param("id")

	var service Service
	err := json.NewDecoder(c.Request.Body).Decode(&service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Service SET petit_dejeuner = ?, phone = ?, bar = ?, tarif_dejeuner = ?, tarif_phone = ?, tarif_bar = ?, Reservation_numeroReservation = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(service.PetitDejeuner, service.Phone, service.Bar, service.TarifDejeuner, service.TarifPhone, service.TarifBar, service.ReservationNumeroReservation, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service mis à jour avec succès"})
}

func deleteService(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Service WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service supprimé avec succès"})
}

// fonction Client

func createClient(c *gin.Context) {
	var client Client
	err := json.NewDecoder(c.Request.Body).Decode(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Client(nom, prenom, ddn, telephone) VALUES(?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(client.Nom, client.Prenom, client.DDN, client.Telephone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Client créé avec succès"})
}

func getAllClients(c *gin.Context) {
	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT id, nom, prenom, ddn, telephone FROM Client")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	clients := make([]Client, 0)
	for rows.Next() {
		var client Client
		var id int
		err := rows.Scan(&id, &client.Nom, &client.Prenom, &client.DDN, &client.Telephone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		clients = append(clients, client)
	}

	c.JSON(http.StatusOK, clients)
}

func getClientByID(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	row := db.QueryRow("SELECT nom, prenom, ddn, telephone FROM Client WHERE id = ?", id)

	var client Client
	err := row.Scan(&client.Nom, &client.Prenom, &client.DDN, &client.Telephone)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Client introuvable"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}
func updateClient(c *gin.Context) {
	id := c.Param("id")

	var client Client
	err := json.NewDecoder(c.Request.Body).Decode(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Client SET nom = ?, prenom = ?, ddn = ?, telephone = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(client.Nom, client.Prenom, client.DDN, client.Telephone, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client mis à jour avec succès"})
}

func deleteClient(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Client WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client supprimé avec succès"})
}

// Fonction  Facture
func createFacture(c *gin.Context) {
	var facture Facture
	err := json.NewDecoder(c.Request.Body).Decode(&facture)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Facture(dateFacture, montant) VALUES(?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(facture.DateFacture, facture.Montant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Facture créée avec succès"})
}

func getAllFactures(c *gin.Context) {
	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT id, dateFacture, montant FROM Facture")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	factures := make([]Facture, 0)
	for rows.Next() {
		var facture Facture
		var id int
		err := rows.Scan(&id, &facture.DateFacture, &facture.Montant)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		factures = append(factures, facture)
	}

	c.JSON(http.StatusOK, factures)
}

func getFactureByID(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	row := db.QueryRow("SELECT dateFacture, montant FROM Facture WHERE id = ?", id)

	var facture Facture
	err := row.Scan(&facture.DateFacture, &facture.Montant)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Facture introuvable"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, facture)
}

func updateFacture(c *gin.Context) {
	id := c.Param("id")

	var facture Facture
	err := json.NewDecoder(c.Request.Body).Decode(&facture)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Facture SET dateFacture = ?, montant = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(facture.DateFacture, facture.Montant, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Facture mise à jour avec succès"})
}

func deleteFacture(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Facture WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Facture supprimée avec succès"})
}

// Fonction reservation 
func createReservation(c *gin.Context) {
	var reservation Reservation
	err := json.NewDecoder(c.Request.Body).Decode(&reservation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Reservation(date_entree, date_sortie, date_reservation, nuite, Client_idClient, Facture_idFacture) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(reservation.DateEntree, reservation.DateSortie, reservation.DateReservation, reservation.Nuite, reservation.ClientIDClient, reservation.FactureIDFacture)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Réservation créée avec succès"})
}

func getAllReservations(c *gin.Context) {
	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT id, date_entree, date_sortie, date_reservation, nuite, Client_idClient, Facture_idFacture FROM Reservation")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	reservations := make([]Reservation, 0)
	for rows.Next() {
		var reservation Reservation
		var id int
		err := rows.Scan(&id, &reservation.DateEntree, &reservation.DateSortie, &reservation.DateReservation, &reservation.Nuite, &reservation.ClientIDClient, &reservation.FactureIDFacture)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		reservations = append(reservations, reservation)
	}

	c.JSON(http.StatusOK, reservations)
}

func getReservationByID(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	row := db.QueryRow("SELECT date_entree, date_sortie, date_reservation, nuite, Client_idClient, Facture_idFacture FROM Reservation WHERE id = ?", id)

	var reservation Reservation
	err := row.Scan(&reservation.DateEntree, &reservation.DateSortie, &reservation.DateReservation, &reservation.Nuite, &reservation.ClientIDClient, &reservation.FactureIDFacture)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Réservation introuvable"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservation)
}

func updateReservation(c *gin.Context) {
	id := c.Param("id")

	var reservation Reservation
	err := json.NewDecoder(c.Request.Body).Decode(&reservation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Reservation SET date_entree = ?, date_sortie = ?, date_reservation = ?, nuite = ?, Client_idClient = ?, Facture_idFacture = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(reservation.DateEntree, reservation.DateSortie, reservation.DateReservation, reservation.Nuite, reservation.ClientIDClient, reservation.FactureIDFacture, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Réservation mise à jour avec succès"})
}

func deleteReservation(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Reservation WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Réservation supprimée avec succès"})
}

// Fonction Chambre 

func createChambre(c *gin.Context) {
	var chambre Chambre
	err := json.NewDecoder(c.Request.Body).Decode(&chambre)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Chambre(numeroChambre, etat, Niveau_numeroNiveau, Reservation_numeroReservation, Categorie_numeroCategorie) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(chambre.NumeroChambre, chambre.Etat, chambre.NiveauNumeroNiveau, chambre.ReservationNumeroReservation, chambre.CategorieNumeroCategorie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "chambre créée avec succès"})
}

func getAllChambres(c *gin.Context) {
	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT id, numeroChambre, etat, Niveau_numeroNiveau, Reservation_numeroReservation, Categorie_numeroCategorie FROM Chambre")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	chambres := make([]Chambre, 0)
	for rows.Next() {
		var chambre Chambre
		var id int
		err := rows.Scan(&id, &chambre.NumeroChambre, &chambre.Etat, &chambre.NiveauNumeroNiveau, &chambre.ReservationNumeroReservation, &chambre.CategorieNumeroCategorie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		chambres = append(chambres, chambre)
	}

	c.JSON(http.StatusOK, chambres)
}


func getChambreByID(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	row := db.QueryRow("SELECT numeroChambre, etat, Niveau_numeroNiveau, Reservation_numeroReservation, Categorie_numeroCategorie FROM Chambre WHERE id = ?", id)

	var chambre Chambre
	err := row.Scan(&chambre.NumeroChambre, &chambre.Etat, &chambre.NiveauNumeroNiveau, &chambre.ReservationNumeroReservation, &chambre.CategorieNumeroCategorie)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Chambre introuvable"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chambre)
}

func updateChambre(c *gin.Context) {
	id := c.Param("id")

	var chambre Chambre
	err := json.NewDecoder(c.Request.Body).Decode(&chambre)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Chambre SET numeroChambre = ?, etat = ?, Niveau_numeroNiveau = ?, Reservation_numeroReservation = ?, Categorie_numeroCategorie = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(chambre.NumeroChambre, chambre.Etat, chambre.NiveauNumeroNiveau, chambre.ReservationNumeroReservation, chambre.CategorieNumeroCategorie, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chambre mise à jour avec succès"})
}

func deleteChambre(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Chambre WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chambre supprimée avec succès "})
}

// Fonction Niveau

func createNiveau(c *gin.Context) {
	var niveau Niveau
	err := json.NewDecoder(c.Request.Body).Decode(&niveau)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Niveau(numeroNiveau, nb_chambres, Hotel_idHotel) VALUES(?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(niveau.NumeroNiveau, niveau.NbChambres, niveau.HotelIDHotel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": " Niveau créé avec succès "})
}

func getAllNiveaux(c *gin.Context) {
	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT id, numeroNiveau, nb_chambres, Hotel_idHotel FROM Niveau")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	niveaux := make([]Niveau, 0)
	for rows.Next() {
		var niveau Niveau
		var id int
		err := rows.Scan(&id, &niveau.NumeroNiveau, &niveau.NbChambres, &niveau.HotelIDHotel)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		niveaux = append(niveaux, niveau)
	}

	c.JSON(http.StatusOK, niveaux)
}

func getNiveauByID(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	row := db.QueryRow("SELECT numeroNiveau, nb_chambres, Hotel_idHotel FROM Niveau WHERE id = ?", id)

	var niveau Niveau
	err := row.Scan(&niveau.NumeroNiveau, &niveau.NbChambres, &niveau.HotelIDHotel)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": " Niveau introuvable "})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, niveau)
}

func updateNiveau(c *gin.Context) {
	id := c.Param("id")

	var niveau Niveau
	err := json.NewDecoder(c.Request.Body).Decode(&niveau)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Niveau SET numeroNiveau = ?, nb_chambres = ?, Hotel_idHotel = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(niveau.NumeroNiveau, niveau.NbChambres, niveau.HotelIDHotel, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": " Niveau mis à jour avec succès "})
}

func deleteNiveau(c *gin.Context) {
	id := c.Param("id")

	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Niveau WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Niveau supprimé avec succès"})
}

package controllers

import (
	"math/rand"
	"net/http"

	"github.com/J-Nokwal/ihar_backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json", "code": "1000"})
		return
	}
	// if user.ID != "manu" || user.Email != "123" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 	return
	// }
	if user.IsAnoymous {
		user.FirstName = "Anonymous"
		user.LastName = getRandomUserName()
	}
	if _, err := user.CreateUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1001"}) // error: "error while insertion"
		return
	}
	c.JSON(http.StatusOK, user)
}
func GetUser(c *gin.Context) {
	id := c.Param("id")
	// var user models.User
	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
		return
	}
	c.JSON(http.StatusOK, user)
}
func PatchUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json", "code": "1000"})
		return
	}
	// if user.ID != "manu" || user.Email != "123" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 	return
	// }
	u, err := models.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1003"}) // error: "error while patching"
		return
	}
	c.JSON(http.StatusOK, u)
}

func getRandomUserName() string {
	usernames := []string{"Coeliac", "Gruntled", "Gunnage", "Typhoon", "Jacobin", "DrongoMayhem", "Laciniate", "Disemvowel", "Fissicostate", "PuppypuppFard", "Apodictic", "Gibberish", "Pachynsis", "CandleMufti", "Urinator", "Blackguard", "Oppilate", "Candelabra", "ChetchetRivel", "Caterwaul", "Cathexis", "Flyspeck", "Nigrescence", "GampodPuss", "Petcock", "LordkurDander", "Hemitery", "BitiamPogo", "Berline", "Gongoozle", "Isomorph", "Hillbilly", "RedthecZurla", "Hodgepodge", "VeldomVatic", "AcrilKazoo", "Shrouds", "Canoodle", "GtminingAstrut", "Zeitgeist", "Sinister", "Pandemonium", "Dinglewong", "Fribble", "Opinable", "Spelunker", "Corbeil", "Ambrosial", "Limousja0303", "Hullabaloo", "Dicerous", "Grudgemental", "Zoochorous", "Foxy", "Astrobolism", "Cahoots", "Thionine", "Cattitude", "Presidial", "Kerfuffle", "Conureger2226", "Kaput", "Volacious", "Slosh", "Stupereezmc", "Skullet", "Planation", "Salpiglossis", "Homophyly", "Piffling", "Vigneron", "Occiput", "Argentreef2000", "Borborygmous", "Abodement", "HoiPolloi", "Flenchheed808", "Shrubbery", "Stolon_bard", "Kahikatea"}
	randomIndex := rand.Intn(len(usernames))
	pick := usernames[randomIndex]
	return pick
}

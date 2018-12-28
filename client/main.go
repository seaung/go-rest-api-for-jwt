package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var myStringKey = os.Getenv("SECRET_KEY")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error : %s", err.Error())
	}

	defer req.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(body))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "seaung"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	toekString, err := token.SignedString(myStringKey)

	if err != nil {
		fmt.Errorf("wrong:%s", err.Error())
		return "", err
	}
	return toekString, nil
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("run client on 9001")

	handleRequests()
}

package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/ClementTeyssa/3PJT-API2/helper"
)

type Transaction struct {
	AccountFrom string  `json:"accountfrom"`
	AccountTo   string  `json:"accountto""`
	Amount      float32 `json:"amount"`
}

type Node struct {
	IpAdress string `json:"ipAdress"`
	Account  string `json:"adress("`
}

type Nodes []Node

const URL_GET_NODES = "https://3pjt-dnode.infux.fr/get-nodes"

func DoVerifications(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//TODO: récupérer la liste des nodes depuis le dnode
	response, err := http.Get(URL_GET_NODES)

	if err != nil {
		helper.ErrorHandlerHttpRespond(w, err.Error())
	}

	if response == nil {
		helper.ErrorHandlerHttpRespond(w, "Problem to connect to DNode")
	}

	// var Nodes Nodes
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "ioutil.ReadAll(response.Body)")
		return
	}

	if body == nil || string(body) == "null" {
		helper.ErrorHandlerHttpRespond(w, "0 nodes connected")
		return
	}

	// json.Unmarshal(body, &Nodes)

	// log.Println("Account : " + Nodes[0].Account)
	// log.Println("IP : " + Nodes[0].IpAdress)

	// var email, password string
	// fmt.Printf("Email: ")
	// fmt.Scan(&email)
	// fmt.Printf("Password: ")
	// fmt.Scan(&password)
	// defs.MyUser.Email = email
	// defs.MyUser.Password = password
	// jsonValue, _ := json.Marshal(defs.MyUser)
	// response, err := http.Post("https://3pjt-api.infux.fr/login", "application/json", bytes.NewBuffer(jsonValue))
	// if err != nil && response == nil {
	// 	fmt.Printf("The HTTP request failed with error %s\n", err)
	// 	return
	// // } else {
	// data, _ := ioutil.ReadAll(response.Body)
	// json.Unmarshal(data, &defs.MyNode)

	//TODO:	choisir 3 nodes aléatoires
	//TODO: envoyer les infos de transac pour qu'ils vérifient
	//TODO: si les 3 sont error on retourne
	//TODO: si les 3 sont OK on ajoute la transaction en l'envoyant à l'API1
	//TODO: on récupère l'id de la transaction qu'on donne à un des 3 node pour qu'il ajoute la transaction à la blockchain
	//TODO: une fois ajouté, le node nous renvoie les informations du block
	//TODO: on ajoute le block à l'API1
	//TODO: on envoie dans l'API3 l'identité des 3 nodes qui ont validé pour qu'il puissent avoir un reward au bout de x validations
}

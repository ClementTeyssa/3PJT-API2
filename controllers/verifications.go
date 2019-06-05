package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	"github.com/ClementTeyssa/3PJT-API2/helper"
)

type Transaction struct {
	AccountFrom string  `json:"accountfrom"`
	AccountTo   string  `json:"accountto""`
	Amount      float32 `json:"amount"`
	Private     []byte  `json:"privatekey"`
}

type NodeGet struct {
	IpAdressAndPort string `json:"ipAdress"`
	Account         string `json:"adress"`
}

type NodesGet struct {
	Nodes []NodeGet `json:"nodes"`
}

type Node struct {
	Ip     string
	Port   string
	Adress string
}

type Nodes []Node

type AdressApiKey struct {
	Adress string `json:"adress"`
	ApiKey string `json:"apikey"`
}

type Solde struct {
	Solde float32 `json:"solde"`
}

type GoodResult struct {
	Good string `json:"good"`
}

const URL_GET_NODES = "https://3pjt-dnode.infux.fr/get-nodes"
const URL_GET_SOLDE_API = "https://3pjt-api.infux.fr/soldeapi"
const URL_GET_VERIF = "https://3pjt-api.infux.fr/transactions/verify"

func DoVerifications(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// on récupère la liste des nodes depuis le dnode
	GetNodes, err := getNodes()
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, err.Error())
		return
	}

	var Nodes Nodes

	// on vérifie que les nodes ont un solde suffisant pour être sûr
	for _, GetNode := range GetNodes.Nodes {
		enoughSolde, err := verifyEnoughSolde(GetNode.Account)
		if err != nil {
			helper.ErrorHandlerHttpRespond(w, err.Error())
			return
		}

		// si il y a suffisament de solde
		if enoughSolde == true {
			vars := strings.Split(GetNode.IpAdressAndPort, "/")
			var Node Node
			Node.Ip = vars[2]
			Node.Port = vars[4]
			Node.Adress = GetNode.Account
			Nodes = append(Nodes, Node)

		}
	}

	// si il n'y a plus suffisament de node après la verif du solde
	if len(Nodes) < 3 {
		helper.ErrorHandlerHttpRespond(w, "No enought nodes (min 3)")
		return
	}

	// choisir 3 nodes aléatoires
	Nodes = choose3Nodes(Nodes)

	// var validNode []Node
	// append(validNode, Nodes[indexNode1], Nodes[indexNode2], Nodes[indexNode3])

	//TODO: envoyer les infos de transac pour qu'ils vérifient

	body2, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "ioutil.ReadAll(r.Body)")
		return
	}

	var transaction Transaction
	err = json.Unmarshal(body2, &transaction)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "json.Unmarshal(body, &transaction)")
		println(string(body2))
		return
	}

	verif1, err1 := askForVerifToNodes(Nodes[0], transaction)
	println(verif1)
	println(err1.Error())
	verif2, err2 := askForVerifToNodes(Nodes[1], transaction)
	println(verif2)
	println(err2.Error())
	verif3, err3 := askForVerifToNodes(Nodes[2], transaction)
	println(verif3)
	println(err3.Error())

	// si les un des 3 retourne une erreur -> on retourne

	if !verif1 || !verif2 || !verif3 {
		if !verif1 {
			helper.ErrorHandlerHttpRespond(w, err1.Error())
			return
		}
		if !verif2 {
			helper.ErrorHandlerHttpRespond(w, err2.Error())
			return
		}
		if !verif3 {
			helper.ErrorHandlerHttpRespond(w, err3.Error())
			return
		}
	}

	//TODO: si les 3 sont OK on ajoute la transaction en l'envoyant à l'API1

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
	// } else {
	// data, _ := ioutil.ReadAll(response.Body)
	// json.Unmarshal(data, &defs.MyNode)

	//TODO: on récupère l'id de la transaction qu'on donne à un des 3 node pour qu'il ajoute la transaction à la blockchain
	//TODO: une fois ajouté, le node nous renvoie les informations du block
	//TODO: on ajoute le block à l'API1
	//TODO: on envoie dans l'API3 l'identité des 3 nodes qui ont validé pour qu'il puissent avoir un reward au bout de x validations
}

func getNodes() (NodesGet, error) {
	response, err := http.Get(URL_GET_NODES)

	var NodesGet NodesGet

	if err != nil {
		return NodesGet, err
	}

	if response == nil {
		return NodesGet, errors.New("Problem to connect to DNode")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return NodesGet, errors.New("ioutil.ReadAll(response.Body)")
	}

	if body == nil || string(body) == "null" {
		return NodesGet, errors.New("0 nodes connected")
	}

	json.Unmarshal(body, &NodesGet)
	return NodesGet, nil
}

func choose3Nodes(pNodes Nodes) Nodes {
	indexNode1 := rand.Intn(len(pNodes))

	indexNode2 := rand.Intn(len(pNodes))
	for indexNode2 == indexNode1 {
		indexNode2 = rand.Intn(len(pNodes))
	}

	indexNode3 := rand.Intn(len(pNodes))
	for indexNode2 == indexNode1 || indexNode3 == indexNode1 {
		indexNode3 = rand.Intn(len(pNodes))
	}

	var Nodes Nodes
	Nodes = append(Nodes, pNodes[indexNode1])
	Nodes = append(Nodes, pNodes[indexNode2])
	Nodes = append(Nodes, pNodes[indexNode3])
	println(indexNode1)
	println(indexNode2)
	println(indexNode3)

	return Nodes
}

func verifyEnoughSolde(adress string) (bool, error) {
	var AdressApiKey AdressApiKey
	AdressApiKey.Adress = adress
	AdressApiKey.ApiKey = helper.ApiKey

	jsonToSend, err := json.Marshal(AdressApiKey)

	if err != nil {
		return false, err
	}
	response, err := http.Post(URL_GET_SOLDE_API, "application/json", bytes.NewBuffer(jsonToSend))
	if err != nil {
		return false, err
	}

	if response == nil {
		return false, errors.New("Response is null")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	var Error helper.MyError
	json.Unmarshal(body, &Error)
	if Error.Error != "" {
		return false, errors.New(Error.Error)
	}

	var Solde Solde
	json.Unmarshal(body, &Solde)

	// si le solde n'est pas suffisant pour être "sûr"
	if Solde.Solde < 10 {
		return false, nil
	}

	return true, nil
}

func askForVerifToNodes(node Node, transac Transaction) (bool, error) {
	jsonToSend, err := json.Marshal(transac)
	if err != nil {
		return false, err
	}
	println(node.Ip)
	println(node.Port)

	//send to a node
	response, err := http.Post("http://"+node.Ip+":"+node.Port, "application/json", bytes.NewBuffer(jsonToSend))
	if err != nil {
		return false, err
	}

	if response == nil {
		return false, errors.New("Response is null")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	var Error helper.MyError
	json.Unmarshal(body, &Error)
	if Error.Error != "" {
		return false, errors.New(Error.Error)
	}

	var Good GoodResult
	json.Unmarshal(body, &Good)
	if Good.Good != "OK" {
		return false, errors.New("Not the good result")
	}

	return true, nil
}

// func sendTransactionToApi() (Transaction, error) {

// }

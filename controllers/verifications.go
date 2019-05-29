package controllers

import "net/http"

func DoVerifications(w http.ResponseWriter, r *http.Request) {
	//TODO: récupérer la liste des nodes depuis le dnode
	//TODO:	choisir 3 nodes aléatoires
	//TODO: envoyer les infos de transac pour qu'ils vérifient
	//TODO: si les 3 sont error on retourne
	//TODO: si les 3 sont OK on ajoute la transaction en l'envoyant à l'API1
	//TODO: on récupère l'id de la transaction qu'on donne à un des 3 node pour qu'il ajoute la transaction à la blockchain
	//TODO: une fois ajouté, le node nous renvoie les informations du block
	//TODO: on ajoute le block à l'API1
	//TODO: on envoie dans l'API3 l'identité des 3 nodes qui ont validé pour qu'il puissent avoir un reward au bout de x validations
}

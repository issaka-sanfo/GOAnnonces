Issaka SANFO

# GOAnnonces

1- Lancer l'Application : 
*********************************************************************************************************

$ docker-compose up



2- Requêtes curl , https://curl.se/:
*********************************************************************************************************

Créer une Annonce, POST:

$ curl http://localhost:10000/annonces/ -d "titre=Car&contenu=Good state&categorie=Automobile&model=CrossBack ds 3"


Avoir liste Annonces, GET:

$ curl http://localhost:10000/annonces/



Modifier une Annonce, PUT:

curl -X PUT http://localhost:10000/annonces/4 -d '{"titre":"Camionette",contenu":"Second Hand","categorie":"Automobile","model":"CrossBack ds 3"}'



Récupere Une Annonce, GET:

$ curl http://localhost:10000/annonces/1



Supprimer une Annonce, DELETE:

$ curl -X DELETE http://localhost:10000/annonces/1




Unit Tests:

Les tests unitaires se trouvent dans le sous dossier "testapis"




Je suis ouvert aux commentaires et prêt à améliorer mon travail de jour en jour! Merci:)



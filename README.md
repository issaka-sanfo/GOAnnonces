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




3- Unit Tests:
*********************************************************************************************************
Les tests unitaires se trouvent dans le sous dossier "testapis"


4- AWS
*********************************************************************************************************

EC2 - AMI Instance with Ubuntu DISTRIB_DESCRIPTION="Ubuntu 20.04.3 LTS"

J'ai installé Docker sur cette Instance avec les instruction indiquées sur:
https://docs.docker.com/engine/install/ubuntu/

En suite, j'ai installé docker-compose avec la commande: 

$ sudo apt  install docker-compose

Le container est lancé au port 10000 de l'Instance, on peut faire les Tests APIs avec le DNS suivant:

ec2-3-143-220-107.us-east-2.compute.amazonaws.com:10000/annonces/





Je suis ouvert aux commentaires et prêt à améliorer mon travail de jour en jour! Merci:)



package connectors

import (
	"context"
	"log"
	"os"
	"path"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/timPrachasri/app-chat/server/app/environments"
	"google.golang.org/api/option"
)

var (
	firestoreInstance *firestore.Client
)

// ConnectFirestore is a connector function for connecting firebase's firestore
func ConnectFirestore() *firestore.Client {
	return connectToFirestore(
		environments.FirebaseFilename,
	)
}

func connectToFirestore(firebaseFilename string) *firestore.Client {
	mut.Lock()
	defer mut.Unlock()
	if firestoreInstance == nil {
		ctx := context.Background()
		pwdPath, err := os.Getwd()
		if err != nil {
			log.Println("errr from google firebase")
			log.Fatalln(err)
		}
		log.Println(pwdPath, "pwd path")
		sa := option.WithCredentialsFile(path.Join(pwdPath, "../../server", firebaseFilename+".json"))
		app, err := firebase.NewApp(ctx, nil, sa)
		if err != nil {
			log.Println("errr from google firebase")
			log.Fatalln(err)
		}
		firestoreInstance, err = app.Firestore(ctx)
		if err != nil {
			log.Println("errr from google firestore")
			log.Fatalln(err)
		}
	}
	return firestoreInstance
}

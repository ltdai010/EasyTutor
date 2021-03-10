package drivers

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/beego/beego/v2/server/web"
	"google.golang.org/api/option"
	"log"
	"sync"
)

var (
	onceDo          = &sync.Once{}
	driverInterface DriverInterface
)

type DriverFactory struct {
	teacherIndex search.IndexInterface
	requestIndex search.IndexInterface
	cloudStore   *firestore.Client
	bucket       *storage.BucketHandle
}

func (d *DriverFactory) GetSearchTeacher() search.IndexInterface {
	return d.teacherIndex
}

func (d *DriverFactory) GetSearchRequest() search.IndexInterface {
	return d.requestIndex
}

func (d *DriverFactory) GetCloudStore() *firestore.Client {
	return d.cloudStore
}

func (d *DriverFactory) GetStorage() *storage.BucketHandle {
	return d.bucket
}

type DriverInterface interface {
	GetSearchTeacher() search.IndexInterface
	GetSearchRequest() search.IndexInterface
	GetCloudStore() *firestore.Client
	GetStorage() 	*storage.BucketHandle
}

func GetDriver() DriverInterface {
	onceDo.Do(func() {
		driverInterface = &DriverFactory{
			teacherIndex: initSearchTeacher(),
			requestIndex: initSearchRequest(),
			cloudStore:   initCloudStore(),
			bucket:       initStorage(),
		}
	})
	return driverInterface
}

func initCloudStore() *firestore.Client {
	ctx := context.Background()
	path, err := web.AppConfig.String("firestore::path")
	if err != nil {
		log.Fatal(err, " driver/driver_factory.go:66")
	}
	sa := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	cloudStore, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return cloudStore
}

func initStorage() *storage.BucketHandle {
	ctx := context.Background()
	bucketString, err := web.AppConfig.String("storage::bucket")
	if err != nil {
		log.Fatal(err, " drivers/driver_factory.go:84")
	}
	config := &firebase.Config{
		StorageBucket: bucketString,
	}
	path, err := web.AppConfig.String("firestore::path")
	if err != nil {
		log.Fatal(err, " driver/driver_factory.go:91")
	}
	sa := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(ctx, config, sa)
	if err != nil {
		log.Fatalln(err, " drivers/driver_factory.go:95")
	}
	clientStorage, err := app.Storage(ctx)
	if err != nil {
		log.Fatalln(err, " drivers/driver_factory.go:100")
	}
	bucket, err := clientStorage.DefaultBucket()
	if err != nil {
		log.Fatalln(err, " drivers/driver_factory.go:104")
	}
	return bucket
}

func initSearchTeacher() search.IndexInterface {
	appID, err := web.AppConfig.String("algolia::app_id")
	if err != nil {
		log.Fatal(err)
	}
	key, err := web.AppConfig.String("algolia::api_key")
	if err != nil {
		log.Fatal(err)
	}
	clientSearch := search.NewClient(appID, key)
	searchTeacher := clientSearch.InitIndex("teacher")
	return searchTeacher
}

func initSearchRequest() search.IndexInterface {
	appID, err := web.AppConfig.String("algolia::app_id")
	if err != nil {
		log.Fatal(err)
	}
	key, err := web.AppConfig.String("algolia::api_key")
	if err != nil {
		log.Fatal(err)
	}
	clientSearch := search.NewClient(appID, key)
	searchRequest := clientSearch.InitIndex("request")
	return searchRequest
}

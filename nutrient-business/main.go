package main

import (
	"log"
	"os"

	"github.com/suhdev/nutrient-business/thirdparty"

	"github.com/suhdev/nutrient-business/handler"

	"github.com/micro/go-micro"
	"github.com/suhdev/nutrient/proton"
)

func main() {
	service := micro.NewService(micro.Name("DS.Client"))
	service.Init()
	ds := proton.NewDataStoreService("DS", service.Client())
	var (
		id, key string
		ok      bool
	)
	if id, ok = os.LookupEnv(proton.EnvThirdPartyAppID); !ok {
		log.Fatal("Could not find edmam app id")
	}
	if key, ok = os.LookupEnv(proton.EnvThirdPartyAppKey); !ok {
		log.Fatal("Could not find edmam app key")
	}
	api := thirdparty.NewEdmamAPI(id, key)
	bh := handler.NewBusinessHandler(ds, api)
	server := micro.NewService(micro.Name("Business"))
	server.Init()
	proton.RegisterBusinessHandler(service.Server(), bh)

	// go func() {
	// 	http.HandleFunc("/search", func(resp http.ResponseWriter, req *http.Request) {
	// 		result := &proton.Result{}
	// 		err := bh.Search(context.Background(), &proton.QueryRequest{
	// 			Q: "chicken",
	// 		}, result)
	// 		b, err := json.Marshal(result.Entries)
	// 		resp.Write(b)
	// 		fmt.Println(len(result.Entries), err)
	// 	})

	// 	http.ListenAndServe(":8585", nil)
	// }()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	// nc, _ := nats.Connect("nats://nutrient-nats:4222")
	// ticker := time.NewTicker(time.Second)
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		nc.Publish("box", []byte("GOOD ONE"))
	// 	}
	// }
}

package pkg

type SurveysServiceClient struct {
	// Client pb.SurveyServiceClient
}

func InitSurveyServiceClient(url string) SurveysServiceClient {
	// create secure connection to the gRPC server with the TLS credentials
	// clientConnection, err := grpc.Dial(url, grpc.WithInsecure())
	// if err != nil {
	// 	fmt.Println("Could not connect:", err)
	// }
	// TODO: create Client
	client := SurveysServiceClient{
		// Client: pb.NewSurveyServiceClient(clientConnection),
	}
	return client
}

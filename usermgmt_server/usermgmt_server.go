package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"

	pb "github.com/tech-with-moss/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(100))
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
}
func (os *UserManagementServer) SendDruidQuery(context.Context, *pb.DruidData) {
	client := &http.Client{}
	log.Printf("Entre a senddruidquery")
	// Si quieres agregar parámetros a la URL simplemente haz una
	// concatenación :)
	url := "https://videoamp.implycloud.com/druid/113f3d62-2b0e-43b2-8b3c-b3c32f98f90b/druid/v2/sql"

	type Query struct {
		query string
	}
	query := &Query{query: "SELECT COUNT(*) AS TheCount FROM 'csv-short-v4'"}
	//query = "SELECT COUNT(*) AS TheCount FROM 'csv-short-v4' "
	sqlJson, err := json.Marshal(query)

	if err != nil {
		log.Fatalf("Couldn't create json query", err)
		return
	}
	/*caCert, err := ioutil.ReadFile("/home/nioma/Desktop/Videoamp/113f3d62-2b0e-43b2-8b3c-b3c32f98f90b.crt")
	if err != nil {
		log.Fatalf("Could'nt read", err)
		return
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}*/
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(sqlJson))
	if err != nil {
		log.Fatalf("Error in request creation: %v", err)
	}

	request.SetBasicAuth("admin", "3Uxuwzz61lwWvy1l3VOYRg==")

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Request error: %v", err)
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Response error: %v", err)
	}

	// Aquí puedes decodificar la respuesta si es un JSON, o convertirla a cadena
	responseString := string(responseBody)
	log.Printf("Response code: %d", response.StatusCode)
	log.Printf("Header: '%q'", response.Header)
	contentType := response.Header.Get("Content-Type")
	log.Printf("Content type: '%s'", contentType)
	log.Printf("Response body: '%s'", responseString)
	return
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

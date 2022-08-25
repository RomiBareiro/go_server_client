package main

import (
	"context"
	"log"

	proto "github.com/RomiBareiro/go_server_client/tree/master/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewOrderServiceClient(conn)
	url := "https://videoamp.implycloud.com/druid/113f3d62-2b0e-43b2-8b3c-b3c32f98f90b/druid/v2/sql"

	//orders, err := client.ListOrders(context.Background(), &emptypb.Empty{})
	//order, err := client.GetOrder(context.Background(), &proto.GetRequest{Id: int32(5)})
	order, err := client.SendDruidQuery(context.Background(), &proto.DruidData{Url: url, User: "admin", Password: "3Uxuwzz61lwWvy1l3VOYRg=="})

	//	order, err := client.GetCount(context.Background(), &proto.TableName{TableName: "tabla"})

	/*if err != nil {
		log.Fatalln("client failed to fetch orders")
		return
	}*/
	/*if len(orders.Orders) > 0 {*/
	//log.Println("Data:", orders.Orders[0].Customer)
	log.Println("Data:", order)

}

/*
   	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
   	if err != nil {
   		log.Fatalf("did not connect: %v", err)
   	}
   	defer conn.Close()
   	c := pb.NewUserManagementClient(conn)

   	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
   	defer cancel()
   	var new_users = make(map[string]int32)
   	new_users["Alice"] = 43
   	new_users["Bob"] = 30
   	for name, age := range new_users {
   		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
   		if err != nil {
   			log.Fatalf("could not create user: %v", err)
   		}
   		log.Printf(`User Details:

   NAME: %s
   AGE: %d
   ID: %d`, r.GetName(), r.GetAge(), r.GetId())

   	}
*/

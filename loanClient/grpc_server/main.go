package main

// import (
// "log"
// "context"
// "google.golang.org/grpc"
// "loan/x/loan/types"
// "fmt"
// "net"
// "flag"
// )

// var (
// 	port = flag.Int("port", 50051, "the port to serve on")
// )

// // server is used to implement helloworld.GreeterServer.
// type server struct {
// 	types.UnimplementedMsgServer
// }

// // SayHello implements helloworld.GreeterServer
// func (s *server) SayHello(ctx context.Context, in *types.MsgRequestLoan) (*types.MsgRequestLoanResponse, error) {
// 	log.Printf("Received: %v", in.GetCreator())
// 	return &types.MsgRequestLoanResponse{}, nil
// }

// // func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// // 	return &types.MsgCreateLoanRequest{Message: "Hello again " + in.GetName()}, nil
// // }

// func main() {
// 	flag.Parse()
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	types.RegisterMsgServer(s, &server{})
// 	log.Printf("server listening at %v", lis.Addr())
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
import (
	"log"
	"net"
	// "context"
	"fmt"
	"flag"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	// "github.com/ignite/cli/ignite/pkg/cosmosclient"

	"loan/x/loan/types"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
// type msgserver struct {
// 	types.UnimplementedMsgServer
// }

// func (s *msgserver) RequestLoan(ctx context.Context, in *types.MsgRequestLoan) (*types.MsgRequestLoanResponse, error){
// 	return &types.MsgRequestLoanResponse{}, nil
// }
type queryserver struct {
	types.UnimplementedQueryServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//types.RegisterMsgServer(s, &msgserver{})
	types.RegisterQueryServer(s, &queryserver{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
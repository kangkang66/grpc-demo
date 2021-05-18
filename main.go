package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
	"gomod/mig_one/user"
	"gomod/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var methods map[string]*desc.MethodDescriptor

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	options := methods[info.FullMethod].GetOptions()
	nestedVal, err := proto.GetExtension(options,user.E_Api)
	//fmt.Println(goreflect.TypeOf(nestedVal))
	a,_ := nestedVal.(*user.HttpConfig)
	fmt.Println(a.GetToken())

	m, err := handler(ctx, req)
	return m, err
}

func main()  {
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	user.RegisterUserServer(s, service.UserHandle)
	reflection.Register(s)


	methods = make(map[string]*desc.MethodDescriptor)
	sds, _ := grpcreflect.LoadServiceDescriptors(s)
	for _, sd := range sds {
		for _, md := range sd.GetMethods() {
			methodName := fmt.Sprintf("/%s/%s", sd.GetFullyQualifiedName(), md.GetName())
			methods[methodName] = md
		}
	}
	fmt.Println(methods)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

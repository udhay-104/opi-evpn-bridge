// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.

// Package evpn is the main package of the application
package evpn

import (
	"context"
	"fmt"
	"log"

	"github.com/milosgajdos/tenus"
	pb "github.com/opiproject/opi-api/network/cloud/v1alpha1/gen/go"
)

// Server represents the Server object
type Server struct {
	pb.UnimplementedCloudInfraServiceServer
}

// CreateSubnet executes the creation of the subnet
func (s *Server) CreateSubnet(_ context.Context, in *pb.CreateSubnetRequest) (*pb.Subnet, error) {
	log.Printf("CreateSubnet: got %v", in)
	// Create a new network bridge
	br, err := tenus.NewBridgeWithName("mybridge")
	if err != nil {
		log.Fatal(err)
	}

	// Bring the bridge up
	if err = br.SetLinkUp(); err != nil {
		fmt.Println(err)
	}

	// Create a dummy link
	dl, err := tenus.NewLink("mydummylink")
	if err != nil {
		log.Fatal(err)
	}

	// Add the dummy link into bridge
	if err = br.AddSlaveIfc(dl.NetInterface()); err != nil {
		log.Fatal(err)
	}

	// Bring the dummy link up
	if err = dl.SetLinkUp(); err != nil {
		fmt.Println(err)
	}
	// TODO: replace cloud -> evpn
	return &pb.Subnet{}, nil
}
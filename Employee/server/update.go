package main

import (
	"context"
	"fmt"
	pb "github.com/sbbhagate/GoCode/Employee/proto"
	sng "github.com/sbbhagate/GoCode/Employee/singleton"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateEmployee(ctx context.Context, in *pb.Employee) (*pb.UpdateEmployeeResponse, error) {
	str := fmt.Sprintf("UpdateEmployee was invoked with %v\n", in)
	sng.DebugLogger.Debug(ctx, str, data)

	data := Emp{
		EmpId:       in.EmpId,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		DisplayName: in.DisplayName,
		Age:         in.Age,
		Salary:      in.Salary,
		Designation: in.Designation,
		Department:  in.Department,
	}

	res, err := sng.MongoService.Collection.UpdateOne(
		ctx,
		bson.M{"empid": in.EmpId},
		bson.M{"$set": data},
	)

	if err != nil {
		sng.ErrorLogger.Error(ctx,err,"Could not Update Record",data)
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	if res.MatchedCount == 0 {
		sng.ErrorLogger.Error(ctx,err,"Cannot find Employee with given Employee Id",data)
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find Employee with given Id",
		)
	}

	CreateMatView(ctx)

	return &pb.UpdateEmployeeResponse{Response: "Updated Successfully"}, nil
}
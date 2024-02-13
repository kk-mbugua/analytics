package auth

import (
	"context"
	authclient "main/cmd/client"
	"main/pkg/proto/pb"
)

type RequestMetadata struct {
	Authorization  string
	OrganisationID string
	UserID         string
	BranchID       string
	RequestAuth    string
	Roles          []string
}

func AuthRequest(ctx context.Context, accessToken string, methodName string) (*RequestMetadata, error) {
	authClient, err := authclient.NewAuthClient(
		"localhost:50050",
		"certs/client/server.crt",
		"certs/client/server.key",
		"certs/ca/server.crt",
	)
	if err != nil {
		return nil, err
	}
	data := &pb.GetAuthContextRequest{
		AccessToken: accessToken,
		MethodName:  methodName,
	}
	// Get the metadata from the context
	contextData, err := authClient.Client.GetAuthContext(ctx, data)
	if err != nil {
		return nil, err
	}

	return &RequestMetadata{
		Authorization:  contextData.Metadata.Authorization,
		OrganisationID: contextData.Metadata.OrganisationId,
		UserID:         contextData.Metadata.UserId,
		BranchID:       contextData.Metadata.BranchId,
		RequestAuth:    contextData.Metadata.RequestAuth,
		Roles:          contextData.Metadata.Roles,
	}, nil
}

func GetRequestMetadata(ctx context.Context) (*RequestMetadata, error) {
	data := RequestMetadata{
		UserID:         "",
		OrganisationID: "",
		BranchID:       "",
		RequestAuth:    "",
	}
	// retrieve already set values from the context.Value with keys coresponding to data keys and assert the values to be of type string. if it is not, assign empty string
	if userID, ok := ctx.Value(UserID).(string); ok {
		data.UserID = userID
	}
	if organisationID, ok := ctx.Value(OrganisationID).(string); ok {
		data.OrganisationID = organisationID
	}
	if branchID, ok := ctx.Value(BranchID).(string); ok {
		data.BranchID = branchID
	}
	if requestAuth, ok := ctx.Value(RequestAuth).(string); ok {
		data.RequestAuth = requestAuth
	}
	return &data, nil
}

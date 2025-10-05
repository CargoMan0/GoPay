package grpc_clients

/*
func NewOperationFeedClient(addr string) (currency.CurrencyServiceClient, *grpc.ClientConn, error) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(addr, dialOptions...)
	if err != nil {
		return nil, nil, fmt.Errorf("grpc.NewClient: %w", err)
	}

	client := currency.NewCurrencyServiceClient(conn)

	return client, conn, nil
}
*/

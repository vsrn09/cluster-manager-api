package cmavmware

import (
	"context"
	"crypto/tls"
	pb "github.com/samsung-cnct/cma-vmware/pkg/generated/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type VMWareClient struct {
	conn   *grpc.ClientConn
	client pb.ClusterClient
}

func CreateNewClient(hostname string, insecure bool) (VMWareClientInterface, error) {
	output := VMWareClient{}
	err := output.CreateNewClient(hostname, insecure)
	if err != nil {
		return nil, err
	}
	return &output, err
}

func (a *VMWareClient) CreateNewClient(hostname string, insecure bool) error {
	var err error
	if insecure {
		// This is for non TLS traffic
		a.conn, err = grpc.Dial(hostname, grpc.WithInsecure())
		if err != nil {
			return err
		}
	} else {
		// If TLS is enabled, we're going to create credentials, also using built in certificates
		var tlsConf tls.Config
		creds := credentials.NewTLS(&tlsConf)

		a.conn, err = grpc.Dial(hostname, grpc.WithTransportCredentials(creds))
		if err != nil {
			return err
		}
	}
	a.client = pb.NewClusterClient(a.conn)
	return nil
}

func (a *VMWareClient) Close() error {
	return a.conn.Close()
}

func (a *VMWareClient) SetClient(client pb.ClusterClient) {
	a.client = client
}

func (a *VMWareClient) CreateCluster(input CreateClusterInput) (CreateClusterOutput, error) {
	var machines []*pb.MachineSpec
	for _, j := range input.VMWare.Machines {
		machines = append(machines, &pb.MachineSpec{
			Host:     j.Host,
			Port:     int32(j.Port),
			Username: j.Username,
		})
	}
	result, err := a.client.CreateCluster(context.Background(), &pb.CreateClusterMsg{
		Name: input.Name,
		Provider: &pb.CreateClusterProviderSpec{
			Name:       VMWareProvider,
			K8SVersion: input.K8SVersion,
			Vmware: &pb.CreateClusterVMWareSpec{
				Namespace: input.VMWare.Namespace,
				Machines:  machines,
			},
			HighAvailability: input.HighAvailability,
			NetworkFabric:    input.NetworkFabric,
		},
	})
	if err != nil {
		return CreateClusterOutput{}, err
	}
	output := CreateClusterOutput{
		Cluster: ClusterItem{
			ID:     result.Cluster.Id,
			Name:   result.Cluster.Name,
			Status: result.Cluster.Status,
		},
	}
	return output, nil
}

func (a *VMWareClient) GetCluster(input GetClusterInput) (GetClusterOutput, error) {
	result, err := a.client.GetCluster(context.Background(), &pb.GetClusterMsg{
		Name: input.Name,
	})
	if err != nil {
		return GetClusterOutput{}, err
	}
	output := GetClusterOutput{
		Cluster: ClusterDetailItem{
			ID:         result.Cluster.Id,
			Name:       result.Cluster.Name,
			Status:     result.Cluster.Status,
			Kubeconfig: result.Cluster.Kubeconfig,
		},
	}
	return output, nil
}

func (a *VMWareClient) DeleteCluster(input DeleteClusterInput) (DeleteClusterOutput, error) {
	result, err := a.client.DeleteCluster(context.Background(), &pb.DeleteClusterMsg{
		Name: input.Name,
	})
	if err != nil {
		return DeleteClusterOutput{}, err
	}
	output := DeleteClusterOutput{
		Status: result.Status,
	}
	return output, nil
}

func (a *VMWareClient) ListClusters(input ListClusterInput) (ListClusterOutput, error) {
	var clusters []ClusterItem
	result, err := a.client.GetClusterList(context.Background(), &pb.GetClusterListMsg{})
	if err != nil {
		return ListClusterOutput{}, err
	}
	for _, j := range result.Clusters {
		clusters = append(clusters, ClusterItem{
			ID:     j.Id,
			Name:   j.Name,
			Status: j.Status,
		})
	}
	output := ListClusterOutput{
		Clusters: clusters,
	}
	return output, nil
}

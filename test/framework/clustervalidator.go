package framework

import (
	"context"
	"fmt"
	"time"

	machinerytypes "k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/controller/clientutil"
	"github.com/aws/eks-anywhere/pkg/retrier"
	clusterf "github.com/aws/eks-anywhere/test/framework/cluster"
	"github.com/aws/eks-anywhere/test/framework/cluster/validations"
)

func validationsForExpectedObjects() []clusterf.StateValidation {
	mediumRetier := retrier.NewWithMaxRetries(60, 5*time.Second)
	longRetier := retrier.NewWithMaxRetries(120, 5*time.Second)
	return []clusterf.StateValidation{
		clusterf.RetriableStateValidation(mediumRetier, validations.ValidateClusterReady),
		clusterf.RetriableStateValidation(mediumRetier, validations.ValidateEKSAObjects),
		clusterf.RetriableStateValidation(longRetier, validations.ValidateControlPlaneNodes),
		clusterf.RetriableStateValidation(longRetier, validations.ValidateWorkerNodes),
		clusterf.RetriableStateValidation(mediumRetier, validations.ValidateCilium),
	}
}

func validationsForClusterDoesNotExist() []clusterf.StateValidation {
	return []clusterf.StateValidation{
		clusterf.RetriableStateValidation(retrier.NewWithMaxRetries(60, 5*time.Second), validations.ValidateClusterDoesNotExist),
	}
}

func (e *ClusterE2ETest) buildClusterStateValidationConfig(ctx context.Context) {
	managementClusterClient, err := buildClusterClient(e.managementKubeconfigFilePath())
	if err != nil {
		e.T.Fatalf("failed to create management cluster client: %s", err)
	}
	clusterClient := managementClusterClient
	if e.managementKubeconfigFilePath() != e.kubeconfigFilePath() {
		clusterClient, err = buildClusterClient(e.kubeconfigFilePath())
	}
	if err != nil {
		e.T.Fatalf("failed to create cluster client: %s", err)
	}
	spec, err := buildClusterSpec(ctx, managementClusterClient, e.ClusterConfig)
	if err != nil {
		e.T.Fatalf("failed to build cluster spec with kubeconfig %s: %v", e.kubeconfigFilePath(), err)
	}

	e.clusterStateValidationConfig = &clusterf.StateValidationConfig{
		ClusterClient:           clusterClient,
		ManagementClusterClient: managementClusterClient,
		ClusterSpec:             spec,
	}
}

func newClusterStateValidator(config *clusterf.StateValidationConfig) *clusterf.StateValidator {
	return clusterf.NewStateValidator(*config)
}

func buildClusterClient(kubeconfigFileName string) (client.Client, error) {
	var clusterClient client.Client
	// Adding the retry logic here because the connection to the client does not always
	// succedd on the first try due to connection failure after the kubeconfig becomes
	// available in the cluster.
	err := retrier.Retry(12, 5*time.Second, func() error {
		c, err := kubernetes.NewRuntimeClientFromFileName(kubeconfigFileName)
		if err != nil {
			return fmt.Errorf("failed to build cluster client: %v", err)
		}
		clusterClient = c
		return nil
	})

	return clusterClient, err
}

func buildClusterSpec(ctx context.Context, client client.Client, config *cluster.Config) (*cluster.Spec, error) {
	clus := &v1alpha1.Cluster{}
	key := machinerytypes.NamespacedName{Namespace: config.Cluster.Namespace, Name: config.Cluster.Name}
	if err := client.Get(ctx, key, clus); err != nil {
		return nil, fmt.Errorf("failed to get cluster to build spec: %s", err)
	}
	configCp := config.DeepCopy()
	configCp.Cluster = clus
	spec, err := cluster.BuildSpecFromConfig(ctx, clientutil.NewKubeClient(client), configCp)
	if err != nil {
		return nil, fmt.Errorf("failed to build cluster spec from config: %s", err)
	}
	return spec, nil
}
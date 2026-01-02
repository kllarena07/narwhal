package orchestrator

import (
	"context"
	"fmt"

	"github.com/liquidmetal-dev/flintlock/api/services/microvm/v1alpha1"
	"github.com/liquidmetal-dev/flintlock/api/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	DefaultKernelImage = "ghcr.io/liquidmetal-dev/flintlock-kernel:5.10.77"
	DefaultOSImage     = "ghcr.io/liquidmetal-dev/capmvm-kubernetes:1.23.5"
)

type MicroVMOrchestrator struct {
	conn   *grpc.ClientConn
	client v1alpha1.MicroVMClient
	ctx    context.Context
	addr   string
}

func NewMicroVMOrchestrator(flintlockAddr string) (*MicroVMOrchestrator, error) {
	ctx := context.Background()
	
	conn, err := grpc.Dial(flintlockAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Flintlock: %w", err)
	}

	client := v1alpha1.NewMicroVMClient(conn)

	return &MicroVMOrchestrator{
		conn:   conn,
		client: client,
		ctx:    ctx,
		addr:   flintlockAddr,
	}, nil
}

func (m *MicroVMOrchestrator) CreateMicroVM(name, namespace, kernelPath, rootfsPath string, vcpus, memoryMB int32) (string, error) {
	if kernelPath == "" {
		kernelPath = DefaultKernelImage
	}
	if rootfsPath == "" {
		rootfsPath = DefaultOSImage
	}

	req := &v1alpha1.CreateMicroVMRequest{
		Microvm: &types.MicroVMSpec{
			Id:        name,
			Namespace: namespace,
			Vcpu:      vcpus,
			MemoryInMb: memoryMB,
			Kernel: &types.Kernel{
				Image: kernelPath,
				Cmdline: map[string]string{
					"console": "ttyS0",
					"reboot":  "k",
					"panic":   "1",
					"pci":     "off",
				},
			},
			RootVolume: &types.Volume{
				Id:         "root",
				IsReadOnly: false,
				Source: &types.VolumeSource{
					ContainerSource: &rootfsPath,
				},
			},
		},
	}

	resp, err := m.client.CreateMicroVM(m.ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to create microVM: %w", err)
	}

	if resp.Microvm.Spec != nil {
		if resp.Microvm.Spec.Uid != nil && *resp.Microvm.Spec.Uid != "" {
			return *resp.Microvm.Spec.Uid, nil
		}
		return resp.Microvm.Spec.Id, nil
	}
	return "", fmt.Errorf("microVM created but no ID returned")
}

func (m *MicroVMOrchestrator) GetMicroVM(uid string) (*types.MicroVM, error) {
	req := &v1alpha1.GetMicroVMRequest{
		Uid: uid,
	}

	resp, err := m.client.GetMicroVM(m.ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get microVM: %w", err)
	}

	return resp.Microvm, nil
}

func (m *MicroVMOrchestrator) ListMicroVMs(namespace string) ([]*types.MicroVM, error) {
	req := &v1alpha1.ListMicroVMsRequest{
		Namespace: namespace,
	}

	resp, err := m.client.ListMicroVMs(m.ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list microVMs: %w", err)
	}

	return resp.Microvm, nil
}

func (m *MicroVMOrchestrator) DeleteMicroVM(uid string) error {
	req := &v1alpha1.DeleteMicroVMRequest{
		Uid: uid,
	}

	_, err := m.client.DeleteMicroVM(m.ctx, req)
	if err != nil {
		return fmt.Errorf("failed to delete microVM: %w", err)
	}

	return nil
}

func (m *MicroVMOrchestrator) Close() error {
	if m.conn != nil {
		return m.conn.Close()
	}
	return nil
}

func (m *MicroVMOrchestrator) GetConnection() *grpc.ClientConn {
	return m.conn
}

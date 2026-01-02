package orchestrator

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MicroVMOrchestrator struct {
	conn   *grpc.ClientConn
	ctx    context.Context
	addr   string
}

func NewMicroVMOrchestrator(flintlockAddr string) (*MicroVMOrchestrator, error) {
	ctx := context.Background()
	
	conn, err := grpc.Dial(flintlockAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Flintlock: %w", err)
	}

	return &MicroVMOrchestrator{
		conn: conn,
		ctx:  ctx,
		addr: flintlockAddr,
	}, nil
}

func (m *MicroVMOrchestrator) CreateMicroVM(name, namespace, kernelPath, rootfsPath string, vcpus, memoryMB int64) (string, error) {
	return "", fmt.Errorf("not implemented: requires Flintlock protobuf definitions")
}

func (m *MicroVMOrchestrator) GetMicroVM(name, namespace string) (interface{}, error) {
	return nil, fmt.Errorf("not implemented: requires Flintlock protobuf definitions")
}

func (m *MicroVMOrchestrator) ListMicroVMs(namespace string) ([]interface{}, error) {
	return nil, fmt.Errorf("not implemented: requires Flintlock protobuf definitions")
}

func (m *MicroVMOrchestrator) DeleteMicroVM(name, namespace string) error {
	return fmt.Errorf("not implemented: requires Flintlock protobuf definitions")
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

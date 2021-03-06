package brokers

import (
	"github.com/gooops/machinery/v1/signatures"
)

// Broker - a common interface for all brokers
type Broker interface {
	SetRegisteredTaskNames(names []string)
	IsTaskRegistered(name string) bool
	StartConsuming(consumerTag string, p TaskProcessor) (bool, error)
	StopConsuming()
	Publish(task *signatures.TaskSignature) error
	GetPendingTasks(queue string) ([]*signatures.TaskSignature, error)
	GetALLPendingTasks(queue string) ([]*signatures.TaskSignature, error)
	DelPendingTask(signature *signatures.TaskSignature) (int64, error)
}

// TaskProcessor - can process a delivered task
// This will probably always be a worker instance
type TaskProcessor interface {
	Process(signature *signatures.TaskSignature) error
}

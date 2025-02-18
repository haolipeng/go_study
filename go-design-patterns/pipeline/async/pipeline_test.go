package async

import (
	"context"
	"testing"
)

func TestProcessorManager(t *testing.T) {
	m := NewProcessorManager()

	// pipeline组装
	m.AddSource(NewTimerSource(1, 2, 3, 4, 5))
	//m.AddSource(NewTimerSource(1, 2, -1, 3, 4, 5))
	m.AddSink(NewConsoleSink())

	m.AddError(NewErrorPolicyExit())

	m.AddProcessor(&SqProcessor{})
	m.AddProcessor(&SumProcessor{})

	// 执行
	m.Run(context.Background())
}

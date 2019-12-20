package snowflake

import (
    "testing"
)

func TestGenerate(t *testing.T) {
    worker, err := NewWorker(1, 0)
    if err != nil {
        t.Error(err)
    }
    for i := 0; i < 1000000; i++ {
        worker.Generate()
    }
}

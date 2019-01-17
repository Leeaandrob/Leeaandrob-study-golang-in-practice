package example_1_10_test

import "testing"

func TestName(t *testing.T) {
  name := getName()

  if name != "World" {
    t.Error("Response from getName is unexpected value")
  }
}

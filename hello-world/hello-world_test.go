package helloworld

import "testing"

func TestHelloWorld(t  *testing.T) {
  want := "fuck"
  got  := HelloWorld()

  if got != want {
    t.Errorf("HelloWorld() = %v; want %v", got, want)
  }
}

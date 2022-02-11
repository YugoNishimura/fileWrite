package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("YugoNishimura")
    want := "Aloha, YugoNishimura. Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}
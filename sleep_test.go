package Sleep
import "testing"
import "time"
func TestSleep(t *testing.T) {
    begin := time.Now()
    Sleep(6)
    end := time.Since(begin).Seconds()
    if end <  6 || end > 6.2 {
        t.Error("Incorrect sleep function")
    }
}
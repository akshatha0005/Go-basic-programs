package Sleep
import "fmt"
import "time"
func Sleep(p int) {
	fmt.Println("waiting for time",p,"Seconds")
	fmt.Println(time.Now())
	<-time.After(time.Duration(p) * time.Second)
	fmt.Println(time.Now())
}
	
package gateway

import "fmt"

func main() {

}

func run(svcPort int) {
	cmdStr := fmt.Sprintf("There  service is running in port : %s ", svcPort)
	fmt.Println(cmdStr)
}

package files

import (
	"fmt"
	"os"
)

func ReadFile() {

}

func WriteFile(contant string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(contant)
	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
	file.Close()
}

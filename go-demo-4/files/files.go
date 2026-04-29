package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func WriteFile(contant []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(contant)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}

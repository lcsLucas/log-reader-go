package utils

import (
	"fmt"
	"regexp"
)

func ParsingLineLog(s string) error {
	pattern, err := regexp.Compile(`(?P<IP>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s-\s-\s\[(?P<datahora>[^\[\]]+)\]\s"(?P<metodo>[A-Z]+)\s(?P<recurso>[^"]+)\sHTTP\/1\.[01]"\s(?P<status>\d{3})\s(?P<tamanho>\d+)\s"(?P<referer>[^"]+)"\s"(?P<user_agent>[^"]+)"`)

	if err != nil {
		return err
	}

	matches := pattern.FindAllStringSubmatch(s, -1)

	for _, v := range matches {
		fmt.Println("IP", v[1])
		fmt.Println("Data", v[2])
		fmt.Println("Método", v[3])
		fmt.Println("Recurso", v[4])
		fmt.Println("Status", v[5])
		fmt.Println("Tamanho", v[6])
		fmt.Println("User Agent", v[8])
	}

	return nil
}

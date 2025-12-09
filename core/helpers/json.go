package helpers

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

func ReadJsonFile(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var builder strings.Builder
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}
	return builder.String()
}

func SaveToJsonFile[T any](file *os.File, t *T) error {
	enc := json.NewEncoder(file)
	err := enc.Encode(t)
	return err
}

func LoadOrInitializeFromJsonFile[T any](file *os.File) (*T, error) {
	out := new(T)
	dec := json.NewDecoder(strings.NewReader(ReadJsonFile(file)))
	err := dec.Decode(out)
	if err == io.EOF {
		log.Println("JSON was empty - recreating...")
		err_enc := SaveToJsonFile(file, out)
		if err_enc != nil {
			return nil, err_enc
		}
	} else if err != nil {
		return nil, err
	}
	return out, nil
}

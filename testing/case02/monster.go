package monster

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) serialization() (bool, error) {
	var str = "D:\\Workspace\\go\\path\\src\\golang_learn\\data\\io\\monster.txt"
	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Println("parse struct 失败！")
		return false, err
	}
	file, err := os.OpenFile(str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file error=", err)
		return false, err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(string(jsonStr) + "\n")
	writer.Flush()
	return true, err
}

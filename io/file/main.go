package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var fileName = "D:\\Workspace\\go\\path\\src\\golang_learn\\data\\io\\test.txt"
	//var toFileName = "D:\\Workspace\\go\\path\\src\\golang_learn\\data\\io\\totest.txt"
	//ReadFileByCache(fileName)
	//ReadFileAll(fileName)

	//CopyFile(toFileName,fileName)
	count(fileName)

}

type CharCount struct {
	Chcount    int //记录英文个数
	NumCount   int //记录数字个数
	SpaceCount int //记录空格个数
	OtherCount int //记录其他字符个数
}

func count(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, world := range line {

			if world >= 'a' && world <= 'z' {
				count.Chcount++
			} else if world >= 'A' && world <= 'Z' {
				count.Chcount++
			} else if world == ' ' || world == '\t' {
				count.SpaceCount++
			} else if world >= 0 && world <= 9 {
				count.NumCount++
			} else {
				count.OtherCount++
			}
		}
	}

	fmt.Println("Chcount:", count.Chcount)
	fmt.Println("NumCount:", count.NumCount)
	fmt.Println("SpaceCount:", count.SpaceCount)
	fmt.Println("OtherCount:", count.OtherCount)

}

//判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//Copy 文件
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcfile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("-------------")
		fmt.Printf("open file err=%v", err)
	}
	defer srcfile.Close()

	reader := bufio.NewReader(srcfile)
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("-------------")
		fmt.Printf("open file err=%v", err)
	}

	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

//写文件操作实例
func writeDemo(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("创建失败！")
		fmt.Printf("open file err=%v", err)
		return
	}
	str := "hello, Garden\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

//一次性读取文件
func ReadFileAll(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%v", string(content))

}

//带缓冲读文件
func ReadFileByCache(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}

}

package main

import (
	"bufio"
	"fmt"
	"haoprogrammer/myspider/cmd/pipline"
	"os"
	"strconv"
)

func main() {
	//p := createPipline("small.in",512,4)
	//writeToFile(p, "small.out")
	//printFile("small.out")

	//large
	//p := createPipline("large.in",80000000,4)
	//writeToFile(p, "large.out")
	//printFile("large.out")

	//network-pipline
	p := createNetWorkPipline("small.in", 512, 4)
	writeToFile(p, "small.out")
	printFile("small.out")
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	p := pipline.ReaderSourceChunk(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}

}

func writeToFile(p <-chan int, filename string) {

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	pipline.WriterSink(writer, p)

}

//这里reader偷懒了，没有close掉
func createPipline(filename string,
	fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount

	//记录开始时间
	pipline.Init()
	var sortResults []<-chan int

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		//跳转文本中某处
		file.Seek(int64(i*chunkSize), 0)

		source := pipline.ReaderSourceChunk(
			bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipline.InMemSort(source))
	}

	return pipline.MergeN(sortResults...)

}

func createNetWorkPipline(filename string,
	fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount

	//记录开始时间
	pipline.Init()
	//var sortResults []<-chan int

	sortAddr := []string{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		//跳转文本中某处
		file.Seek(int64(i*chunkSize), 0)

		source := pipline.ReaderSourceChunk(
			bufio.NewReader(file), chunkSize)

		//sortResults = append(sortResults, pipline.InMemSort(source))
		addr := ":" + strconv.Itoa(7000+i)

		pipline.NetworktSink(addr, pipline.InMemSort(source))

		sortAddr = append(sortAddr, addr)
	}

	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipline.NetworkSource(addr))
	}

	return pipline.MergeN(sortResults...)

}

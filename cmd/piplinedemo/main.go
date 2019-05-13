package main

import (
	"bufio"
	"fmt"
	"haoprogrammer/myspider/cmd/pipline"
	"os"
)

func main() {
	const filename = "large.in"
	const n = 10000000
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipline.RandomSource(n)
	//pipline.WriterSink(file, p)
	//包装下，使得生成文件速度变快
	writer := bufio.NewWriter(file)
	pipline.WriterSink(writer, p)

	//需要Flush，确保数据都写入
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//p = pipline.ReaderSource(file)
	//p = pipline.ReaderSource(bufio.NewReader(file))
	p = pipline.ReaderSourceChunk(bufio.NewReader(file), -1)

	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}

}

func mergeDemo() {
	p := pipline.Merge(
		pipline.InMemSort(pipline.ArraySource(3, 2, 1, 5, 4, 6, 8)),
		pipline.InMemSort(pipline.ArraySource(3, 21, 5, 7, 10)))

	//for {
	//	//这里p  chan会关闭,需要ok判断chan状态
	//	if num, ok := <- p; ok {
	//		fmt.Println(num)
	//	}else {
	//		break
	//	}
	//}

	//用range,chan 需要关闭，不然range不知道什么时候结束,（我一直在等你发，你却没有通知我就走了）。range心态爆炸,会报错
	for v := range p {
		fmt.Println(v)
	}
}

package pipline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

//使用函数的人,返回值<-chan int  从chan里面拿东西
//对应函数内部,需要往chan里面放东西
func ArraySource(a ...int) <-chan int {
	out := make(chan int)

	//go 的 channel（信道）需要和goroutine结合使用
	go func() {
		for _, v := range a {
			out <- v
		}
		//关闭channel,一般不需要要手动关闭的，这里我们知道数据已传输完成，所以可以关闭
		close(out)
	}()

	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		//read into memory
		var a []int
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read done", time.Now().Sub(startTime))
		//sort
		sort.Ints(a)
		fmt.Println("InMemSort done", time.Now().Sub(startTime))

		//output
		for _, v := range a {
			out <- v
		}
		close(out)

	}()

	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)

	//合并两个结果
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}

		close(out)

		fmt.Println("Merge done", time.Now().Sub(startTime))

	}()

	return out
}

//函数名不要起成file这种具体的概念
//read节点
func ReaderSource(reader io.Reader) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				v := int(
					binary.BigEndian.Uint64(buffer))
				out <- v
			}

			if err != nil {
				break
			}
		}
		close(out)
	}()

	return out
}

func ReaderSourceChunk(reader io.Reader, chunkSize int) <-chan int {
	//out := make(chan int)
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				v := int(
					binary.BigEndian.Uint64(buffer))
				out <- v
			}

			//这里假定chunkSize -1 代表全部读
			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()

	return out
}

func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))

		writer.Write(buffer)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()

	return out
}

//递归完成N路两两归并
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2

	return Merge(
		MergeN(inputs[:m]...),
		MergeN(inputs[m:]...))

}

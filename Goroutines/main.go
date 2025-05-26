/*
package main

import (
	"fmt"
	"sync"
)

// A goroutine is a lightweight thread of execution managed by the Go runtime.
//Think of it like a background task or a thread, but way more lightweight and efficient.
// Multi thread we can achieve this via Go keyword

//When you use goroutines, your program may finish before the goroutines complete.
//A sync.WaitGroup lets you:
//Wait for all goroutines to finish ✅
//Prevent the main function from exiting too early 🛑

//func task[t any](id []t) {
//	fmt.Println("doing task", id)
//}

// defer keyword will execute after the fininshing of the func

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // Decrease the counter when done
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		//go task([]int{i})
		//go task([]string{"sa", "m", "b"})
		//go func(i int) {
		//	fmt.Println([]int{i})
		//}(i)
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait() // Wait until counter becomes 0

	//time.Sleep(2 * time.Second)

}
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadFile(url string, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("file downloading", url)
	time.Sleep(2 * time.Second)
	fmt.Println("file downloading finsihed", url)
}

func main() {
	var wg sync.WaitGroup
	files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"}

	for _, file := range files {
		wg.Add(1)
		go downloadFile(file, &wg)
	}

	wg.Wait()
	fmt.Println("All downloads complete ✅")
}

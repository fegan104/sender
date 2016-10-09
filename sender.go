package main

import (
    "bufio"
    "fmt"
    "github.com/kyokomi/emoji"
    "log"
    "os"
    "strings"
    "time"
    "github.com/fegan104/sender/foo"
)

func main() {
  foo.New()
  foo.Blah()
}

func PrintEomji() {
    fmt.Println("Hello Wolrd Emoji!")
    emoji.Println(":beer: Beer!!!")
    pizzaMessage := emoji.Sprint("I like a :pizza: and :sushi:!!")
    fmt.Println(pizzaMessage)
}

func ReadTestFile() {
    file, err := os.Open("test.txt")
    if err != nil {
      log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
       fmt.Println(scanner.Text())
    }
}

func ReadStdin() {
    reader := bufio.NewReader(os.Stdin)
    input := ""
    for input != "exit" {
      fmt.Print("Enter text: ")
      input, err := reader.ReadString('\n')
      if (err != nil){
        log.Fatal(err)
      }
      input = strings.TrimRight(input, "\n")
      fmt.Println(input)
    }
}

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func DoWork() {
    //In order to use our pool of workers we need to send them work and collect their results. We make 2 channels for this.
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    //This starts up 3 workers, initially blocked because there are no jobs yet.
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    //Here we send 9 jobs and then close that channel to indicate that’s all the work we have.
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    //Finally we collect all the results of the work.
    for a := 1; a <= 9; a++ {
        <-results
    }
}
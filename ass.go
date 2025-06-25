package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "sync"
)

func ProcessLogs(inputFiles []string, outputFile string) error {
    errorCh := make(chan string)      
    var wg sync.WaitGroup            

    done := make(chan error, 1)
    go func() {
        file, err := os.Create(outputFile)
        if err != nil {
            done <- fmt.Errorf("failed to create output file: %w", err)
            return
        }
        defer file.Close()

        writer := bufio.NewWriter(file)
        defer writer.Flush()

        for errLine := range errorCh {
            if _, err := writer.WriteString(errLine + "\n"); err != nil {
                done <- fmt.Errorf("failed to write to output file: %w", err)
                return
            }
        }
        done <- nil 
    }()

    for _, filename := range inputFiles {
        wg.Add(1)
        go func(file string) {
            defer wg.Done()

            f, err := os.Open(file)
            if err != nil {
                log.Printf("failed to open file %s: %v", file, err)
                return
            }
            defer f.Close()

            scanner := bufio.NewScanner(f)
            for scanner.Scan() {
                line := scanner.Text()
                if strings.Contains(line, "ERROR") {
                    errorCh <- line
                }
            }
            if err := scanner.Err(); err != nil {
                log.Printf("error reading file %s: %v", file, err)
            }
        }(filename)
    }
    go func() {
        wg.Wait()
        close(errorCh)
    }()

    return <-done
}

func main() {
    inputFiles := []string{"server1.log", "server2.log", "server3.log"}
    err := ProcessLogs(inputFiles, "errors.log")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Log processing complete.")
}

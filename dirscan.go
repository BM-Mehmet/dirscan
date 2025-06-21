/*
Bu program, bir hedef URL'deki "MSEC" kısmını verilen wordlist dosyasındaki kelimelerle değiştirerek HTTP istekleri gönderir.
Kullanıcıdan hedef URL (-u), wordlist dosya yolu (-w), eşzamanlı iş parçacığı sayısı (-t) ve eşleşecek HTTP durum kodu (-mc) parametrelerini alır.
Wordlist dosyasındaki her kelime için, URL'deki "MSEC" kısmı o kelimeyle değiştirilir ve HTTP GET isteği yapılır.
Eğer -mc parametresi verilmemişse (0 ise), tüm durum kodları ekrana yazdırılır. Eğer verilmişse, sadece o durum koduna sahip yanıtlar ekrana yazdırılır.
İşlemler, belirtilen sayıda eşzamanlı iş parçacığı (goroutine) ile gerçekleştirilir.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {

	var url string
	var wordlist string
	var threads int
	var matchStatus int

	flag.StringVar(&url, "u", "", "Target URL to scan")

	flag.StringVar(&wordlist, "w", "", "Path to the wordlist file")
	flag.IntVar(&threads, "t", 10, "Number of concurrent threads (default: 10)")
	flag.IntVar(&matchStatus, "mc",0, "HTTP status code to match (default: 0)")

	flag.Parse()
	if url == "" || wordlist == "" {
		fmt.Println("Usage: dirscan -u http://site.com/MSEC -w wordlist.txt")
		return
	}

	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Println("Error opening wordlist file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words []string

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			words = append(words, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading wordlist file:", err)
		return
	}
	jobs := make(chan string, threads)
	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for word := range jobs {
				msecURL := strings.ReplaceAll(url, "MSEC", word)
				resp, err := http.Get(msecURL)
				if err != nil {
					fmt.Println("[HATA]", msecURL, "-", err)
					continue
				}
				if matchStatus == 0 || resp.StatusCode == matchStatus {
					fmt.Printf("[HATA] Status %d for %s\n", resp.StatusCode, msecURL)
					resp.Body.Close()
				}
			}
		}()
	}

	for _, word := range words {
		jobs <- word
	}
	close(jobs)

	wg.Wait()
}

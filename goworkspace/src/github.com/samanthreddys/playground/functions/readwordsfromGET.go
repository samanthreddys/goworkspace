package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		words[sc.Text()] = ""
		//wordlist = append(wordlist, _)

	}

	//fmt.Println(words)

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading")

	}
	i := 0
	for k, _ := range words {
		fmt.Println(k)
		if i == 2 {
			break
		}
		i++
	}

	/* bs, _ := ioutil.ReadAll(res.Body)
	fmt.Println(bs)
	s := string(bs)
	fmt.Println(s)
	*/
}

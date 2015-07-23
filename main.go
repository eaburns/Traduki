package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var (
	srcLang = flag.String("f", "eo", "Source language")
	dstLang = flag.String("t", "en", "Destination language")
)

const (
	rawUrl     = "http://en.lernu.net/cgi-bin/serchi.pl"
	wordArg    = "modelo"
	srcLangArg = "delingvo"
	dstLangArg = "allingvo"
)

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		traduki(arg)
	}
}

func traduki(word string) {
	v := url.Values{}
	v.Set(srcLangArg, *srcLang)
	v.Set(dstLangArg, *dstLang)
	v.Set(wordArg, word)
	u, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}
	u.RawQuery = v.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("%s:	GET failed %s", word, err)
		return
	}
	defer resp.Body.Close()

	// TODO(eaburns): Actually parse the output.
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Printf("%s: failed to read response: %s", word, err)
		return
	}
	os.Stdout.WriteString("\n")
}

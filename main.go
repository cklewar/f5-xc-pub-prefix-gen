package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/nfx/go-htmltable"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const selector = "#content > div > div > div.data > table:nth-child(6)"
const url = "https://docs.cloud.f5.com/docs/reference/network-cloud-ref"
const prefixOffset = 3
const prefixFile = "prefixes.json"

type PublicIPv4SubnetRanges struct {
	Geography string   `header:"Geography"`
	Protocol  string   `header:"Protocol"`
	Ports     []string `header:"Ports"`
	Prefixes  []string `header:"Prefixes"`
	Notes     string   `header:"Notes"`
}

type PublicIPv4EgressSubnetRanges struct {
	Egress []string `json:"egress"`
}

type extendTable struct {
	*htmltable.Page
}

func (p *extendTable) Each4(a, b, c, d string, f func(a, b, c, d string) error) error {
	table, err := p.FindWithColumns(a, b, c, d)
	if err != nil {
		return err
	}
	offsets := map[string]int{}
	for idx, header := range table.Header {
		offsets[header] = idx
	}
	_1, _2, _3, _4 := offsets[a], offsets[b], offsets[c], offsets[d]
	for idx, row := range table.Rows {
		if len(row) < 3 {
			continue
		}
		err = f(row[_1], row[_2], row[_3], row[_4])
		if err != nil {
			return fmt.Errorf("row %d: %w", idx, err)
		}
	}
	return nil
}

func main() {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc2, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	expand := doc2.Find(selector)
	expand = expand.BeforeHtml("<table>")
	expand = expand.AfterHtml("</table>")
	var b bytes.Buffer
	err = goquery.Render(&b, expand)
	if err != nil {
		return
	}

	pisrs := make([]PublicIPv4SubnetRanges, 0)
	piesrs := PublicIPv4EgressSubnetRanges{}

	jsonFile, err := os.Open(prefixFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully opened file: %s\n", prefixFile)
	defer func(jsonFile *os.File) error {
		err := jsonFile.Close()
		if err != nil {
			log.Fatalf("Error closing file: %s\n", prefixFile)
			return err
		}
		return nil
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &piesrs)
	if err != nil {
		log.Fatalf("Error reading regions json data %s\n", err.Error())
	}

	table, _ := htmltable.NewFromString(b.String())
	eTable := &extendTable{table}
	_ = eTable.Each4("Geography", "Protocol", "IP Address", "Ports", func(geography string, protocol string, ipPrefixes string, ports string) error {
		currentPrefixes := map[string]bool{}
		_ports := map[string]bool{}
		fmt.Printf("Geography: %s ----- Protocol: %s ----- Ports: %s ----- IP Prefixes: %s\n", geography, protocol, ports, ipPrefixes)
		lenIpPrefixes := len(ipPrefixes)
		lenCurrentPrefix := strings.Index(ipPrefixes, "/") + prefixOffset
		lenNewIpPrefixes := len(ipPrefixes) - lenCurrentPrefix
		currentPrefix := ipPrefixes[:len(ipPrefixes)-lenNewIpPrefixes]
		fmt.Println("Current prefix", currentPrefix, "Length current prefix:",
			lenCurrentPrefix, "Length new prefix list:", lenNewIpPrefixes, "Length prefix list:", lenIpPrefixes)
		ipPrefixes = ipPrefixes[lenCurrentPrefix:]
		pisr := PublicIPv4SubnetRanges{
			Geography: geography,
			Protocol:  protocol,
		}

		_, ok := currentPrefixes[currentPrefix]
		if !ok {
			currentPrefixes[currentPrefix] = true
		}

		_tmp := strings.Split(ports, ",")
		for _, p := range _tmp {
			_, ok := _ports[strings.TrimSpace(p)]
			if !ok {
				_ports[strings.TrimSpace(p)] = true
			}
		}

		for i := 0; i <= len(ipPrefixes); i++ {
			lenIpPrefixes = len(ipPrefixes)
			lenCurrentPrefix = strings.Index(ipPrefixes, "/") + prefixOffset
			lenNewIpPrefixes = len(ipPrefixes) - lenCurrentPrefix
			currentPrefix = ipPrefixes[:len(ipPrefixes)-lenNewIpPrefixes]
			fmt.Println("Current prefix", currentPrefix, "Length current prefix:",
				lenCurrentPrefix, "Length new prefix list:", lenNewIpPrefixes, "Length prefix list:", lenIpPrefixes)
			ipPrefixes = ipPrefixes[lenCurrentPrefix:]
			_, ok := currentPrefixes[currentPrefix]
			if !ok {
				currentPrefixes[currentPrefix] = true
			}

			_tmp := strings.Split(ports, ",")
			for _, p := range _tmp {
				_, ok := _ports[strings.TrimSpace(p)]
				if !ok {
					_ports[strings.TrimSpace(p)] = true
				}
			}
		}

		for k := range currentPrefixes {
			pisr.Prefixes = append(pisr.Prefixes, k)
		}

		for k := range _ports {
			pisr.Ports = append(pisr.Ports, k)
		}

		pisrs = append(pisrs, pisr)

		return nil
	})
	// fmt.Println(pisrs)
}

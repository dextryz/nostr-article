package nostrarticlecli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	np "github.com/dextryz/nostr-pipeline"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

type Config struct {
	Npub   string   `json:"npub"`
	Relays []string `json:"relays"`
}

func StringEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env variable \"%s\" not set, usual", key)
	}
	return value
}

var CONFIG = StringEnv("NOSTR_CONFIG")

func loadConfig() (*Config, error) {
	b, err := os.ReadFile(CONFIG)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}
	if len(cfg.Relays) == 0 {
		log.Println("please set atleast on relay in your config.json")
	}
	return &cfg, nil
}

func Titles(cfg *Config, limit int) {

	_, pk, err := nip19.Decode(cfg.Npub)
	if err != nil {
		panic(err)
	}

	f := nostr.Filter{
		Kinds:   []int{nostr.KindArticle},
		Authors: []string{pk.(string)},
		Limit:   limit,
	}
	naddrs := np.New(cfg.Relays[0]).Filter(f).Query().Naddrs()

	for _, naddr := range naddrs {

		prefix, data, err := nip19.Decode(naddr)
		if err != nil {
			log.Fatalf("shouldn't error: %s", err)
		}

		if prefix != "naddr" {
			log.Fatalf("returned invalid prefix")
		}

		ep := data.(nostr.EntityPointer)
		if ep.Kind != nostr.KindArticle {
			log.Fatalf("returned wrong kind")
		}

		fmt.Println(ep.Identifier)
		fmt.Println(naddr)
		fmt.Println("")
	}
}

func Main() int {

	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	args := os.Args[1:]

	if args[0] == "list" {

        if len(args) != 2 {
            log.Fatalln("provide the number of articles to list")
        }

		limit, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalln(err)
		}
		Titles(cfg, limit)
	} else {
		log.Fatalln("option not aviable")
		return 1
	}

	return 0
}

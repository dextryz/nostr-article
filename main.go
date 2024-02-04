package nostrarticlecli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	np "github.com/dextryz/nostr-pipeline"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

type article struct {
	input  io.Reader
	output io.Writer
}

type option func(*article) error

func WithInputFromArgs(args []string) option {
	return func(c *article) error {
		if len(args) < 1 {
			return nil
		}
		c.input = strings.NewReader(args[0])
		return nil
	}
}

func New(opts ...option) (*article, error) {
	c := &article{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *article) Title() int {
	lines := 0
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		//pipeline := np.New("wss://relay.damus.io/")
		//titles := wpipeline.Authors([]string{input.Text()}).Kinds([]int{nostr.KindArticle}).Query().Titles()

        ids := np.New("wss://relay.damus.io/").Authors([]string{input.Text()}).Kinds([]int{nostr.KindArticle}).Query().Ids("wss://relay.damus.io/", "npub14ge829c4pvgx24c35qts3sv82wc2xwcmgng93tzp6d52k9de2xgqq0y4jk")

        for _, naddr := range ids {

            prefix, data, err := nip19.Decode(naddr)
            if err != nil {
                fmt.Errorf("shouldn't error: %s", err)
            }

            if prefix != "naddr" {
                fmt.Errorf("returned invalid prefix")
            }

            ep := data.(nostr.EntityPointer)
            if ep.Kind != nostr.KindArticle {
                fmt.Errorf("returned wrong kind")
            }

			fmt.Println(ep.Identifier)
			fmt.Println(naddr)
			fmt.Println("")
		}

	}
	return lines
}

func (c *article) Tags() int {
	lines := 0
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		pipeline := np.New("wss://relay.damus.io/")
		pipeline.Authors([]string{input.Text()}).Kinds([]int{nostr.KindArticle}).Query().Tags().SortByCount().Stdout()
	}
	return lines
}

func Main() int {
	c, err := New(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	c.Title()
	return 0
}

package nostrarticlecli

import (
	np "github.com/dextryz/nostr-pipeline"
	"github.com/nbd-wtf/go-nostr"
)

func Main() {

	npub := "npub14ge829c4pvgx24c35qts3sv82wc2xwcmgng93tzp6d52k9de2xgqq0y4jk"

	pipeline := np.New("wss://relay.damus.io/")
	pipeline.Authors([]string{npub}).Kinds([]int{nostr.KindArticle}).Query().Tags().SortByCount().Stdout()

}

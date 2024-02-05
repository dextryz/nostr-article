# Nostr Article CLI

Read-only CLI for long form content on [nostr](www.nostr.com). To manage your article with write permissions use the [Nostr Knowledge Management](https://github.com/dextryz/nkm) client.

## Setup

Create your config file in `~/.config/nostr/article.json` containing:

```
{
  "relays": ["wss://relay.damus.io"],
  "npub": "npub14ge829c4pvgx24c35qts3sv82wc2xwcmgng93tzp6d52k9de2xgqq0y4jk"
}
```

Finally, set your environment variable:

```shell
export NOSTR_CONFIG=~/.config/nostr/article.json`
```

Build the executable

```shell
make build
```

## Usage

View two articles

```shell
> nart list 2
The Art of War
naddr...

Channels in Go
naddr...
```

## Todo

```shell
# List articles whoes content contains the following keyword.
> nart search "hello friend"

# List articles with the following tags.
> nart tagged "coding,go"
```

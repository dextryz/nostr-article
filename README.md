# Nostr Article CLI

Nostr CLI to manage long form content

Create your `config.json` in `~/.config/nostr/config.json` like below

```
{
  "relays": ["wss://relay.damus.io"],
  "nsec": "nsec1xxxxxx"
}
```

```shell
> nart list
The Art of War
naddr...

Channels in Go
naddr...

> nart search "hello friend"

> nart tagged "coding,go"
```

Tag a note

```shell
> nart tag 202402060643.md coding go nostr

> nart title 202402060643.md 'My New Note'
```

Commit a markdown file to the set of relays

```shell
> nart commit 202402060643.md
```

## Ideas

Myabe have it sudh that you env set the current article:

```shell
> nart new
touch 202402060643.md
export ARTICLE=202402060643.md
```


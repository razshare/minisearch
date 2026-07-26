# What is this?

This is a very naive proof of concept search engine written 
in [Go](https://go.dev/) and [Svelte](https://svelte.dev/) using 
[Frizzante](https://github.com/razshare/frizzante).

# Get Started (Prebuilt Binaries)

Download the [migrate and serve programs](https://github.com/razshare/minisearch/releases).

Run the migrate program.

```sh
./migrate
```

Run the server.

```sh
./serve
```

# Get Started (Build From Source)

> [!NOTE]
> You will need the frizzante cli, for more details see the [official documentation](https://razshare.github.io/frizzante-docs/).

Configure project.

```sh
make configure
```

Build project.

```sh
make build
```

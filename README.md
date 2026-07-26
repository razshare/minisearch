# What is this?

This is a very naive proof of concept search engine written 
in [Go](https://go.dev/) and [Svelte](https://svelte.dev/) using 
[Frizzante](https://github.com/razshare/frizzante).

# Get Started (Prebuilt Binaries)

Download the [migrate program and the server](https://github.com/razshare/minisearch/releases)

Run the migrate program.

```sh
./migrate
```

> [!NOTE]
> This will create a `./source.sqlite` file, your database.


Run the server.

```sh
./serve
```

> [!NOTE]
> This will serve the website at `http://127.0.0.1:38123`.

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

> [!NOTE]
> This will generate your binaries to `.gen/bin`.

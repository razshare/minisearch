# What is this?

This is a very naive proof of concept search engine written 
in [Go](https://go.dev/) and [Svelte](https://svelte.dev/) using 
[Frizzante](https://github.com/razshare/frizzante).

# Get Started (Prebuilt Binaries)

Download the [migrate program and the server](https://github.com/razshare/minisearch/releases)

Run the migrate program.

```sh
./migrate-linux-amd64
```

> [!NOTE]
> This will create a `./source.sqlite` file, your database.


Run the server.

```sh
./serve-linux-amd64
```

> [!NOTE]
> This will serve the website at `http://127.0.0.1:38123`.

# Get Started (Build From Source)

> [!NOTE]
> You will need the frizzante cli, for more details see the [official documentation](https://razshare.github.io/frizzante-docs/).

Configure project.

```sh
frizzante configure
```

Migrate development.

```sh
frizzante migrate
```

Start development.

```sh
frizzante dev
```

Build production.

```sh
frizzante build
```
    
> [!NOTE]
> This will create two executables, `.gen/bin/migrate` and `.gen/bin/serve`.

Migrate production.

```sh
.gen/bin/migrate
```

Serve production.

```sh
.gen/bin/serve
    ```
# Perfecty Push Server ⚡️

![Tests](https://github.com/rwngallego/perfecty-push/workflows/Tests/badge.svg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

Self-hosted Push Notifications server written in Go.

![Perfecty Push for Wordpress](.github/assets/logo-white.png)

Send thousands of notifications from your server:
No hidden fees, no third party dependencies and you own your data. 👏

**Perfecty Push Server** also integrates with:
 - [Wordpress](https://wordpress.org/plugins/perfecty-push-notifications/)

## Local Setup 👨🏻‍💻

Generate a sample TLS certificate:

```sh
mkdir examples/
openssl req -newkey rsa:2048 -nodes -keyout examples/certs/server.key -x509 -days 365 -out examples/certs/server.crt

# Update the configs:
vi configs/perfecty.yml
  ...
  ssl:
    enabled: true
    cert_file: examples/server.crt
    key_file: examples/server.key
```

Execute the project:

```sh
go run cmd/perfecty/main.go
```

Generate executable:

```shell
go build ./cmd/perfecty/
./perfecty
```

## Configuration 🛠

You can change the values in `config/perfecty.yml`.

## Unit tests 🧪

Run all the tests:

```sh
go test -v ./...
```

## License 💡

This project is licensed under [MIT](LICENSE).

## Collaborators 🔥

[<img alt="rwngallego" src="https://avatars3.githubusercontent.com/u/691521?s=460&u=ceab22655f55101b66f8e79ed08007e2f8034f34&v=4" width="117">](https://github.com/rwngallego) |
:---: |
[Rowinson Gallego](https://www.linkedin.com/in/rwngallego/) |

## Special Thanks

[<img alt="Jetbrains" src="https://github.com/rwngallego/perfecty-push-wp/raw/master/.github/assets/jetbrains-logo.svg" width="120">](https://www.jetbrains.com/?from=PerfectyPush)

Thanks to Jetbrains for supporting this Open Source project with their magnificent tools.
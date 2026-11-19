# VeBoost

## Run VeMerkle with Docker Compose

Run the following commands from the repository root. Docker Engine and Docker
Compose v2 or newer are required; Go is only required for builds outside Docker.

```sh
cp compose.env.example .env
cp merkle/.env.example merkle/.env
cp merkle/configs/airdrop.docker.example.yaml merkle/configs/airdrop.docker.yaml
```

Skip copying any file that already exists. Before starting:

- Set the PostgreSQL DSN, RPC URL, airdrop contract address and Swagger password
  in `merkle/configs/airdrop.docker.yaml`.
- Set `SLACK_BOT_TOKEN` and `SLACK_SIGNING_SECRET` in `merkle/.env` for Slack
  integration. This file must exist because the application loads it at startup.
- Select the image and host port in the root `.env`. The two `.env` files serve
  different purposes: the root file configures Compose, while `merkle/.env`
  contains application settings.

PostgreSQL remains an external dependency; this Compose file does not create or
replace a database. `127.0.0.1` inside the container points to the container, not
the Docker host. Use `host.docker.internal` for a database on the host, or the
appropriate reachable hostname for a remote database. PostgreSQL must accept
connections from the container network. Use the appropriate SSL settings for
your database.

### Deploy a published image

The default image is `ghcr.io/bedrock-technology/veboost:dev`. If the GHCR package
is private, first authenticate using `docker login ghcr.io` with a token that
has permission to read the package.

```sh
docker compose pull
docker compose up -d --no-build
docker compose logs -f vemerkle
```

The service runs `VeMerkle run --port 8080`. It is available on host port `8080`
by default, or the port selected with `VEMERKLE_PORT`. Swagger is available at
`http://localhost:8080/docs/index.html` using the configured credentials.

For production, set `VEMERKLE_IMAGE` to a specific release image tag rather than
the moving `dev` or `latest` tag. After changing the image, repeat `pull` and
`up -d --no-build`.

### Build locally

```sh
docker compose build
docker compose up -d --no-build --pull never
```

If Docker Hub is unreachable, build with Docker Official Images mirrored on
Amazon ECR Public, then run the same `up` command:

```sh
docker compose build --build-arg BASE_REGISTRY=public.ecr.aws/docker/library
```

The multi-stage build uses Go 1.27.4, runs the Go tests, and produces a minimal
runtime image with CA certificates. The service runs as UID/GID `10001`, not
root. Mounted configuration files must be readable by this user.

The image includes the binary and Swagger specification, but not application
configuration, `.env` files, or credentials. Configuration files are mounted
read-only. Logs under `/app/logs` and exported files under `/app/data` use named
volumes; use `/app/data/<filename>.csv` when requesting a persistent CSV export.

```sh
docker compose cp vemerkle:/app/data/rewards.csv ./rewards.csv
docker compose down
```

`docker compose down` preserves the named volumes. Do not add `--volumes` unless
you intend to delete the persisted logs and exported files.

## Image builds in GitHub Actions

Workflows live in `.github/workflows/`:

- Pushing the `dev` branch publishes `dev`, `dev-last`, and a unique
  `dev-<version>-<UTC date>-<commit>` tag.
- Pushing any Git tag publishes the normalized release tag, `pro`, `latest`,
  and a unique `pro-<version>-<UTC date>-<commit>` tag.

Both workflows build `linux/amd64` images from the root `Dockerfile` and publish
to `ghcr.io/<lowercase repository name>`. For this repository, that is
`ghcr.io/bedrock-technology/veboost`. They authenticate with `GITHUB_TOKEN`; the
repository must allow GitHub Actions to write packages. Local builds use the
Docker host's architecture.

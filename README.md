[![PRs Welcome][pr-welcome-shield]][pr-welcome-url]
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <img alt="delineate.io" src="https://github.com/delineateio/.github/blob/master/assets/logo.png?raw=true" height="75" />
  <h2 align="center">delineate.io</h2>
  <p align="center">portray or describe (something) precisely.</p>

  <h3 align="center">Multi-Service Redis Cache Example</h3>

  <p align="center">
    Demonstrates using Redis as a low latency cache serving a Go & Python microservices
    <br />
    <a href="https://github.com/delineateio/multi-service-shared-cache"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/delineateio/multi-service-shared-cache">View Demo</a>
    ·
    <a href="https://github.com/delineateio/multi-service-shared-cache/issues">Report Bug</a>
    ·
    <a href="https://github.com/delineateio/multi-service-shared-cache/issues">Request Feature</a>
  </p>
</p>

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [About The Project](#about-the-project)
- [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Local Dependencies](#local-dependencies)
  - [Local Setup](#local-setup)
- [Running Local Services](#running-local-services)
- [Usage](#usage)
  - [Go Services](#go-services)
  - [Go Python](#go-python)
  - [Running & Debugging Locally](#running--debugging-locally)
- [Service Implementation](#service-implementation)
- [Go Services Implementation](#go-services-implementation)
- [Python Services Implementation](#python-services-implementation)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

<!-- ABOUT THE PROJECT -->
## About The Project

This repo demonstrates a few key cloud native integrations/elements:

* Using [Traefik Proxy](https://doc.traefik.io/traefik/) as a reverse proxy across multiple microservices
* Using [Redis](https://redis.io/) as an in memory cache across microservices
* Usage of various different popular Go packages (see [Go Services Implementation](#go-services-implementation))
* Usage of [FastAPI](https://fastapi.tiangolo.com/) as a web development framework in Python
* Multistage Docker image build to ensure optimised images

## Built With

Further logos can be inserted to highlight the specific technologies used to create the solution from [here](https://github.com/Ileriayo/markdown-badges).

| Syntax | Description |
| --- | ----------- |
| ![pre-commit](https://img.shields.io/badge/precommit-%235835CC.svg?style=for-the-badge&logo=precommit&logoColor=white) | Pre-commit `git` hooks that perform checks before pushes|
| ![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white) | Source control management platform  |
| ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) | Containerise applications and provide local environment |
| ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) | Containerise applications and provide local environment |
| ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) | Implementation of an example microservice |
| ![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) | Implementation of an example microservice |
| ![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white) | Provides im memory key value cache |

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Local Dependencies

A number of local dependencies are required.  To review the local dependencies run `task dependencies:list`.  If new local dependencies then they should be added to the correct Taskfile in `./os` e.g. `taskfile.darwin.yaml`.

> Note that currently only `macOS` is configured and a PR should be submitted if either `Linux` or `Windows` are required.

### Local Setup

This repo follows the principle of minimal manual setup of the local development environment.

 A `task` target has been provided for simplicity ```task init```, the `taskfile.yaml` file can be inspected for more details.

## Running Local Services

A `docker-compose` file is provided that can be used to manage local environment services.  The stack can be found at `ops/local/stack.yaml`.

```shell
# stands up the local services
task local:up

# tears down the local services
task local:down
```

<!-- USAGE EXAMPLES -->

## Usage

Once the local environment is running then the microservice APIs can be called.

### Go Services

```shell
# checks that the container is up
http localhost/healthz host:go.delineate.io

# increments the counter
http POST localhost/increment host:go.delineate.io

# returns the current count
http localhost/count host:go.delineate.io
```

### Go Python

```shell
# checks that the container is up
http localhost/healthz host:python.delineate.io

# increments the counter
http POST localhost/increment host:python.delineate.io

# returns the current count
http localhost/count host:python.delineate.io
```

### Running & Debugging Locally

```shell
# runs the golang app locally
cd ./dev/services/go-service/src
go run *.go

# see other commands above
http :1102/count

# runs the python app locally
cd ./dev/services/python-service/src
pip install -r ../requirements.txt
uvicorn main:app --reload

# see other commands above
http :8000/count
```

## Service Implementation

## Go Services Implementation

The following 3rd party libraries where used in the implementation...

* [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
* [github.com/rs/zerolog](https://github.com/rs/zerolog)
* [github.com/go-redis/redis/v9](https://github.com/go-redis/redis)
* [github.com/spf13/viper](https://github.com/spf13/viper)

## Python Services Implementation

The following 3rd party libraries where used in the implementation...

* [FastAPI](https://fastapi.tiangolo.com/)

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/delineateio/multi-service-shared-cache/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

If you would like to contribute to any delineate.io OSS projects please read:

* [Code of Conduct](https://github.com/delineateio/.github/blob/master/CODE_OF_CONDUCT.md)
* [Contributing Guidelines](https://github.com/delineateio/.github/blob/master/CONTRIBUTING.md)

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

* [Best README Template](https://github.com/othneildrew/Best-README-Template)
* [Markdown Badges](https://github.com/Ileriayo/markdown-badges)
* [DocToc](https://github.com/thlorenz/doctoc)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[pr-welcome-shield]: https://img.shields.io/badge/PRs-welcome-ff69b4.svg?style=for-the-badge&logo=github
[pr-welcome-url]: https://github.com/delineateio/multi-service-shared-cache/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue
[contributors-shield]: https://img.shields.io/github/contributors/delineateio/multi-service-shared-cache.svg?style=for-the-badge&logo=github
[contributors-url]: https://github.com/delineateio/multi-service-shared-cache/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/delineateio/multi-service-shared-cache.svg?style=for-the-badge&logo=github
[forks-url]: https://github.com/delineateio/multi-service-shared-cache/network/members
[stars-shield]: https://img.shields.io/github/stars/delineateio/multi-service-shared-cache.svg?style=for-the-badge&logo=github
[stars-url]: https://github.com/delineateio/multi-service-shared-cache/stargazers
[issues-shield]: https://img.shields.io/github/issues/delineateio/multi-service-shared-cache.svg?style=for-the-badge&logo=github
[issues-url]: https://github.com/delineateio/multi-service-shared-cache/issues
[license-shield]: https://img.shields.io/github/license/delineateio/multi-service-shared-cache.svg?style=for-the-badge&logo=github
[license-url]: https://github.com/delineateio/multi-service-shared-cache/blob/master/LICENSE

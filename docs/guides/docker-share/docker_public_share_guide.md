---
sidebar_position: 10
sidebar_label: Public Share
---

# Docker Public Share

With zrok, you can publicly share a server app that's running in another Docker container, or any server that's reachable by the zrok container.

## Walkthrough Video

<iframe width="100%" height="315" src="https://www.youtube.com/embed/ycov--9ZtB4" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

## Before You Begin

To follow this guide you will need [Docker](https://docs.docker.com/get-docker/) and [the Docker Compose plugin](https://docs.docker.com/compose/install/) for running `docker compose` commands in your terminal.

## Public Share with Docker Compose

1. Make a folder on your computer to use as a Docker Compose project for your zrok public share.
1. In your terminal, change directory to your newly-created project folder.
1. Download [the zrok-public-share Docker Compose project file](pathname:///zrok-public-share/docker-compose.yml) into your new project folder.
1. Copy your zrok environment token from the zrok web console to your clipboard and paste it in a file named `.env` in the same folder like this:

    ```bash
    # file name ".env"
    ZROK_ENABLE_TOKEN="8UL9-48rN0ua"
    ```

1. If you are self-hosting zrok then it's important to set your API endpoint URL too. If you're using the hosted zrok service then you can skip this step.

    ```bash
    # file name ".env"
    ZROK_API_ENDPOINT="https://zrok.example.com"
    ```

1. Run your Compose project to start sharing the built-in demo web server:

    ```bash
    docker compose up
    ```

1. Read the public share URL from the output. One of the last lines is like this:

    ```bash
    zrok-public-share-1  |  https://w6r1vesearkj.in.zrok.io/
    ```

    You can swap in a different server app container instead of the demo server, or you could change the Docker network to "host" and share something running on the Docker host's localhost interface.

1. Edit the file `docker-compose.yml`. Replace the following line:

    ```yaml
    command: share public --headless http://zrok-test:9090
    ```

    Replace that line with this to start sharing the HTTPBin server app container instead of the zrok test endpoint.

    ```yaml
    command: share public --headless http://httpbin-test:8080
    ```

1. Re-run your project to load the new server configuration.

    ```bash
    docker-compose up --force-recreate
    ```

    Now you'll have a new public share URL for the `httpbin` API testing server.

1. Run "down" to destroy the Compose project when you're done. Then delete the selected zrok environment by clicking "Actions" in the web console.

    ```bash
    docker compose down --remove-orphans --volumes
    ```

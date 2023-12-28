# CertGuardian

## Overview

CertGuardian is a monitoring interface for ssl certificates.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [License](#license)

## Installation

Running CertGuardian

```bash
docker run --name certguardian -d -p 7070:7070 1704mori/certguardian:latest
```
or with compose file
```bash
version: "3"

services:
  dokka:
    image: 1704mori/certguardian:latest
    container_name: certguardian
    ports:
      - 7070:7070
```

CertGuardian will be available at http://localhost:7070

## Configuration

CertGuardian can be configured using environment variables.

| Variable | Default | Description |
| --- | --- | --- |
| `PORT` | `7070` | Port to listen on |
| `CRON_INTERVAL` | `1d` | Interval for the cron job. e.g: 15d (15 days) or 3h (3 hours) |
| `NEAR_EXPIRY_THRESHOLD` | `10d` | Threshold to determine if a certificate is near expiration. e.g: 15d (15 days) or 3h (3 hours) |

## Todo

- [ ] Alerts

## License

This project is licensed under the [MIT License](LICENSE).

# srvctl

Lightweight CLI tool for everyday server operations written in Go.

srvctl is a simple and fast command-line utility for managing and diagnosing Linux servers.  
It provides quick access to system status, systemd services, logs, and Docker operations from a single interface.

Designed for DevOps and SRE workflows without unnecessary complexity.

## Features

- System status overview (CPU, memory, disk)
- Manage systemd services (start, stop, restart, status)
- View logs via journalctl
- Basic Docker operations
- Fast access to common server tasks
- Minimal and clean CLI interface

## Installation

Build from source:

git clone https://github.com/EgorIstomin/srvctl
cd srvctl
go build -o srvctl

(Optional) Install globally:

sudo mv srvctl /usr/local/bin/

## Usage

Show system status:

srvctl status

Manage services:

srvctl service restart nginx
srvctl service status docker

View logs:

srvctl logs nginx

Docker commands:

srvctl docker ps
srvctl docker images

## Requirements

- Linux
- systemd
- journalctl
- Docker (optional for docker commands)

## Why

srvctl was created as a fast CLI tool for everyday server management without needing to remember long system commands.

---

Minimal. Fast. Practical.

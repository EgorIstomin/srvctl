# srvctl

srvctl is a small CLI tool for basic server operations like checking system status, managing systemd services, viewing logs and running simple Docker commands.

It wraps common Linux admin tasks into one simple command so you don’t have to remember long systemctl, journalctl or docker commands.

Build:

go build -o srvctl .

Optional install:

sudo install -m 0755 srvctl /usr/local/bin/srvctl

Usage:

srvctl <command> [args]

Commands:

srvctl sys uptime        → show uptime  
srvctl sys df            → show disk usage  
srvctl sys free          → show memory usage  
srvctl sys status        → hostname + uptime + disk + memory  

srvctl svc status nginx  → check service status  
srvctl svc restart nginx → restart service  

srvctl logs nginx        → show last 100 log lines  
srvctl logs nginx --lines 200 → custom log lines  

srvctl docker ps         → show running containers  

Requires Linux with systemd.  
Service restart needs proper permissions.

srvctl executes system commands with timeouts to avoid hanging processes and returns real stderr when something fails.

Use it for quick diagnostics, service checks, log viewing and simple operational tasks.

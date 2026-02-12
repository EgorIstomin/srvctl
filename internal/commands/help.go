package commands

import "fmt"

func PrintHelp() {
	fmt.Println("srvctl - server control tool")
	fmt.Println()
	fmt.Println("usage:")
	fmt.Println("  srvctl <command> [args]")
	fmt.Println()
	fmt.Println("commands:")
	fmt.Println("  sys       system info")
	fmt.Println("  svc       systemd service commands")
	fmt.Println("  logs      view service logs (journalctl)")
	fmt.Println("  docker    docker commands")
	fmt.Println()
	fmt.Println("examples:")
	fmt.Println("  srvctl sys uptime")
	fmt.Println("  srvctl sys df")
	fmt.Println("  srvctl sys free")
	fmt.Println("  srvctl svc status nginx")
	fmt.Println("  srvctl svc restart nginx")
	fmt.Println("  srvctl logs nginx --lines 100")
	fmt.Println("  srvctl docker ps")
}


package main  // импортиуем главный entry-point пакет 

import ( 
	"fmt" // пакет для вывода в консоль
	"os"  // пакет для взаимодействия с окружением 

	"srvctl/internal/commands" // импортуем пакет 
)

func main() {
	args := os.Args // 

	if len(args) < 2 {
		commands.PrintHelp()
		return
	}

	command := args[1]

	switch command {
	case "help", "-h", "--help":
		commands.PrintHelp()
		return

	case "sys":
		if err := commands.HandleSys(args[2:]); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		return

	case "svc":
		if err := commands.HandleSvc(args[2:]); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		return

	case "logs":
		if err := commands.HandleLogs(args[2:]); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		return

	case "docker":
		if err := commands.HandleDocker(args[2:]); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		return

	default:
		fmt.Println("unknown command:", command)
		commands.PrintHelp()
		os.Exit(1)
	}
}

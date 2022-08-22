package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"brexs-test/domain"
	"brexs-test/http"
	"brexs-test/services"

	"github.com/sirupsen/logrus"
)

type args struct {
	action   string
	fileName string
}

var (
	httpServer *http.Server
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	})
}

func main() {

	logrus.Infof("application started")

	args, err := parseArgs()
	if err != nil {
		printHelp()
		return
	}

	routesBusiness := services.NewRoutesBusiness(args.fileName)

	switch args.action {
	case "server":
		ctx, cancel := context.WithCancel(context.Background())
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			_ = cleanup()
			cancel()
		}()
		//start http server
		httpServer = http.NewServer("4000", routesBusiness)
		httpServer.Open()
		<-ctx.Done()

	case "console":
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("please enter the route: ")
			text, _ := reader.ReadString('\n')
			st := strings.Split(text, "-")
			input := domain.RouteSchema{
				Origin:  strings.ReplaceAll(st[0], "\n", ""),
				Destiny: strings.ReplaceAll(st[1], "\n", ""),
				// Cost: 0,
			}

			routesBusiness.ReadFile()
			bestRoute := routesBusiness.FindBestRoute(input)

			fmt.Println("best route: ", bestRoute)
		}

	default:
		printHelp()
	}

}

// clean the application
func cleanup() error {
	if httpServer != nil {
		if err := httpServer.Close(); err != nil {
			return err
		}
	}
	return nil
}

func parseArgs() (*args, error) {
	// the first index of the cliArgs slice is always the Go filename.
	cliArgs := os.Args

	if len(cliArgs) < 3 {
		return nil, fmt.Errorf("invalid input")
	}

	return &args{
		action:   cliArgs[1],
		fileName: cliArgs[2],
	}, nil
}

func printHelp() {
	fmt.Println(`
usage:
- go run cmd/main.go server LISTA_ROTAS.CSV
	Inicializa o servidor de API
- go run cmd/main.go console LISTA_ROTAS.CSV
	Inicializa a aplicação no console`)
}

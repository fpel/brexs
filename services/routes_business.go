package services

import (
	"brexs-test/domain"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type RoutesBusiness struct {
	fileName string
	Content  []*domain.RouteSchema
}

func NewRoutesBusiness(fileName string) *RoutesBusiness {
	_, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return &RoutesBusiness{
		fileName: fileName,
	}
}

func (r *RoutesBusiness) ReadFile() {
	content := openFile(r.fileName)
	lines := strings.Split(content, "\n")

	c := make([]*domain.RouteSchema, len(lines))
	for i, linha := range lines {
		d := strings.Split(linha, ",")
		intVar, _ := strconv.Atoi(d[2])
		c[i] = &domain.RouteSchema{
			Origin:  d[0],
			Destiny: d[1],
			Cost:    intVar,
		}
	}
	r.Content = c
}

func (r *RoutesBusiness) FindBestRoute(route domain.RouteSchema) string {
	rout, cost := r.findRoute(route.Origin, route.Destiny)
	return fmt.Sprintf("%s > $%d", rout, cost)
}

func (r *RoutesBusiness) SaveFile(route domain.RouteSchema) error {
	content := openFile(r.fileName)
	content = content + "\n" + route.Origin + "," + route.Destiny + "," + strconv.Itoa(route.Cost)
	return os.WriteFile(r.fileName, []byte(content), 0666)

}

func (r *RoutesBusiness) findRoute(ori string, de string) (string, int) {
	cost := 0
	routeList := make(map[string]int)
	route := ""
	for _, d := range r.Content {

		if d.Origin == ori {
			route = route + d.Origin + "-"
			cost += d.Cost

			if d.Destiny != de {
				//troca a origem para buscar a outra linha
				origin := d.Destiny
				stop := false
				for !stop {
					for i, d1 := range r.Content {
						if d1.Origin == origin {
							route = route + origin + "-"
							cost += d1.Cost
							//encontrei o ultimo destino
							if d1.Destiny == de {
								routeList[route+de] = cost
								cost = 0
								route = ""
								stop = true
								break
							} else {
								origin = d1.Destiny
								break
							}
						}
						//Quando chega ao final e nao encontra o destino, precisa sair do laco
						if (i + 1) == len(r.Content) {
							stop = true
						}
					}
				}
			} else {
				routeList[d.Origin+"-"+d.Destiny] = d.Cost
				route = ""
				cost = 0
			}
		}
	}
	best := 0
	ret := ""
	x := 0
	for i, f := range routeList {
		if x == 0 {
			best = f
			ret = i
		}
		if best > f {
			best = f
			ret = i
		}
		x++
	}

	return ret, best
}

func openFile(fileName string) string {
	file, err := os.Open(fileName)
	fileStats, _ := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, fileStats.Size())
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes\n", count)

	return string(data)
}

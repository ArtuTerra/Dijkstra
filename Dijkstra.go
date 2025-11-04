package main

import (
	"dijsktra/grafo"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

type Resultado struct {
	Caminho        []string `json:"caminho"`
	DistanciaTotal float64  `json:"distanciaTotal"`
	TempoTotal     int      `json:"tempoTotal"`
	CriterioUsado  string   `json:"criterioUsado"`
}

var verticesNome = []string{
	"Av. Paulista, 1000",
	"Rua Augusta, 500",
	"Rua Consolação, 200",
	"Rua Oscar Freire, 100",
	"Alameda Santos, 300",
	"Praça da República, 50",
	"Av. Brigadeiro, 800",
	"Shopping Iguatemi",
	"Av. Rebouças, 1500",
	"Av. Ibirapuera, 2000",
	"Parque Ibirapuera - Portão 2",
	"Av. São João, 1200",
	"Terminal Parque Dom Pedro",
	"Marginal Pinheiros, Km 5",
}

func main() {
	http.HandleFunc("/", postHandler)

	fmt.Println("Servidor iniciado na porta 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}

func iniciarDijkstra(origem string, destino string, criterio string) (*Resultado, error) {
	g := grafo.NewGrafo()

	return encontrarMelhorRota(g, origem, destino, criterio)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusInternalServerError)
		return
	}

	type Dados struct {
		Origem   string `json:"origem"`
		Destino  string `json:"destino"`
		Criterio string `json:"criterio"`
	}
	var dados Dados

	err = json.Unmarshal(body, &dados)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Dados recebidos: %+v\n", dados.Origem)
	fmt.Printf("Dados recebidos: %+v\n", dados.Destino)
	fmt.Printf("Dados recebidos: %+v\n", dados.Criterio)

	resultado, err := iniciarDijkstra(dados.Origem, dados.Destino, dados.Criterio)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	jsonBytes, err := json.Marshal(resultado)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(jsonBytes))
}

func encontrarMelhorRota(g grafo.Grafo, origem string, destino string, criterio string) (*Resultado, error) {
	if err := validarRota(g, origem, destino); err != nil {
		return nil, err
	}

	if err := validarCriterio(criterio); err != nil {
		return nil, err
	}

	pontuacao := map[string]float64{}
	anterior := map[string]string{}
	explorado := map[string]bool{}

	for _, nome := range verticesNome {
		explorado[nome] = false
		pontuacao[nome] = math.Inf(1)
		if g.Vertices[nome].Nome == origem {
			pontuacao[nome] = 0
		}
	}

	for {

		menor := struct {
			nome      string
			pontuacao float64
		}{
			nome:      "",
			pontuacao: math.Inf(1),
		}

		for _, nome := range verticesNome {
			if explorado[nome] == false && pontuacao[nome] < menor.pontuacao {
				menor.nome = nome
				menor.pontuacao = pontuacao[nome]
			}
		}

		if menor.nome == "" || menor.nome == destino {
			break
		}

		explorado[menor.nome] = true

		for _, aresta := range g.Vertices[menor.nome].Arestas {
			var novaPontuacao float64
			if criterio == "tempo" {
				novaPontuacao = pontuacao[menor.nome] + float64(aresta.Dados.Tempo)
			} else {
				novaPontuacao = pontuacao[menor.nome] + aresta.Dados.Distancia
			}

			if novaPontuacao < pontuacao[aresta.Nome] {
				pontuacao[aresta.Nome] = novaPontuacao
				anterior[aresta.Nome] = menor.nome
			}

		}
	}

	if pontuacao[destino] == math.Inf(1) {
		return nil, errors.New("não existe caminho entre origem e destino")
	}

	caminho := []string{destino}
	atual := destino
	for atual != origem {
		atual = anterior[atual]
		caminho = append([]string{atual}, caminho...)
	}

	distanciaTotal := 0.0
	tempoTotal := 0

	for i := 0; i < len(caminho)-1; i++ {
		verticeAtual := caminho[i]
		verticeProximo := caminho[i+1]

		for _, aresta := range g.Vertices[verticeAtual].Arestas {
			if aresta.Nome == verticeProximo {
				distanciaTotal += aresta.Dados.Distancia
				tempoTotal += aresta.Dados.Tempo
				break
			}
		}
	}

	return &Resultado{
		Caminho:        caminho,
		DistanciaTotal: distanciaTotal,
		TempoTotal:     tempoTotal,
		CriterioUsado:  criterio,
	}, nil
}

func validarRota(g grafo.Grafo, origem string, destino string) error {
	if _, exists := g.Vertices[origem]; !exists {
		return errors.New("origem não existe")
	}

	if _, exists := g.Vertices[destino]; !exists {
		return errors.New("destino não existe")
	}
	return nil
}

func validarCriterio(criterio string) error {
	if criterio != "tempo" && criterio != "distancia" {
		return errors.New("critério inválido")
	}
	return nil
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

type distanciaTempo struct {
	distancia float64
	tempo     int
}

type aresta struct {
	nome  string
	dados distanciaTempo
}

type vertice struct {
	nome    string
	arestas []aresta
}

type grafo struct {
	vertices map[string]vertice
}

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

func inicializarGrafo() grafo {
	g := grafo{
		vertices: make(map[string]vertice),
	}

	for _, nome := range verticesNome {
		g.vertices[nome] = vertice{nome: nome, arestas: []aresta{}}
	}

	g.adicionarAresta("Av. Paulista, 1000", "Rua Augusta, 500", 2.3, 8)
	g.adicionarAresta("Av. Paulista, 1000", "Rua Consolação, 200", 1.5, 12)

	g.adicionarAresta("Rua Augusta, 500", "Av. Paulista, 1000", 2.3, 8)
	g.adicionarAresta("Rua Augusta, 500", "Rua Oscar Freire, 100", 1.8, 6)
	g.adicionarAresta("Rua Augusta, 500", "Alameda Santos, 300", 1.2, 5)

	g.adicionarAresta("Rua Consolação, 200", "Av. Paulista, 1000", 1.5, 12)
	g.adicionarAresta("Rua Consolação, 200", "Praça da República, 50", 2.0, 15)
	g.adicionarAresta("Rua Consolação, 200", "Alameda Santos, 300", 1.0, 7)

	g.adicionarAresta("Rua Oscar Freire, 100", "Rua Augusta, 500", 1.8, 6)
	g.adicionarAresta("Rua Oscar Freire, 100", "Shopping Iguatemi", 3.5, 10)
	g.adicionarAresta("Rua Oscar Freire, 100", "Av. Rebouças, 1500", 2.2, 9)

	g.adicionarAresta("Alameda Santos, 300", "Rua Augusta, 500", 1.2, 5)
	g.adicionarAresta("Alameda Santos, 300", "Rua Consolação, 200", 1.0, 7)
	g.adicionarAresta("Alameda Santos, 300", "Av. Brigadeiro, 800", 1.8, 6)
	g.adicionarAresta("Alameda Santos, 300", "Praça da República, 50", 1.5, 10)

	g.adicionarAresta("Praça da República, 50", "Rua Consolação, 200", 2.0, 15)
	g.adicionarAresta("Praça da República, 50", "Alameda Santos, 300", 1.5, 10)
	g.adicionarAresta("Praça da República, 50", "Av. São João, 1200", 1.0, 8)
	g.adicionarAresta("Praça da República, 50", "Terminal Parque Dom Pedro", 2.5, 18)

	g.adicionarAresta("Av. Brigadeiro, 800", "Alameda Santos, 300", 1.8, 6)
	g.adicionarAresta("Av. Brigadeiro, 800", "Av. Ibirapuera, 2000", 3.0, 11)
	g.adicionarAresta("Av. Brigadeiro, 800", "Shopping Iguatemi", 2.5, 8)

	g.adicionarAresta("Shopping Iguatemi", "Rua Oscar Freire, 100", 3.5, 10)
	g.adicionarAresta("Shopping Iguatemi", "Av. Brigadeiro, 800", 2.5, 8)
	g.adicionarAresta("Shopping Iguatemi", "Av. Ibirapuera, 2000", 4.0, 12)
	g.adicionarAresta("Shopping Iguatemi", "Av. Rebouças, 1500", 1.5, 5)

	g.adicionarAresta("Av. Rebouças, 1500", "Rua Oscar Freire, 100", 2.2, 9)
	g.adicionarAresta("Av. Rebouças, 1500", "Shopping Iguatemi", 1.5, 5)
	g.adicionarAresta("Av. Rebouças, 1500", "Marginal Pinheiros, Km 5", 3.8, 14)

	g.adicionarAresta("Av. Ibirapuera, 2000", "Av. Brigadeiro, 800", 3.0, 11)
	g.adicionarAresta("Av. Ibirapuera, 2000", "Shopping Iguatemi", 4.0, 12)
	g.adicionarAresta("Av. Ibirapuera, 2000", "Parque Ibirapuera - Portão 2", 1.0, 4)
	g.adicionarAresta("Av. Ibirapuera, 2000", "Terminal Parque Dom Pedro", 5.5, 25)

	g.adicionarAresta("Parque Ibirapuera - Portão 2", "Av. Ibirapuera, 2000", 1.0, 4)
	g.adicionarAresta("Parque Ibirapuera - Portão 2", "Av. São João, 1200", 6.0, 28)

	g.adicionarAresta("Av. São João, 1200", "Praça da República, 50", 1.0, 8)
	g.adicionarAresta("Av. São João, 1200", "Terminal Parque Dom Pedro", 2.0, 12)
	g.adicionarAresta("Av. São João, 1200", "Parque Ibirapuera - Portão 2", 6.0, 28)

	g.adicionarAresta("Terminal Parque Dom Pedro", "Praça da República, 50", 2.5, 18)
	g.adicionarAresta("Terminal Parque Dom Pedro", "Av. São João, 1200", 2.0, 12)
	g.adicionarAresta("Terminal Parque Dom Pedro", "Av. Ibirapuera, 2000", 5.5, 25)
	g.adicionarAresta("Terminal Parque Dom Pedro", "Marginal Pinheiros, Km 5", 8.0, 22)

	g.adicionarAresta("Marginal Pinheiros, Km 5", "Av. Rebouças, 1500", 3.8, 14)
	g.adicionarAresta("Marginal Pinheiros, Km 5", "Terminal Parque Dom Pedro", 8.0, 22)

	return g
}

func (g *grafo) adicionarAresta(origem string, destino string, distancia float64, tempo int) {
	if v, exists := g.vertices[origem]; exists {
		v.arestas = append(v.arestas, aresta{
			nome: destino,
			dados: distanciaTempo{
				distancia: distancia,
				tempo:     tempo,
			},
		})

		g.vertices[origem] = v
	}
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
	g := grafo{}
	g = inicializarGrafo()

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

func encontrarMelhorRota(g grafo, origem string, destino string, criterio string) (*Resultado, error) {
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
		if g.vertices[nome].nome == origem {
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

		for _, aresta := range g.vertices[menor.nome].arestas {
			var novaPontuacao float64
			if criterio == "tempo" {
				novaPontuacao = pontuacao[menor.nome] + float64(aresta.dados.tempo)
			} else {
				novaPontuacao = pontuacao[menor.nome] + aresta.dados.distancia
			}

			if novaPontuacao < pontuacao[aresta.nome] {
				pontuacao[aresta.nome] = novaPontuacao
				anterior[aresta.nome] = menor.nome
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

		for _, aresta := range g.vertices[verticeAtual].arestas {
			if aresta.nome == verticeProximo {
				distanciaTotal += aresta.dados.distancia
				tempoTotal += aresta.dados.tempo
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

func validarRota(g grafo, origem string, destino string) error {
	if _, exists := g.vertices[origem]; !exists {
		return errors.New("origem não existe")
	}

	if _, exists := g.vertices[destino]; !exists {
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

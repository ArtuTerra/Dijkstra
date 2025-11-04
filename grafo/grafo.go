package grafo

type distanciaTempo struct {
	Distancia float64
	Tempo     int
}

type aresta struct {
	Nome  string
	Dados distanciaTempo
}

type vertice struct {
	Nome    string
	Arestas []aresta
}

type Grafo struct {
	Vertices map[string]vertice
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

func NewGrafo() Grafo {
	g := Grafo{
		Vertices: make(map[string]vertice),
	}

	for _, nome := range verticesNome {
		g.Vertices[nome] = vertice{Nome: nome, Arestas: []aresta{}}
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

func (g *Grafo) adicionarAresta(origem string, destino string, distancia float64, tempo int) {
	if v, exists := g.Vertices[origem]; exists {
		v.Arestas = append(v.Arestas, aresta{
			Nome: destino,
			Dados: distanciaTempo{
				Distancia: distancia,
				Tempo:     tempo,
			},
		})

		g.Vertices[origem] = v
	}
}

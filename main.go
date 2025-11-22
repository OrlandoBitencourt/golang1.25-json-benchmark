package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/OrlandoBitencourt/golang1.25-json-benchmark/models"
)

func main() {
	fmt.Println("🚀 Gerador de dados para benchmark JSON v1 vs v2")
	fmt.Println()

	// Seed para gerar dados aleatórios consistentes
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Gerar dados de teste
	generateTestData()

	fmt.Println()
	fmt.Println("✅ Dados gerados com sucesso!")
	fmt.Println()
	fmt.Println("📊 Para rodar os benchmarks, execute:")
	fmt.Println("   # V1 (sem experimental)")
	fmt.Println("   go test ./benchmarks/... -bench=V1 -benchmem -benchtime=10s > v1_results.txt")
	fmt.Println()
	fmt.Println("   # V2 (com experimental)")
	fmt.Println("   GOEXPERIMENT=jsonv2 go test ./benchmarks/... -bench=V2 -benchmem -benchtime=10s > v2_results.txt")
	fmt.Println()
	fmt.Println("   # Comparar resultados")
	fmt.Println("   benchstat v1_results.txt v2_results.txt")
}

func generateTestData() {
	// Criar diretório se não existir
	if err := os.MkdirAll("testdata", 0755); err != nil {
		fmt.Printf("❌ Erro ao criar diretório testdata: %v\n", err)
		return
	}

	// Pequeno: 1 usuário (~1KB)
	fmt.Print("Gerando testdata/small.json... ")
	small := generateUsers(1)
	if err := saveJSON("testdata/small.json", small); err != nil {
		fmt.Printf("❌ Erro: %v\n", err)
		return
	}
	fmt.Println("✓ (~1KB)")

	// Médio: 100 usuários (~100KB)
	fmt.Print("Gerando testdata/medium.json... ")
	medium := generateUsers(100)
	if err := saveJSON("testdata/medium.json", medium); err != nil {
		fmt.Printf("❌ Erro: %v\n", err)
		return
	}
	fmt.Println("✓ (~100KB)")

	// Grande: 10000 usuários (~10MB)
	fmt.Print("Gerando testdata/large.json... ")
	large := generateUsers(10000)
	if err := saveJSON("testdata/large.json", large); err != nil {
		fmt.Printf("❌ Erro: %v\n", err)
		return
	}
	fmt.Println("✓ (~10MB)")

	// Logs: 1000 entradas de log para teste de streaming (~500KB)
	fmt.Print("Gerando testdata/logs.jsonl... ")
	logs := generateLogs(1000)
	if err := saveJSONLines("testdata/logs.jsonl", logs); err != nil {
		fmt.Printf("❌ Erro: %v\n", err)
		return
	}
	fmt.Println("✓ (~500KB)")
}

func generateUsers(count int) []models.User {
	users := make([]models.User, count)
	for i := 0; i < count; i++ {
		users[i] = models.User{
			ID:        i + 1,
			Username:  fmt.Sprintf("user%d", i+1),
			Email:     fmt.Sprintf("user%d@example.com", i+1),
			FirstName: randomName(),
			LastName:  randomName(),
			Active:    rand.Intn(2) == 1,
			CreatedAt: randomDate(),
			UpdatedAt: time.Now(),
			Profile: models.Profile{
				Bio:      randomBio(),
				Avatar:   fmt.Sprintf("https://avatar.example.com/%d", i+1),
				Location: randomLocation(),
				Website:  fmt.Sprintf("https://user%d.example.com", i+1),
				Metadata: map[string]string{
					"theme":    randomTheme(),
					"language": randomLanguage(),
					"timezone": "UTC-3",
				},
			},
			Roles: randomRoles(),
		}
	}
	return users
}

func generateLogs(count int) []models.LogEntry {
	logs := make([]models.LogEntry, count)
	levels := []string{"DEBUG", "INFO", "WARN", "ERROR"}
	services := []string{"api", "worker", "scheduler", "mailer"}

	for i := 0; i < count; i++ {
		logs[i] = models.LogEntry{
			Timestamp: time.Now().Add(-time.Duration(rand.Intn(3600)) * time.Second),
			Level:     levels[rand.Intn(len(levels))],
			Service:   services[rand.Intn(len(services))],
			Message:   randomLogMessage(),
			Context: map[string]interface{}{
				"user_id":     rand.Intn(1000),
				"request_id":  fmt.Sprintf("req_%d", rand.Intn(10000)),
				"duration_ms": rand.Intn(1000),
			},
			TraceID: fmt.Sprintf("trace_%d", rand.Intn(100000)),
		}
	}
	return logs
}

func saveJSON(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func saveJSONLines(filename string, logs []models.LogEntry) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	for _, log := range logs {
		if err := encoder.Encode(log); err != nil {
			return err
		}
	}
	return nil
}

// Funções auxiliares para gerar dados aleatórios

func randomName() string {
	names := []string{
		"João", "Maria", "Pedro", "Ana", "Carlos", "Juliana",
		"Lucas", "Fernanda", "Rafael", "Beatriz", "Felipe", "Camila",
		"Gabriel", "Amanda", "Bruno", "Patricia", "Rodrigo", "Letícia",
	}
	return names[rand.Intn(len(names))]
}

func randomBio() string {
	bios := []string{
		"Desenvolvedor apaixonado por Go e sistemas distribuídos",
		"Engenheira de software especialista em backend e APIs",
		"Tech lead focado em performance e escalabilidade",
		"Arquiteto de soluções cloud native e microsserviços",
		"Full stack developer com foco em Go e React",
		"DevOps engineer automatizando tudo que é possível",
		"Site Reliability Engineer mantendo sistemas no ar 24/7",
		"Backend developer construindo APIs que escalam",
	}
	return bios[rand.Intn(len(bios))]
}

func randomLocation() string {
	locations := []string{
		"São Paulo, BR",
		"Rio de Janeiro, BR",
		"Blumenau, BR",
		"Florianópolis, BR",
		"Curitiba, BR",
		"Porto Alegre, BR",
		"Belo Horizonte, BR",
		"Brasília, BR",
	}
	return locations[rand.Intn(len(locations))]
}

func randomTheme() string {
	themes := []string{"dark", "light", "auto"}
	return themes[rand.Intn(len(themes))]
}

func randomLanguage() string {
	languages := []string{"pt-BR", "en-US", "es-ES", "fr-FR"}
	return languages[rand.Intn(len(languages))]
}

func randomRoles() []string {
	allRoles := []string{"user", "admin", "moderator", "developer", "viewer"}
	count := rand.Intn(3) + 1 // Entre 1 e 3 roles
	roles := make([]string, count)

	// Garantir que não há duplicatas
	used := make(map[int]bool)
	for i := 0; i < count; i++ {
		var idx int
		for {
			idx = rand.Intn(len(allRoles))
			if !used[idx] {
				used[idx] = true
				break
			}
		}
		roles[i] = allRoles[idx]
	}
	return roles
}

func randomDate() time.Time {
	// Data aleatória nos últimos 365 dias
	daysAgo := rand.Intn(365)
	return time.Now().AddDate(0, 0, -daysAgo)
}

func randomLogMessage() string {
	messages := []string{
		"Request processed successfully",
		"Database query completed in 45ms",
		"Cache miss, fetching from database",
		"User authentication successful",
		"Rate limit exceeded for IP address",
		"Background job started",
		"Email notification sent",
		"File upload completed",
		"API request validation failed",
		"Connection pool exhausted",
		"Retry attempt 3 of 5",
		"Transaction committed successfully",
	}
	return messages[rand.Intn(len(messages))]
}

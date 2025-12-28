package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	var (
		path  = flag.String("path", "migrations", "pasta das migrations")
		cmd   = flag.String("cmd", "up", "comando: up|down|steps|version|force")
		steps = flag.Int("steps", 0, "quantidade de steps (para cmd=steps)")
		force = flag.Int("force", -1, "força versão (para cmd=force)")
	)
	flag.Parse()

	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN não definido")
	}

	m, err := migrate.New(
		"file://"+*path,
		dsn,
	)
	if err != nil {
		log.Fatalf("erro criando migrate: %v", err)
	}

	switch *cmd {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "steps":
		if *steps == 0 {
			log.Fatal("use -steps N com cmd=steps (ex: -cmd=steps -steps=1)")
		}
		err = m.Steps(*steps)
	case "version":
		v, dirty, verr := m.Version()
		if verr == migrate.ErrNilVersion {
			fmt.Println("version: (nil)")
			return
		}
		if verr != nil {
			log.Fatalf("erro pegando versão: %v", verr)
		}
		fmt.Printf("version: %d (dirty=%v)\n", v, dirty)
		return
	case "force":
		if *force < 0 {
			log.Fatal("use -force N com cmd=force (ex: -cmd=force -force=1)")
		}
		err = m.Force(*force)
	default:
		log.Fatalf("cmd inválido: %s", *cmd)
	}

	if err == migrate.ErrNoChange {
		fmt.Println("no change")
		return
	}
	if err != nil {
		log.Fatalf("migração falhou: %v", err)
	}

	fmt.Println("ok")
}

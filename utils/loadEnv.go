package utils

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const strTopPath = "../"
const strEnvFile = ".env"
const strIdeRoot = "go.mod"

var (
	ErrEnvNotFound = errors.New("no .env file found")
	ErrEnvPathLoad = errors.New("não foi possivel localizar o diretorio da aplicação")
	ErrEnvRootPath = errors.New("não foi possivel localizar o diretorio principal")
)

func LoadEnvMem() error {
	// Tenta carregar o .env local (sem falhar caso não exista)
	_ = godotenv.Load(strEnvFile)

	// Tenta também carregar a versão do topo (../.env), se aplicável
	_ = godotenv.Load(strTopPath + strEnvFile)

	// Opcionalmente, tenta carregar do root path detectado
	_ = LoadSysEnv()

	return nil // Nunca retorna erro por .env ausente
}

func LoadSysEnv() error {
	rootPath, err := GetRootPath()

	if err != nil {
		return err
	}

	if err := godotenv.Load(filepath.Join(rootPath, strEnvFile)); err != nil {
		envFilePath := GetSysPath(rootPath)

		if err := godotenv.Load(envFilePath); err != nil {
			return ErrEnvNotFound
		}
	}

	return nil
}

func GetRootPath() (string, error) {
	var rootPath string = ""

	strPathApp, err := os.Getwd()

	if err != nil {
		return rootPath, ErrEnvPathLoad
	}

	rootPath, err = filepath.Abs(filepath.Dir(strPathApp))

	if err != nil {
		return rootPath, ErrEnvRootPath
	}

	return rootPath, nil
}

func GetSysPath(strPath string) string {
	var topDir string = ""

	for {
		if _, err := os.Stat(filepath.Join(strPath, strIdeRoot)); err != nil {
			topDir = filepath.Join(strPath, strEnvFile)
			break
		}

		topPath := filepath.Dir(strPath)
		if topPath == strPath {
			break
		}

		strPath = topPath
	}

	return topDir
}

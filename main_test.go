package monologger

import (
	"os"
	"testing"
)

func TestInfo(t *testing.T) {
	logger, err := New(os.Stdout, "ms-teste-tag ", false)
	if err != nil {
		panic("Não foi possivel inicial filelogger")
	}

	logger.Info("Exemplo de linha de info")
}

func TestError(t *testing.T) {
	logger, err := New(os.Stdout, "ms-teste-tag ", false)
	if err != nil {
		panic("Não foi possivel inicial filelogger")
	}
	logger.Error("Exemplo de error")
}

func TestDebug(t *testing.T) {
	logger, err := New(os.Stdout, "ms-teste-tag ", false)
	if err != nil {
		panic("Não foi possivel inicial filelogger")
	}
	t.Log("Antes de logar debug off")
	logger.Debug("Isso nao deve logar")
	t.Log("Depois de Off, antes de logar debug on")
	logger.SetDebug(true)
	logger.Debug("Isso DEVE logar")
}

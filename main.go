package monologger

import (
	"fmt"
	"io"
	"log"
	"os"
)

//GetArquivoLog criar/concatenar abre arquivo para adição com permissões necessarias
func GetArquivoLog(arquivoLogNome string) (*os.File, error) {
	arquivoLog, err := os.OpenFile(arquivoLogNome, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("Nao foi possivel escrever no arquivo erro1:%s", err.Error())
	}
	return arquivoLog, nil
}

//New - informe destino io.Writer arquivo, stdout, para também escrever em arquivo utilize NewWithFile
func New(destino io.Writer, tag string, debugMode bool) (*Monologger, error) {
	logInstance := &Monologger{
		Logger:    log.New(destino, tag, log.Ldate|log.Lmicroseconds),
		debugMode: debugMode}
	return logInstance, nil
}

//SetDebug configura modo debug
func (l *Monologger) SetDebug(mode bool) {
	l.debugMode = mode
}

//Monologger mantém arquivo e metodos para log
type Monologger struct {
	*log.Logger
	debugMode bool
}

//Warning adiciona nova linha no arquivo de log com rotulo WARNING
//Warning log line "DATETIME TAG WARNING"
func (l *Monologger) Warning(message string, params ...interface{}) {
	l.Println("WARNING ", message, params)
}

//Info adiciona nova linha no log com rotulo INFO
//
//Info log line "DATETIME TAG INFO"
func (l *Monologger) Info(message string, params ...interface{}) {
	l.Println("INFO	", message, params)
}

//Fatal log e finaliza
func (l *Monologger) Fatal(message string, params ...interface{}) {
	l.Logger.Fatal("FATAL ", message, params)
}

//Error adiciona nova linha no arquivo de log
//
//message é inserida no arquivo de log com rotulo ERROR
func (l *Monologger) Error(message string, params ...interface{}) {
	l.Println("ERROR ", message, params)
}

//Debug adiciona nova linha no arquivo de log
//
//TODO: adicionar uma configuração por variavel de ambiente
//que permite ligar/desligar
func (l *Monologger) Debug(message string, params ...interface{}) {
	if l.debugMode {
		l.Println("DEBUG ", message, params)
	}
}
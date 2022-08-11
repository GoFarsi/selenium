package worker

import (
	"bufio"
	"github.com/Ja7ad/selenium/errors"
	"log"
	"os"
)

type Worker struct {
	Address     []string
	ProxyPath   string
	NumOfWorker uint
	proxyList   []string
}

func NewWorker(address []string, proxyPath string, worker uint) *Worker {
	return &Worker{
		Address:     address,
		ProxyPath:   proxyPath,
		NumOfWorker: worker,
	}
}

func (w *Worker) Start() error {
	if err := w.checkAddress(); err != nil {
		return err
	}
	if err := w.setProxyListFromFile(); err != nil {
		return err
	}
	if err := w.checkNumOfWorker(); err != nil {
		return err
	}
	return nil
}

func (w *Worker) checkAddress() error {
	if len(w.Address) == 0 {
		return errors.ERR_ADDRESS_IS_EMPTY
	} else if len(w.Address) > 3 {
		return errors.ERR_ADDRESS_LIST_LIMITED
	}
	return nil
}

func (w *Worker) setProxyListFromFile() error {
	if len(w.ProxyPath) == 0 {
		return errors.ERR_PROXY_PATH_IS_INVALID
	}
	file, err := os.Open(w.ProxyPath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		w.proxyList = append(w.proxyList, scanner.Text())
		if i > 50 {
			break
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	log.Println("the proxy number is limited to 50 and has been loaded")
	return nil
}

func (w *Worker) checkNumOfWorker() error {
	if w.NumOfWorker == 0 {
		w.NumOfWorker = 2
	} else if w.NumOfWorker > 3 {
		return errors.ERR_NUM_OF_WORKER_LIMITED
	}
	return nil
}

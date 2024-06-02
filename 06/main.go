package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	contextCancel()
	channel()
	timeout()
}

func contextCancel() {
	// создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	// запускаем горутину с этим контекстом
	go worker(ctx)
	// ждем 3 секунды
	time.Sleep(3 * time.Second)
	// отменяем контекст
	cancel()
	// ждем еще 2 секунды
	time.Sleep(2 * time.Second)
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// контекст отменен, выходим из горутины
			fmt.Println("\033[31mworker stopped\033[0m")
			return
		default:
			// продолжаем работать
			fmt.Println("\033[31mworker working\033[0m")
			time.Sleep(time.Second)
		}
	}
}

func channel() {
	// создаем канал для передачи сигнала остановки
	stop := make(chan bool)
	// запускаем горутину с этим каналом
	go workerChannel(stop)
	// ждем 3 секунды
	time.Sleep(3 * time.Second)
	// отправляем сигнал остановки в канал
	stop <- true
	// ждем еще 2 секунды
	time.Sleep(2 * time.Second)

}

func workerChannel(stop chan bool) {
	for {
		select {
		case <-stop:
			// получили сигнал остановки, выходим из горутины
			fmt.Println("\033[32mworker stopped\033[0m")
			return
		default:
			// продолжаем работать
			fmt.Println("\033[32mworker working\033[0m")
			time.Sleep(time.Second)
		}
	}
}

func timeout() {
	// создаем канал для получения результата работы
	result := make(chan string)
	fmt.Println("\033[33mworker working\033[0m")
	// запускаем горутину с этим каналом
	go workerTimeout(result)
	// устанавливаем таймаут в 3 секунды
	timeout := time.After(3 * time.Second)
	select {
	case res := <-result:
		// получили результат работы до таймаута
		fmt.Println("\033[33mworker finished:\033[0m", res)
	case <-timeout:
		// время вышло, прерываем горутину
		fmt.Println("\033[33mworker timeout\033[0m")
	}
}

func workerTimeout(result chan string) {
	// имитируем долгую работу
	time.Sleep(1 * time.Second)
	// отправляем результат в канал
	result <- "success"
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*




Так же требуется поддерживать функционал fork/exec-команд


Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).


*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}

	switch parts[0] {

	case "cd":
		if len(parts) < 2 {
			return errors.New("cd: no directory")
		}
		err := os.Chdir(parts[1])
		return err

	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, dir)

	case "echo":
		cmd := exec.Command(parts[0], parts[1:]...)
		output, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, string(output))

	case "kill":
		if len(parts) < 2 {
			return errors.New("kill: no pid")
		}
		cmd := exec.Command("kill", parts[1])
		output, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, string(output))
	case "ps":
		cmd := exec.Command("ps", parts[1:]...)
		output, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, string(output))
	default:

	}
	return nil
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sigChan
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

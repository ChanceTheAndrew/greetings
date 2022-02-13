package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello возвращает приветствие для указанного человека.
func Hello(name string) (string, error) {
	// Если имя не было указано, возвращаем ошибку с сообщением.
	if name == "" {
		return name, errors.New("empty name")
	}

	// Создаем сообщение в произвольном формате.
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprint(randomFormat())
	return message, nil
}

// Hellos возвращает карту,
// которая связывает каждого из названных людей
// с приветственным сообщением.
func Hellos(names []string) (map[string]string, error) {
	// Карта для связывания имен с сообщениями.
	messages := make(map[string]string)
	// Перебираем полученный срез имен, вызываем
	// функция Hello для получения сообщения для каждого имени.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// В карте свяжите полученное сообщение с
		// именем.
		messages[name] = message
	}
	return messages, nil
}

// Init устанавливает начальные значения для переменных,
// используемых в функции.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat возвращает одно
// из набора приветственных сообщений.
// Возвращаемое сообщение выбирается случайным образом.
func randomFormat() string {
	// Срез форматов сообщений.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Возвращаем один из случайно выбранных форматов сообщений.
	return formats[rand.Intn(len(formats))]
}

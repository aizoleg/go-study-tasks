package main

import "fmt"

// WeatherProvider - интерфейс, который ожидается нашей системой для предоставления данных о погоде
type WeatherProvider interface {
	GetWeather() string
}

// ExternalWeatherService - это имитация стороннего сервиса погоды, который имеет свой собственный интерфейс
type ExternalWeatherService struct{}

// RetrieveWeather - метод, который предоставляет данные о погоде в формате, специфичном для ExternalWeatherService
func (ews *ExternalWeatherService) RetrieveWeather() string {
	return "Sunny"
}

// WeatherAdapter - адаптер, который "переводит" методы ExternalWeatherService в тот формат, который ожидается нашей системой
type WeatherAdapter struct {
	Service *ExternalWeatherService // Ссылка на объект стороннего сервиса
}

// GetWeather - реализация метода нашего интерфейса, использующая сторонний сервис
func (wa *WeatherAdapter) GetWeather() string {
	return wa.Service.RetrieveWeather() // Здесь происходит "перевод" метода
}

// ShowWeather - функция, демонстрирующая погоду на основе предоставленного WeatherProvider
func ShowWeather(wp WeatherProvider) {
	fmt.Println("Today's weather is:", wp.GetWeather())
}

func main() {
	// Создание экземпляра стороннего сервиса
	externalService := &ExternalWeatherService{}
	// Создание адаптера для стороннего сервиса
	adapter := &WeatherAdapter{Service: externalService}

	// Использование адаптера вместо прямого вызова стороннего сервиса
	ShowWeather(adapter)
}

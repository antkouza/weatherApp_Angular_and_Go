# 🌦️ Weather App (Angular + Go)

## 📌 Introduction

This project is a simple full-stack Weather Application that allows users to search for weather information by city.

The frontend is built with **Angular**, providing a user-friendly interface, while the backend is implemented in **Go**, acting as a lightweight API server that fetches real-time weather data from the OpenWeather API.\
Moreover, the backend uses a simple **in-memory cache** to cache entries to achieve faster responses for repeated requests and reduced external API calls. Entries expire after a few minutes.

<img width="208" height="315" alt="image" src="https://github.com/user-attachments/assets/ed8759c5-c659-47a4-a372-2f27f0c4d9dd" />

## 🚀 Key Technologies

### Frontend

- Angular
- TypeScript
- HTML / CSS
- Angular Forms & HTTP Client

### Backend

- Go (Golang)
- net/http package
- REST API integration

### External API

- OpenWeatherMap API


## ⚙️ Getting Started

### Prerequisites

Make sure you have the following installed:

- Node.js (v16+ recommended)
- Angular CLI
- Go (v1.18+)
- OpenWeatherMap API Key

## 🛠️ Installation & Run

1. Clone the repository

```bash
git clone https://github.com/antkouza/weatherApp_Angular_and_Go.git
cd weather-app
```

2. Backend Setup (Go)
```bash
cd backend
Create .env file with
API_KEY=your_openweather_api_key
```

Run the Go server

```bash
go build (or go run .)
.\weatherApp.exe
```

The backend will start on:
http://localhost:8080
You can also test the backend alone from browser e.g. http://localhost:8080/weather/London,uk

3. Frontend Setup (Angular)

```bash
cd frontend/weather-app
(npm install if you need to install dependencies)
ng serve (run app)
```

The frontend will start on:
http://localhost:4200.

## 🔄 How It Works
User enters a city in the Angular UI.\
Frontend sends a GET request: http://localhost:8080/weather?city=tokyo to backend.

Go backend:
- Receives the request from the frontend
- Checks the **in-memory cache**:
  - If recent data for the requested city exists → returns cached response
  - If not → proceeds to fetch fresh data
- Calls the OpenWeather API using the API key from the `.env` file
- Processes and transforms the response into a simplified JSON format
- Stores the result in the cache
- Returns JSON weather data to the frontend

Frontend Rendering: Angular receives the JSON response and renders
  - Temperature (°C / °F)
  - Humidity
  - Wind data
  - Weather condition and icon

Backend showcase request with and w/o caching
<img width="1681" height="111" alt="image" src="https://github.com/user-attachments/assets/3024f2d3-7021-42ff-8ee6-84d7fe88758c" />

## 🧪 Testing
Run Angular tests (Karma/Jasmine)

```bash
ng test
```

## 📁 Structure
weatherApp_Angular_and_Go (GitHub Repo)
```bash
├── backend/
│ ├── models.go
│ ├── weather.go
│ └── weather_service.go
├── frontend/
│ └── weather-app/
│ ├── src/
│ └── package.json
└── .gitignore
└── README.md
```

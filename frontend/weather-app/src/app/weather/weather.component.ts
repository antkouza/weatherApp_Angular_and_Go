import { Component } from '@angular/core';
import { WeatherService } from '../weather.service';
import { Weather } from '../models/weather.model';

@Component({
  selector: 'app-weather',
  templateUrl: './weather.component.html',
})
export class WeatherComponent {
  city: string = '';
  weatherData: Weather | null = null;

  constructor(private weatherService: WeatherService) {}

  errorMessage: string = '';
  loading: boolean = false;

  getWeather() {
    const queryCity = this.city?.trim();

    if (!queryCity) {
      this.errorMessage = 'Please enter a city name.';
      this.weatherData = null; // Hide old data
      return;
    }

    this.errorMessage = ''; // Clear previous error
    this.loading = true;
    this.weatherData = null; // Optional: hide data while loading

    this.weatherService.getWeather(queryCity).subscribe({
      next: (data) => {
        this.weatherData = data;
        this.loading = false;
      },
      error: (err) => {
        this.loading = false;
        this.errorMessage = 'City not found. Please try again.';
      },
    });
  }
}

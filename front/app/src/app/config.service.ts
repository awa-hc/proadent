import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ConfigService {
  private config = {
    API_URL: 'http://localhost:8080/',
  };

  get apiUrl(): string {
    return this.config.API_URL;
  }
}

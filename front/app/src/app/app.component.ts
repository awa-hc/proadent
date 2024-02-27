import { HttpClientModule } from '@angular/common/http';
import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { StorageService } from './storage.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, HttpClientModule],
  templateUrl: './app.component.html',
})
export class AppComponent {
  title = 'app';
  constructor(public storage: StorageService) {
    if (typeof window !== 'undefined') {
      const currentTheme = storage.getItem('theme');
      if (currentTheme === 'dark') {
        document.body.classList.add('dark');
      }
    }
  }

  toggleTheme() {
    if (typeof window !== 'undefined') {
      const isDarkMode = document.body.classList.contains('dark');
      if (isDarkMode) {
        document.body.classList.remove('dark');
        this.storage.setItem('theme', 'light');
      } else {
        document.body.classList.add('dark');
        this.storage.setItem('theme', 'dark');
      }
    }
  }
}

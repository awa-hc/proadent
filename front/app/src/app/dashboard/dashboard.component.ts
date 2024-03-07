import { Component, NgZone, OnInit } from '@angular/core';
import { NavbarComponent } from '../components/navbar/navbar.component';
import { StorageService } from '../storage.service';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [NavbarComponent, RouterModule],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.css',
})
export class DashboardComponent {
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

import { HttpClient } from '@angular/common/http';
import { Component, HostListener, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { Route, Router, RouterModule } from '@angular/router';

@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [ReactiveFormsModule, RouterModule],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css',
})
export class ProfileComponent {
  constructor(private router: Router) {}
  isSidebarOpen = false;

  gotoMe() {
    this.router.navigate(['/profile/me']);
  }

  gotoAppointments() {
    this.router.navigate(['/profile/appointment']);
  }

  isMobile(): boolean {
    const win = this.getWindow();
    return win ? win.innerWidth < 1024 : false;
  }

  gotoClinics() {
    this.router.navigate(['/profile/clinics']);
  }

  toggleSidebar() {
    if (window.innerWidth < 1024) {
      this.isSidebarOpen = !this.isSidebarOpen;
    }
  }

  getWindow(): Window | null {
    return typeof window !== 'undefined' ? window : null;
  }
}

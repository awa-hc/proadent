import { HttpClient } from '@angular/common/http';
import { Component, HostListener, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { Route, Router, RouterModule } from '@angular/router';
import { CookieService } from '../../cookie.service';
import { StorageService } from '../../storage.service';

@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [ReactiveFormsModule, RouterModule],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css',
})
export class ProfileComponent implements OnInit {
  constructor(
    private router: Router,
    private _cookieService: CookieService,
    private storage: StorageService
  ) {}
  isSidebarOpen = false;

  ngOnInit(): void {
    const tokenexpiration = this.storage.getItem('AuthExpiration');
    const currentTime = new Date().getTime();

    if (tokenexpiration && currentTime > Number(tokenexpiration)) {
      this.storage.removeItem('AuthExpiration');
      this.storage.removeItem('Auth');
      this._cookieService.deleteCookie('Auth');
    }
    if (
      this._cookieService.getCookie('Auth') == null ||
      this.storage.getItem('Auth') == null
    ) {
      this.router.navigate(['/login']);
    }
  }

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

import { Component, NgZone, OnInit } from '@angular/core';
import { LocalStorageService } from '../local-storage.service';
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatTooltipModule } from '@angular/material/tooltip';
import { CookieService } from '../cookie.service';

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [MatSnackBarModule, MatTooltipModule],
  templateUrl: './navbar.component.html',
})
export class NavbarComponent implements OnInit {
  TimeNow: string = '';
  UserName: string = '';
  UserProfile: boolean = false;
  User: any = {};
  token: string | null = '';
  isLoading: boolean = false;

  constructor(
    private ngZone: NgZone,
    private localStorageService: LocalStorageService,
    private router: Router,
    private snackBar: MatSnackBar,
    private cookieService: CookieService
  ) {}
  ngOnInit(): void {
    this.ngZone.runOutsideAngular(() => {
      setInterval(() => {
        const now = new Date();
        this.TimeNow =
          now.toLocaleDateString() + ' ' + now.toLocaleTimeString();
        this.ngZone.run(() => {});
      }, 1000);
    });
    this.fetchUser();
  }

  async fetchUser() {
    this.token = this.localStorageService.getItem('user');
    if (this.token !== null || this.cookieService.getCookie('Auth')) {
      this.isLoading = true;
      try {
        const response = await fetch('http://localhost:8080/user/logged', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${this.token}`,
          },
          credentials: 'include',
        });
        const data = await response.json();
        this.isLoading = false;
        this.User = data.user;
        this.UserName = this.User.FullName;
      } catch (error) {
        console.error('Error fetching user:', error);
      } finally {
        this.isLoading = false;
      }
    } else {
      this.isLoading = false;
    }
  }

  async logout() {
    try {
      const response = await fetch('http://localhost:8080/auth/logout', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${this.token}`,
        },
        credentials: 'include',
      });
      const data = await response.json();
      console.log(data);
      if (data.message) {
        this.snackBar.open(data.message, 'Close', {
          duration: 3000,
          verticalPosition: 'top',
        });
        this.localStorageService.removeItem('user');

        this.cookieService.deleteCookie('Auth');
        this.router.navigate(['/']);
      } else {
        this.snackBar.open('Error logging out', 'Close', {
          duration: 3000,
        });
      }
    } catch (error) {
      console.error('Error logging out:', error);
    }
  }

  toggleProfile() {
    this.UserProfile = !this.UserProfile;
  }
}

import { Component, OnInit, OnDestroy, NgZone } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { DataService } from '../../data.service';
import { StorageService } from '../../storage.service';
import { Router } from '@angular/router';
import { translateYAnimation } from '../Animations/animations';

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [MatButtonModule, MatTooltipModule],
  providers: [DataService],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.css',
  animations: [translateYAnimation],
})
export class NavbarComponent implements OnInit {
  time = { dayname: '', day: '', time: '', month: '' };
  user: any = [];
  userProfile: boolean = false;
  mobileMenu: boolean = false;

  constructor(
    private ngZone: NgZone,
    private _router: Router,
    private _dataService: DataService,
    private _storageService: StorageService
  ) {}

  ngOnInit(): void {
    this.ngZone.runOutsideAngular(() => {
      this.updateTime();
      setInterval(() => {
        this.updateTime();
        this.ngZone.run(() => {});
      }, 6000);
    });

    this.getUserDetails();
  }

  updateTime(): void {
    const now = new Date();
    this.time.dayname = now.toLocaleDateString('es', { weekday: 'long' });
    this.time.day = now.toLocaleDateString('es', { day: '2-digit' });
    this.time.time = now.toLocaleTimeString('es', {
      hour: '2-digit',
      minute: '2-digit',
    });
    this.time.month = now.toLocaleDateString('es', { month: 'long' });
  }

  ngOnDestroy(): void {
    this.ngZone.run(() => {});
  }

  getUserDetails(): void {
    this._dataService.getuserinfo().subscribe((data) => {
      // console.log(data);
      this.user = data;
    });
  }
  toggleProfile() {
    this.userProfile = !this.userProfile;
  }
  logout() {
    this._storageService.removeItem('token');
    this._router.navigate(['/']);
  }

  toggleMenu() {
    this.mobileMenu = !this.mobileMenu;
  }
}

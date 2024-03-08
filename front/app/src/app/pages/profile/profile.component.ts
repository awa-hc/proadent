import { Component, OnInit } from '@angular/core';
import { StorageService } from '../../storage.service';
import { Router } from '@angular/router';
import { DataService } from '../../data.service';
import { CommonModule } from '@angular/common';
import { BrowserModule } from '@angular/platform-browser';

@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [],
  providers: [DataService, CommonModule, BrowserModule],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css',
})
export class ProfileComponent implements OnInit {
  email: string = '';
  fullname: string = '';
  ci: string = '';
  birthdate: string = '';
  appointments: any[] = [];
  phone: string = '';
  createdat: string = '';
  updatedat: string = '';
  roledescription: string = '';

  constructor(
    private dataService: DataService,
    private storageService: StorageService,
    private router: Router
  ) {}

  ngOnInit() {
    if (!this.storageService.getItem('token')) {
      this.router.navigate(['/login']);
    }
    this.dataService.getuserinfo().subscribe((response) => {
      console.log(response);
      console.log(response.appointments);
      this.email = response.email;
      this.fullname = response.fullName;
      this.ci = response.ci;
      this.birthdate = response.birthDay;
      this.phone = response.phone;
      this.createdat = response.createdAt;
      this.updatedat = response.updatedAt;
      this.roledescription = response.role.description;
      this.appointments = response.appointments;
    });
  }

  getAppointmentStyle(status: string): string {
    switch (status) {
      case 'pending':
        return 'background-color: yellow; color: black;';
      case 'completed':
        return 'background-color: green; color: white;';
      case 'cancelled':
        return 'background-color: red; color: white;';
      case 'confirmed':
        return 'background-color: blue; color: white;';
      default:
        return '';
    }
  }
}

import { Component, HostListener, OnInit } from '@angular/core';
import { DataService } from '../../../data.service';
import { Router } from '@angular/router';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { provideNativeDateAdapter } from '@angular/material/core';

@Component({
  selector: 'app-all',
  standalone: true,
  imports: [MatDatepickerModule],
  providers: [DataService, provideNativeDateAdapter()],
  templateUrl: './all.component.html',
  styleUrl: './all.component.css',
})
export class AllComponent implements OnInit {
  data: any[] = [];
  dates: Date[] = [];

  ngOnInit() {
    this.getAllappointments();
  }

  constructor(private _dataService: DataService, private _router: Router) {}

  getDates() {
    this.data.forEach((appointment) => {
      const alldate = appointment.date.split(' ');
      const parts = alldate[0].split('/');
      const newdate = parts[2] + '/' + parts[1] + '/' + parts[0];
      const date = new Date(newdate);
      this.dates.push(date);
    });
    console.log(this.dates);
  }

  getAllappointments() {
    this._dataService.getallappointments().subscribe((data) => {
      this.data = data;
      this.getDates();
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

  changelocation(code: string) {
    console.log(code);
    window.location.href = 'dashboard/appointments/edit/' + code;
    console.log('asdsad');
  }
}

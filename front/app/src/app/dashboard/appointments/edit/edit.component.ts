import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { DataService } from '../../../data.service';
import { Subject, debounceTime, distinctUntilChanged, switchMap } from 'rxjs';

@Component({
  selector: 'app-edit',
  standalone: true,
  providers: [DataService, Router],
  templateUrl: './edit.component.html',
  styleUrl: './edit.component.css',
})
export class EditComponent implements OnInit {
  appointmentcode: string = '';
  date: string = '';
  createdAt: string = '';
  updatedAt: string = '';
  reason: string = '';
  status: string = '';
  userBirthDay: string = '';
  userCI: string = '';
  userEmail: string = '';
  userName: string = '';
  userPhone: string = '';
  isLoading: boolean = false;
  appointmentssuggest: string[] = [];

  constructor(
    private route: ActivatedRoute,
    private _dataService: DataService,
    private _router: Router
  ) {
    this.searchAppointment
      .pipe(
        debounceTime(300),
        distinctUntilChanged(),
        switchMap((term: string) =>
          this._dataService.getappointmentbycode(term)
        )
      )
      .subscribe((suggest: string[]) => {
        this.appointmentssuggest = suggest;
      });
  }

  ngOnInit(): void {
    this.route.params.subscribe((params) => {
      this.appointmentcode = params['code'];
      this.getappointment(this.appointmentcode);
    });
  }

  getappointment(code: string): void {
    this._dataService.getappointmentbycode(code).subscribe((data) => {
      this.appointmentcode = data.code;
      this.createdAt = data.createdAt;
      this.date = data.date;
      this.reason = data.reason;
      this.status = data.status;
      this.userBirthDay = data.userBirthDay;
      this.updatedAt = data.updatedAt;
      this.userName = data.userName;
      this.userPhone = data.userPhone;
      this.userCI = data.userCI;
      this.userEmail = data.userEmail;
      console.log(data);
      this.isLoading = false;
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
  changestatusappointment(code: string, status: string) {
    this.isLoading = true;
    const data = {
      status: status,
    };
    this._dataService.udpateappointmentstatus(code, data).subscribe((data) => {
      console.log(data);
      if (data.message) {
        this.isLoading = false;
        this._router
          .navigateByUrl('/', { skipLocationChange: true })
          .then(() => {
            window.location.reload();
          });
      }
    });
  }

  private searchAppointment = new Subject<string>();
}

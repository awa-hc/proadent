import {
  Component,
  CUSTOM_ELEMENTS_SCHEMA,
  NgModule,
  OnInit,
} from '@angular/core';
import { SchedulerComponent } from '../../../components/scheduler/scheduler.component';
import { HttpClient } from '@angular/common/http';
import { format, parseISO } from 'date-fns';
import { StorageService } from '../../../storage.service';

@NgModule({
  declarations: [SchedulerComponent],
  exports: [SchedulerComponent],
})
export class SchedulerModule {}

@Component({
  selector: 'app-profile-appointment',
  standalone: true,
  templateUrl: './profile-appointment.component.html',
  styleUrl: './profile-appointment.component.css',
  imports: [SchedulerModule],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
})
export class ProfileAppointmentComponent implements OnInit {
  // appointments: any[] = [
  //   {
  //     id: 1,
  //     text: 'Cita de prueba',
  //     start_date: '2024-09-29 10:00',
  //     end_date: '2024-09-29 10:30',
  //   },
  //   {
  //     id: 2,
  //     text: 'Cita de prueba 2',
  //     start_date: '2024-09-29 11:00',
  //     end_date: '2024-09-29 11:30',
  //   },
  // ];

  constructor(private http: HttpClient, private storage: StorageService) {}
  appointments: any[] = [];
  //testing
  doctors = [
    { ci: '23345678', name: 'Dr. Smith' },
    { ci: '34567890', name: 'Dr. Johnson' },
  ];
  selectedDoctor = this.doctors[0].ci;
  reason: string = '';
  dateTime: string = '2024-09-06T12:00:00Z';
  patientCi: string = '';

  ngOnInit(): void {
    this.loadAppointments();
    this.patientCi = this.storage.getItem('ci') || '';
  }

  loadAppointments() {
    const ci = this.storage.getItem('ci');
    if (ci) {
      this.http
        .get<any[]>('http://localhost:8080/appointment/patient/' + ci, {
          withCredentials: true,
        })
        .subscribe(
          (response) => {
            this.appointments = this.transformEvents(response);
          },
          (error) => {
            console.error('Error al obtener las citas', error);
          }
        );
    }
  }

  formateDate(dateISO: string): string {
    const date = parseISO(dateISO);
    return format(date, 'yyyy-MM-dd HH:mm');
  }
  calculateEndDate(dateISO: string): string {
    const start_date = parseISO(dateISO);
    const end_date = new Date(start_date.getTime() + 30 * 60000);
    return format(end_date, 'yyyy-MM-dd HH:mm');
  }

  transformEvents(data: any[]): any[] {
    const transformed = data.map((event) => ({
      id: event.ID,
      text: event.reason,
      start_date: this.formateDate(event.date_time),
      end_date: this.calculateEndDate(event.date_time),
    }));

    return transformed;
  }
}

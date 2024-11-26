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
import { MatDividerModule } from '@angular/material/divider';
import {
  FormBuilder,
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  ValidatorFn,
  Validators,
} from '@angular/forms';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';

import { noPastDateValidator } from '../../../utils/forms/pastimevalidator';

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
  imports: [
    SchedulerModule,
    MatDividerModule,
    ReactiveFormsModule,
    MatDatepickerModule,
    MatNativeDateModule,
  ],
  providers: [MatNativeDateModule],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
})
export class ProfileAppointmentComponent implements OnInit {
  constructor(
    private http: HttpClient,
    private storage: StorageService,
    private fb: FormBuilder
  ) {
    this.createAppointment = this.fb.group({
      reason: ['', [Validators.required, Validators.minLength(5)]],
      start_date: [Validators.required, noPastDateValidator()],
      patient_ci: new FormControl({
        value: this.storage.getItem('ci'),
        disabled: true,
      }),
      doctor_ci: [''],
      start_time: ['', Validators.required],
      requestedBy: [this.storage.getItem('ci') || '', Validators.required],
    });
  }
  createAppointment: FormGroup;
  appointments: any[] = [];
  showForm: boolean = false;
  createAppointmentDate = '';
  reason: string = '';
  dateTime: string = '2024-09-06T12:00:00Z';
  patientCi: string = '';
  availableHours: string[] = [];

  ngOnInit(): void {
    this.loadAppointments();
    this.patientCi = this.storage.getItem('ci') || '';
    this.createAppointment.get('start_date')?.valueChanges.subscribe((date) => {
      this.updateAvailableHours(date);
    });
  }

  loadAppointments() {
    const ci = this.storage.getItem('ci');

    if (ci) {
      this.http
        .get<any[]>('http://localhost:8080/userappointments/user/' + ci, {
          withCredentials: true,
        })
        .subscribe(
          (response: any) => {
            console.log('Citas obtenidas', response);
            this.appointments = this.transformEvents(response.appointments);
          },
          (error: any) => {
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

  onEmptyDateSelected(date: string) {
    this.showForm = true;
    this.createAppointmentDate = date;

    const [selectedDate, selectedTime] = date.split(' ');

    const dateObject = new Date(selectedDate);

    dateObject.setDate(dateObject.getDate() + 1);

    const adjustedDate = dateObject.toISOString().split('T')[0];

    this.createAppointment.patchValue({
      start_date: adjustedDate,
      start_time: selectedTime,
    });
  }

  closeForm() {
    this.showForm = false;
  }

  createAppointmentPOST() {
    console.log('asdasd');
    console.log(this.createAppointment.valid);
    if (this.createAppointment.valid) {
      const formValues = this.createAppointment.value;
      const appointmentData = {
        doctor_ci: '23345678', // Aquí podrías tener una variable o valor dinámico
        patient_ci: formValues.patient_ci,
        date_time: `${formValues.start_date}T${formValues.start_time}:00Z`, // Formato UTC
        reason: formValues.reason,
        requested_by: formValues.requested_by,
        type: 'presential',
      };
                    
      this.http
        .post('http://localhost:8080/appointment/create', appointmentData, {
          withCredentials: true,
        })
        .subscribe(
          (response: any) => {
            console.log('Cita creada', response);
            this.loadAppointments();
            this.showForm = false;
          },
          (error: any) => {
            console.error('Error al crear la cita', error);
          }
        );
    } else {
      console.error('Formulario inválido');
    }
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

  // Función para actualizar las horas basadas en el día seleccionado
  updateAvailableHours(date: string): void {
    const day = new Date(date).getDay(); // Obtenemos el día (0 = Domingo, 6 = Sábado)

    if (day === 6) {
      // Sábado: 9:00 AM - 1:00 PM
      this.availableHours = this.generateTimeSlots(9, 13);
    } else if (day !== 0) {
      // Lunes a Viernes: 9:00 AM - 8:00 PM
      this.availableHours = this.generateTimeSlots(9, 20);
    } else {
      // Domingo no permitimos citas
      this.availableHours = [];
    }
  }

  // Función para generar el rango de horas permitidas
  generateTimeSlots(startHour: number, endHour: number): string[] {
    const slots: string[] = [];
    for (let hour = startHour; hour < endHour; hour++) {
      slots.push(`${this.formatTime(hour, 0)}`);
      slots.push(`${this.formatTime(hour, 30)}`);
    }
    slots.push(`${this.formatTime(endHour, 0)}`); // Añadir la última hora sin minutos
    return slots;
  }

  formatTime(hour: number, minutes: number): string {
    const formattedHour = hour < 10 ? `0${hour}` : hour;
    const formattedMinutes = minutes < 10 ? `0${minutes}` : minutes;
    return `${formattedHour}:${formattedMinutes}`;
  }

  dateFilter = (date: Date | null): boolean => {
    if (date === null) {
      return false;
    }
    const day = date.getDay();
    const now = new Date();
    now.setHours(0, 0, 0, 0); // Pone la hora del día actual a las 00:00

    // Excluir domingos (day === 0) y fechas anteriores al día actual
    return date >= now && day !== 0;
  };
}

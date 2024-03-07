import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { provideNativeDateAdapter } from '@angular/material/core';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import {
  AbstractControl,
  FormBuilder,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatStepperModule } from '@angular/material/stepper';
import { STEPPER_GLOBAL_OPTIONS } from '@angular/cdk/stepper';
import { MatSelectModule } from '@angular/material/select';
import { DataService } from '../../data.service';
import GetUserIdFromToken from '../../utils/token';
import { StorageService } from '../../storage.service';

@Component({
  selector: 'app-home',
  standalone: true,
  providers: [
    provideNativeDateAdapter(),
    {
      provide: STEPPER_GLOBAL_OPTIONS,
      useValue: { showError: true },
    },
    [DataService],
  ],

  imports: [
    MatFormFieldModule,
    MatInputModule,
    MatDatepickerModule,
    MatButtonModule,
    MatCardModule,
    ReactiveFormsModule,
    MatStepperModule,
    FormsModule,
    MatSelectModule,
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
})
export class HomeComponent implements OnInit {
  firstFormGroup!: FormGroup;
  secondFormGroup!: FormGroup;
  thirdFormGroup!: FormGroup;
  userlogged: boolean = false;

  hours: string[] = [
    'MAÑANA',
    '09:00',
    '09:30',
    '10:00',
    '11:00',
    '11:30',
    '12:00',
    'TARDE',
    '15:00',
    '15:30',
    '16:00',
    '16:30',
    '17:00',
    '17:30',
    '18:00',
    '18:30',
  ];

  dayFilter = (d: Date | null): boolean => {
    const day = (d || new Date()).getDay();

    const today = new Date();
    today.setHours(0, 0, 0, 0); // Ajustar las horas a 0 para comparar solo la fecha sin tener en cuenta la hora

    if (!d) {
      return false;
    }

    // Previene la selección de fechas anteriores a hoy y de los días domingo (0).
    return d >= today && day !== 0;
  };

  constructor(
    private fb: FormBuilder,
    private dataService: DataService,
    private storageService: StorageService
  ) {
    this.firstFormGroup = this.fb.group({
      firstCtrl: ['', Validators.required],
    });
    this.secondFormGroup = this.fb.group({
      secondCtrl: ['', [Validators.required, this.timeRangeValidator]],
    });

    this.thirdFormGroup = this.fb.group({
      thirdCtrl: ['', [Validators.required, Validators.maxLength(100)]],
    });
  }

  ngOnInit(): void {
    if (this.storageService.getItem('token')) {
      this.userlogged = true;
    }
  }
  timeRangeValidator(
    control: AbstractControl
  ): { [key: string]: boolean } | null {
    const invalidValues = ['MAÑANA', 'TARDE'];

    if (control.value && invalidValues.includes(control.value.toUpperCase())) {
      return { invalidTime: true };
    }
    return null;
  }

  SubmitData(): void {
    if (!this.userlogged) {
      window.location.href = '/login';
    }

    if (
      this.firstFormGroup.invalid ||
      this.secondFormGroup.invalid ||
      this.thirdFormGroup.invalid
    ) {
      return;
    }

    const selectedDate: Date = this.firstFormGroup.value.firstCtrl;
    const year = selectedDate.getFullYear();
    const month = selectedDate.getMonth();
    const day = selectedDate.getDate();
    const time = this.secondFormGroup.value.secondCtrl;
    const [hours, minutes] = time.split(':').map(Number);
    const dateTime = new Date(year, month, day, hours, minutes);
    console.log(GetUserIdFromToken(this.storageService.getItem('token') || ''));

    const data = {
      userId: GetUserIdFromToken(this.storageService.getItem('token') || ''),
      date: dateTime.toISOString(),
      reason: this.thirdFormGroup.value.thirdCtrl,
    };
    this.dataService.createappointment(data).subscribe((response) => {
      console.log(response);
    });
  }
}

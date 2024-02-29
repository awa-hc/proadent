import { ChangeDetectorRef, Component } from '@angular/core';
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
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';

@Component({
  selector: 'app-home',
  standalone: true,
  providers: [provideNativeDateAdapter()],
  imports: [
    MatFormFieldModule,
    MatInputModule,
    MatDatepickerModule,
    MatButtonModule,
    MatCardModule,
    ReactiveFormsModule,
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
})
export class HomeComponent {
  timeForm!: FormGroup;

  constructor(private fb: FormBuilder) {}

  ngOnInit(): void {
    this.timeForm = this.fb.group({
      time: ['', [Validators.required, this.timeRangeValidator]],
      reason: ['', [Validators.required]],
    });
  }

  timeRangeValidator(control: AbstractControl): { [key: string]: any } | null {
    const value = control.value;
    if (value) {
      const hours = +value.split(':')[0];
      const minutes = +value.split(':')[1];
      const totalMinutes = hours * 60 + minutes;

      const isMorningValid =
        totalMinutes >= 9 * 60 && totalMinutes <= 12 * 60 + 30;
      const isAfternoonValid =
        totalMinutes >= 15 * 60 && totalMinutes <= 19 * 60;

      if (!isMorningValid && !isAfternoonValid) {
        return { timeRangeInvalid: true };
      }
    }
    return null;
  }

  roundTime(event: any): void {
    let [hour, minute] = event.target.value
      .split(':')
      .map((val: string) => parseInt(val, 10));

    if (minute >= 15 && minute < 45) {
      minute = 30;
    } else if (minute < 15) {
      minute = 0;
    } else {
      minute = 0;
      hour += 1;
    }
    const roundedTime = `${hour.toString().padStart(2, '0')}:${minute
      .toString()
      .padStart(2, '0')}`;
    this.timeForm.get('time')?.setValue(roundedTime, { emitEvent: false });
  }
}

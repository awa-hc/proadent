import { Component, HostListener } from '@angular/core';
import { MatTooltip, MatTooltipModule } from '@angular/material/tooltip';
import { Router, RouterModule } from '@angular/router';
import { DataService } from '../../data.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-appointments',
  standalone: true,
  imports: [RouterModule, MatTooltipModule],
  providers: [DataService],
  templateUrl: './appointments.component.html',
  styleUrl: './appointments.component.css',
})
export class AppointmentsComponent {
  showError: boolean = false;
  constructor(
    private _dataService: DataService,
    private _snackBar: MatSnackBar,
    private _router: Router
  ) {}
  @HostListener('document:keydown', ['$event'])
  handleKeyboardEvent(event: KeyboardEvent) {
    if (event.key === 'Enter') {
      this.searchappointment();
    }
  }
  searchappointment(): void {
    let code = (document.getElementById('code') as HTMLInputElement).value;
    this._dataService.getappointmentbycode(code).subscribe((data) => {
      if (data.error) {
        this.showError = true;
        this._snackBar.open('Cita no encontrada', 'Cerrar', {
          duration: 3000,
          verticalPosition: 'top',
        });
      } else {
        this._router.navigate(['dashboard/appointments/edit/' + code]);
      }
    });
  }
}

import { Component, Inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { formatUtcToLocal } from '../../../../utils/date-utils';

@Component({
  selector: 'app-profile-clinics-dialog',
  standalone: true,
  imports: [MatDialogModule, MatButtonModule],
  templateUrl: './profile-clinics-dialog.component.html',
})
export class ProfileClinicsDialogComponent {
  formattedData: any;

  constructor(@Inject(MAT_DIALOG_DATA) public data: any) {}
}

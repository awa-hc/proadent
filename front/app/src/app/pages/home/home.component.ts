import { Component } from '@angular/core';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatCardModule } from '@angular/material/card';
import { provideNativeDateAdapter } from '@angular/material/core';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [MatCardModule, MatDatepickerModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
  providers: [provideNativeDateAdapter()],
})
export class HomeComponent {
  selected!: Date | null;

  
}

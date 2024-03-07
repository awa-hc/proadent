import { Component, OnInit, OnDestroy, NgZone } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [MatButtonModule, MatTooltipModule],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.css',
})
export class NavbarComponent implements OnInit {
  time = { dayname: '', day: '', time: '', month: '' };

  constructor(private ngZone: NgZone) {}

  ngOnInit(): void {
    this.ngZone.runOutsideAngular(() => {
      this.updateTime();
      setInterval(() => {
        this.updateTime();
        this.ngZone.run(() => {});
      }, 6000);
    });
  }

  updateTime(): void {
    const now = new Date();
    this.time.dayname = now.toLocaleDateString('es', { weekday: 'long' });

    this.time.day = now.toLocaleDateString('es', { day: '2-digit' });

    this.time.time = now.toLocaleTimeString('es', {
      hour: '2-digit',
      minute: '2-digit',
    });
    this.time.month = now.toLocaleDateString('es', { month: 'long' });
  }
}

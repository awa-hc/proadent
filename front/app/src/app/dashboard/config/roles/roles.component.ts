import { Component } from '@angular/core';
import { DataService } from '../../../data.service';

@Component({
  selector: 'app-roles',
  standalone: true,
  imports: [],
  providers: [DataService],
  templateUrl: './roles.component.html',
  styleUrl: './roles.component.css',
})
export class RolesComponent {
  roles: any = [];

  constructor(private _dataService: DataService) {
    this.getroles();
  }

  getroles(): void {
    this._dataService.getroles().subscribe((data) => {
      console.log(data);
      this.roles = data;
    });
  }
}

import { Component } from '@angular/core';
import { DataService } from '../../../data.service';

@Component({
  selector: 'app-users',
  standalone: true,
  imports: [],
  providers: [DataService],
  templateUrl: './users.component.html',
  styleUrl: './users.component.css',
})
export class UsersComponent {
  users: any[] = [];

  constructor(private _dataService: DataService) {
    this.getUsers();
  }

  getUsers(): void {
    this._dataService.getUsers().subscribe((data) => {
      this.users = data;
      console.log(data);
    });
  }
}

import { Component, OnInit } from '@angular/core';
import { StorageService } from '../../storage.service';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [],
  templateUrl: './home.component.html',
})
export class HomeComponent implements OnInit {
  constructor(private _storage: StorageService) {}

  ngOnInit(): void {

    

  }
}

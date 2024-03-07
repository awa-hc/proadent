import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, from } from 'rxjs';
import { StorageService } from './storage.service';
import { response } from 'express';

@Injectable({
  providedIn: 'root',
})
export class DataService {
  private url = 'http://localhost:5062/';
  constructor(
    private http: HttpClient,
    private storageService: StorageService
  ) {}

  register(data: any): Observable<any> {
    let response = fetch(this.url + 'user/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .catch((error) => {
        return error;
      });
    return from(response);
  }

  login(data: any): Observable<any> {
    let response = fetch(this.url + 'Auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .catch((error) => {
        return error;
      });
    return from(response);
  }

  createappointment(data: any): Observable<any> {
    let response = fetch(this.url + 'Appointment/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + this.storageService.getItem('token'),
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .catch((error) => {
        return error;
      });
    return from(response);
  }

  getuserinfo(): Observable<any> {
    let response = fetch(this.url + 'user/me', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + this.storageService.getItem('token'),
      },
    })
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .catch((error) => {
        return error;
      });
    return from(response);
  }
}

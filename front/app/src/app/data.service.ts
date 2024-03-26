import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, from } from 'rxjs';
import { StorageService } from './storage.service';
import { response } from 'express';

@Injectable({
  providedIn: 'root',
})
export class DataService {
  private url = 'https://backasp.fly.dev/';
  private urlservices = `http://localhost:8080/`;
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

  getallappointments(): Observable<any> {
    let response = fetch(this.url + 'appointment', {
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
  getappointmentbycode(code: string): Observable<any> {
    let response = fetch(this.url + 'appointment/' + code, {
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
  udpateappointmentstatus(code: string, data: any): Observable<any> {
    let response = fetch(this.url + 'appointment/status/' + code, {
      method: 'PUT',
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
      .then((error) => {
        return error;
      });
    return from(response);
  }

  getUsers(): Observable<any> {
    let response = fetch(this.url + 'user/all', {
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
      .then((error) => {
        return error;
      });
    return from(response);
  }

  contactForm(data: any): Observable<any> {
    let response = fetch(this.urlservices + `contact-form`, {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .then((error) => {
        return error;
      });
    return from(response);
  }

  getroles(): Observable<any> {
    let response = fetch(this.url + 'Role', {
      method: 'GET',
      headers: {
        'Content-type': 'application/json',
        Authorization: 'Bearer ' + this.storageService.getItem('token'),
      },
    })
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .then((error) => {
        return error;
      });
    return from(response);
  }
}

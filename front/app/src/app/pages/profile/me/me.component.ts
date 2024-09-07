import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { StorageService } from '../../../storage.service';

@Component({
  selector: 'app-me',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './me.component.html',
  styleUrl: './me.component.css',
})
export class MeComponent implements OnInit {
  changePasswordForm: FormGroup;

  constructor(
    private http: HttpClient,
    private router: Router,
    private fb: FormBuilder
  ) {
    this.changePasswordForm = this.fb.group({
      currentPassword: '',
      newPassword: '',
      confirmPassword: '',
    });
  }
  window: any;
  profile: any = {};
  isSidebarOpen = false;

  ngOnInit(): void {
    this.getProfile();
  }

  changePassword() {
    console.log('Cambiando contraseña ...');
  }

  getProfile() {
    this.http
      .get('http://localhost:8080/user/me', { withCredentials: true })
      .subscribe(
        (response) => {
          this.profile = response;
          localStorage.setItem('ci', this.profile.ci);
          if (this.profile.birthdate) {
            this.profile.birthdate = new Date(
              this.profile.birthdate
            ).toLocaleDateString();
          }
        },
        (error) => {
          console.error('Error al obtener el perfil', error);
        }
      );
  }

  goToAppointments() {
    this.router.navigate(['/profile/appointment']);
  }
}

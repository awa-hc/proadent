import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { StorageService } from '../../storage.service';
import { CookieService } from '../../cookie.service';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [ReactiveFormsModule],
  providers: [HttpClient],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  loginForm: FormGroup;

  constructor(
    private http: HttpClient,
    private fb: FormBuilder,
    private router: Router,
    private storage: StorageService,
    private _cookieService: CookieService
  ) {
    this.loginForm = this.fb.group({
      email: '',
      password: '',
    });
  }

  ngOnInit(): void {
    if (this.storage.getItem('Auth') || this._cookieService.getCookie('Auth')) {
      this.router.navigate(['/profile']);
    }

    const tokenexpiration = this.storage.getItem('AuthExpiration');
    const currentTime = new Date().getTime();

    if (tokenexpiration && currentTime > Number(tokenexpiration)) {
      this.storage.removeItem('AuthExpiration');
      this.storage.removeItem('Auth');
      this._cookieService.deleteCookie('Auth');

    }
  }

  onSubmit() {
    const loginData = this.loginForm.value;
    this.http
      .post<{ token: string }>(
        'http://localhost:8080/auth/login/email',
        loginData
      )
      .subscribe(
        (response: any) => {
          console.log('Login exitoso', response);
          // Guarda el token en una cookie llamada 'auth'
          const expires = new Date();
          expires.setTime(expires.getTime() + 24 * 60 * 60 * 1000);
          document.cookie = `Auth=${
            response.token
          }; expires=${expires.toUTCString()}; path=/; secure; samesite=strict`;
          this.storage.setItem('Auth', response.token);
          this.storage.setItem('AuthExpiration', expires.getTime().toString());

          this.router.navigate(['/profile']);
        },
        (error: any) => {
          console.error('Error en el login', error);
        }
      );
  }
}

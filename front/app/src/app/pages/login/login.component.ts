import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [ReactiveFormsModule],
  providers: [HttpClient],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent {
  loginForm: FormGroup;

  constructor(
    private http: HttpClient,
    private fb: FormBuilder,
    private router: Router
  ) {
    this.loginForm = this.fb.group({
      email: '',
      password: '',
    });
  }

  onSubmit() {
    const loginData = this.loginForm.value;
    this.http
      .post<{ token: string }>(
        'http://localhost:8080/auth/login/email',
        loginData
      )
      .subscribe(
        (response) => {
          console.log('Login exitoso', response);

          // Guarda el token en una cookie llamada 'auth'
          const expires = new Date();
          expires.setTime(expires.getTime() + 24 * 60 * 60 * 1000); // 7 días de expiración

          document.cookie = `Auth=${
            response.token
          }; expires=${expires.toUTCString()}; path=/; secure; samesite=strict`;
          this.router.navigate(['/profile']);
        },
        (error) => {
          console.error('Error en el login', error);
        }
      );
  }
}

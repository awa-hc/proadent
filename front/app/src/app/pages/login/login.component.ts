import { Component, OnInit } from '@angular/core';
import {
  FormBuilder,
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatCard, MatCardModule } from '@angular/material/card';
import { MatFormField, MatLabel } from '@angular/material/form-field';
import { MatSnackBar } from '@angular/material/snack-bar';
import { DataService } from '../../data.service';
import { StorageService } from '../../storage.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatLabel,
    MatCardModule,
    MatCard,
    MatFormField,
  ],
  providers: [DataService],

  templateUrl: './login.component.html',
  styleUrl: './login.component.css',
})
export class LoginComponent implements OnInit {
  loginForm = this.fb.group({
    email: ['', [Validators.required, Validators.email]],
    password: ['', Validators.required],
  });

  constructor(
    private fb: FormBuilder,
    private _snackBar: MatSnackBar,
    private dataService: DataService,
    private storageService: StorageService,
    private router: Router
  ) {}

  ngOnInit(): void {
    if (this.storageService.getItem('token')) {
      this.router.navigate(['/']);
    }
  }

  onSubmit(): void {
    console.log('asdasd');
    if (this.loginForm.get('email')?.invalid) {
      this._snackBar.open('El campo email es requerido', 'Aceptar', {
        duration: 2000,
      });
    }
    if (this.loginForm.get('password')?.invalid) {
      this._snackBar.open('El campo contraseña es requerido', 'Aceptar', {
        duration: 2000,
      });
    }

    if (this.loginForm.valid) {
      console.log(this.loginForm.value);
      this.dataService.login(this.loginForm.value).subscribe((response) => {
        console.log(response);
        localStorage.setItem('token', response.token);
      });
    }
  }
}

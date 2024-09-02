import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { DataService } from '../../data.service';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [ReactiveFormsModule],
  providers: [DataService],
  templateUrl: './register.component.html',
  styleUrl: './register.component.css',
})
export class RegisterComponent implements OnInit {
  registerForm!: FormGroup;

  constructor(
    private form: FormBuilder,
    private _snackBar: MatSnackBar,
    private dataService: DataService
  ) {}

  ngOnInit(): void {
    this.registerForm = new FormGroup({
      fullname: this.form.control('', [Validators.required]),
      email: this.form.control('', [Validators.required, Validators.email]),
      password_confirmation: this.form.control('', [
        Validators.required,
        Validators.minLength(8),
      ]),
      password: this.form.control('', [
        Validators.required,
        Validators.minLength(8),
      ]),
      phone: this.form.control('+591', [Validators.required]),
      ci: this.form.control('', [Validators.required]),
      birthday: this.form.control('1990-01-01'),
    });
  }

  onSubmit(): void {
    if (
      this.registerForm.get('password')?.value !==
      this.registerForm.get('password_confirmation')?.value
    ) {
      this._snackBar.open('Las contraseñas no coinciden', 'Aceptar', {
        duration: 2000,
      });
      return;
    }

    if (this.registerForm.invalid) {
      if (this.registerForm.get('fullname')?.invalid) {
        this._snackBar.open('El campo nombre es requerido', 'Aceptar', {
          duration: 2000,
        });
        return;
      }
      if (this.registerForm.get('email')?.invalid) {
        this._snackBar.open('El campo email es requerido', 'Aceptar', {
          duration: 2000,
        });
      }
      if (this.registerForm.get('password')?.invalid) {
        this._snackBar.open(
          'El campo contraseña es requerido o Minimo 8 Caracteres',
          'Aceptar',
          {
            duration: 2000,
          }
        );
        return;
      }
      if (this.registerForm.get('password_confirmation')?.invalid) {
        this._snackBar.open(
          'El campo confirmar contraseña es requerido',
          'Aceptar',
          {
            duration: 2000,
          }
        );
        return;
      }
      if (this.registerForm.get('phone')?.invalid) {
        this._snackBar.open('El campo telefono es requerido', 'Aceptar', {
          duration: 2000,
        });
        return;
      }
      if (this.registerForm.get('ci')?.invalid) {
        this._snackBar.open('El campo ci es requerido', 'Aceptar', {
          duration: 2000,
        });
        return;
      }
    }
    if (this.registerForm.valid) {
      this.dataService
        .register(this.registerForm.value)
        .subscribe((response) => {
          if (response) {
            console.log(response);
            // if (response.error) {
            //   this._snackBar.open(response.error, 'Aceptar', {
            //     duration: 10000,
            //     horizontalPosition: 'center',
            //     verticalPosition: 'top',
            //     panelClass: ['error-snackbar', 'snackbar-center'],
            //   });
            // }
          }
        });
    }
  }
}

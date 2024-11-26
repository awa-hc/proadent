import { AbstractControl, ValidationErrors, ValidatorFn } from '@angular/forms';

export function noPastDateValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const today = new Date();
    today.setHours(0, 0, 0, 0); // Establecemos la hora a medianoche para comparar solo la fecha

    const selectedDate = new Date(control.value);
    if (selectedDate < today) {
      return { pastDate: true }; // Si la fecha es anterior, retornamos un error
    }
    return null; // Si la fecha es válida, no hay error
  };
}

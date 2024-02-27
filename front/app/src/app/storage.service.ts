import { isPlatformBrowser } from '@angular/common';
import { Inject, Injectable, PLATFORM_ID } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class StorageService {
  private isBrowser: boolean;

  constructor(@Inject(PLATFORM_ID) private platformId: Object) {
    this.isBrowser = isPlatformBrowser(platformId);
  }

  setItem(key: string, value: any): void {
    if (this.isBrowser) {
      localStorage.setItem(key, value);
    }
    // Opcional: manejar almacenamiento en el servidor
  }

  getItem(key: string): any {
    if (this.isBrowser) {
      const item = localStorage.getItem(key);
      return item ? item : null;
    }
    return null; // O manejar de otra manera
  }
  removeItem(key: string): void {
    if (this.isBrowser) {
      localStorage.removeItem(key);
    }
    // Opcional: manejar almacenamiento en el servidor
  }
  clear(): void {
    if (this.isBrowser) {
      localStorage.clear();
    }
  }
}

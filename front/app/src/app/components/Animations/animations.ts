import { trigger, transition, style, animate } from '@angular/animations';

export const fadeInOutAnimation = trigger('fadeInOut', [
  transition(':enter', [
    style({ opacity: 0 }),
    animate('0.5s', style({ opacity: 1 })),
  ]),
  transition(':leave', [animate('0.5s', style({ opacity: 0 }))]),
]);

export const translateXAnimation = trigger(`translateX`, [
  transition(':enter', [
    style({ transform: 'translateX(-100%)' }), // Posición inicial del elemento
    animate('200ms ease-out', style({ transform: 'translateX(0)' })), // Cambio a la posición final
  ]),
  transition(':leave', [
    style({ transform: 'translateX(0)' }), // Posición inicial del elemento al salir
    animate('200ms ease-in', style({ transform: 'translateX(-100%)' })), // Cambio a la posición final al salir
  ]),
]);

export const translateYAnimation = trigger(`translateY`, [
  transition(':enter', [
    style({ transform: 'translateY(-100%)' }), // Posición inicial del elemento
    animate('200ms ease-out', style({ transform: 'translateY(0)' })), // Cambio a la posición final
  ]),
  transition(':leave', [
    style({ transform: 'translateY(0)' }), // Posición inicial del elemento al salir
    animate('200ms ease-in', style({ transform: 'translateY(-100%)' })), // Cambio a la posición final al salir
  ]),
]);

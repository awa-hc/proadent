import { Routes } from '@angular/router';
import { AppComponent } from './app.component';
import { LoginComponent } from './pages/login/login.component';
import { HomeComponent } from './pages/home/home.component';
import { ProfileComponent } from './pages/profile/profile.component';
import { ProfileAppointmentComponent } from './pages/profile/profile-appointment/profile-appointment.component';
import { MeComponent } from './pages/profile/me/me.component';
import { ProfileClinicsComponent } from './pages/profile/profile-clinics/profile-clinics.component';

export const routes: Routes = [
  { path: '', redirectTo: '/home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'login', component: LoginComponent },
  {
    path: 'profile',
    component: ProfileComponent,
    children: [
      {
        path: '',
        redirectTo: 'me',
        pathMatch: 'full',
      },
      {
        path: 'me',
        component: MeComponent,
      },
      {
        path: 'appointment',
        component: ProfileAppointmentComponent,
      },
      {
        path: 'clinics',
        component: ProfileClinicsComponent,
      },
      {
        path: '**',
        redirectTo: '/profile/me',
      },
    ],
  },
  { path: '**', redirectTo: '/home' },
];

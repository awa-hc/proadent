import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { LoginComponent } from './pages/login/login.component';
import { RegisterComponent } from './pages/register/register.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { AppointmentsComponent } from './dashboard/appointments/appointments.component';
import { CreateComponent as CreateAppointment } from './dashboard/appointments/create/create.component';
import { EditComponent as EditAppointment } from './dashboard/appointments/edit/edit.component';
import { ConfirmedComponent as ConfirmedAppointment } from './dashboard/appointments/confirmed/confirmed.component';
import { CompletedComponent } from './dashboard/appointments/completed/completed.component';
import { CancelledComponent } from './dashboard/appointments/cancelled/cancelled.component';
import { AllComponent as AllAppointment } from './dashboard/appointments/all/all.component';
import { PendingComponent as PendingAppointment } from './dashboard/appointments/pending/pending.component';
import { PriceComponent } from './dashboard/price/price.component';
import { CreateComponent as CreatePrice } from './dashboard/price/create/create.component';
import { EditComponent as EditPrice } from './dashboard/price/edit/edit.component';
import { ConfirmedComponent as ConfirmedPrice } from './dashboard/price/confirmed/confirmed.component';
import { PendingComponent as PendingPrice } from './dashboard/price/pending/pending.component';
import { ConfigComponent } from './dashboard/config/config.component';
import { RolesComponent } from './dashboard/config/roles/roles.component';
import { UsersComponent } from './dashboard/config/users/users.component';
import { AccountreciavableComponent } from './dashboard/accountreciavable/accountreciavable.component';
import { AllComponent as AllAccountreciavable } from './dashboard/accountreciavable/all/all.component';

export const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
  },
  {
    path: 'login',
    component: LoginComponent,
  },
  {
    path: 'register',
    component: RegisterComponent,
  },
  {
    path: 'dashboard',
    component: DashboardComponent,
    children: [
      {
        path: 'appointments',
        component: AppointmentsComponent,
        children: [
          {
            path: 'all',
            component: AllAppointment,
          },
          {
            path: 'cancelled',
            component: CancelledComponent,
          },
          {
            path: 'completed',
            component: CompletedComponent,
          },
          {
            path: 'confirmed',
            component: ConfirmedAppointment,
          },
          {
            path: 'create',
            component: CreateAppointment,
          },
          {
            path: 'edit/:id',
            component: EditAppointment,
          },
          {
            path: 'pending',
            component: PendingAppointment,
          },
          {
            path: '',
            redirectTo: 'create',
            pathMatch: 'full',
          },
        ],
      },
      {
        path: 'price',
        component: PriceComponent,
        children: [
          {
            path: 'create',
            component: CreatePrice,
          },
          {
            path: 'edit',
            component: EditPrice,
          },
          {
            path: 'confirmed',
            component: ConfirmedPrice,
          },
          {
            path: 'pending',
            component: PendingPrice,
          },
          {
            path: '',
            redirectTo: 'create',
            pathMatch: 'full',
          },
        ],
      },
      {
        path: 'config',
        component: ConfigComponent,
        children: [
          {
            path: 'roles',
            component: RolesComponent,
          },
          {
            path: 'users',
            component: UsersComponent,
          },
        ],
      },
      {
        path: 'accountreciavable',
        component: AccountreciavableComponent,
        children: [
          {
            path: 'all',
            component: AllAccountreciavable,
          },
        ],
      },
    ],
  },
];

import { HttpClient } from '@angular/common/http';
import { AfterViewInit, Component, inject, ViewChild } from '@angular/core';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { StorageService } from '../../../storage.service';
import { format, parseISO } from 'date-fns';
import { formatUtcToLocal } from '../../../utils/date-utils';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { ConfigService } from '../../../config.service';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { ProfileClinicsDialogComponent } from './dialog/profile-clinics-dialog';

/**
 * @title Table with pagination
 */
@Component({
  selector: 'app-profile-clinics',
  standalone: true,
  imports: [
    MatTableModule,
    MatPaginatorModule,
    MatButtonModule,
    MatIconModule,
    MatDialogModule,
  ],
  templateUrl: './profile-clinics.component.html',
  styleUrl: './profile-clinics.component.css',
})
export class ProfileClinicsComponent implements AfterViewInit {
  displayedColumns: string[] = [
    'code',
    'date_time',
    'status',
    'price',
    'actions',
  ];
  dataSource = new MatTableDataSource<Clinics>();

  readonly dialog = inject(MatDialog);

  constructor(
    private config: ConfigService,
    private http: HttpClient,
    private storage: StorageService
  ) {}
  @ViewChild(MatPaginator)
  paginator!: MatPaginator;

  click() {
    console.log('click');
  }

  getClinics() {
    this.http
      .get<Clinics[]>(
        'http://localhost:8080/clinic/patient/' + this.storage.getItem('ci'),
        { withCredentials: true }
      )
      .subscribe((data: Clinics[]) => {
        const formattedData = data.map((clinic) => {
          return {
            ...clinic,
            date_time: formatUtcToLocal(clinic.date_time, 'yyyy-MM-dd HH:mm'), // Asigna la fecha formateada
          };
        });

        this.dataSource = new MatTableDataSource(formattedData);
        this.dataSource.paginator = this.paginator;
        console.log('Formatted Data:', formattedData); // Verifica todo el objeto
      });
  }

  code: string | null = null;

  openClinic(code: string | null) {
    if (!code) {
      console.error('El código de clínica es inválido:', code);
      return;
    }

    const url = `${this.config.apiUrl}clinic/code/${code}`;

    this.http.get(url).subscribe(
      (data: any) => {
        console.log('Clinic:', data);

        // Verificar si los datos son un array o un objeto y formatear las fechas
        let formattedData: any;
        if (Array.isArray(data)) {
          formattedData = data.map(
            (clinic: { date_time: string; status: string; price: any }) => ({
              ...clinic,
              date_time: formatUtcToLocal(clinic.date_time, 'yyyy-MM-dd HH:mm'), // Formatear fecha
              price: clinic.price,
            })
          );
        } else {
          formattedData = {
            ...data,
            date_time: formatUtcToLocal(data.date_time, 'yyyy-MM-dd HH:mm'), // Formatear fecha
            price: data.price,
          };
        }

        // Abrir el diálogo con los datos formateados
        const dialogRef = this.dialog.open(ProfileClinicsDialogComponent, {
          data: formattedData,
        });

        dialogRef.afterClosed().subscribe((result) => {
          console.log('Dialog closed:', result);
        });
      },
      (error) => {
        console.error('Error al obtener clínica:', error);
      }
    );

    console.log('Open Clinic:', code);
  }

  parsedDate(date: string): string {
    const parsedTime = parseISO(date);
    return format(parsedTime, 'yyyy-MM-dd HH:mm');
  }

  ngAfterViewInit() {
    this.getClinics();
    this.dataSource.paginator = this.paginator;
  }
}

export interface Clinics {
  code: string;
  date_time: string;
  status: string;
  price: number;
}

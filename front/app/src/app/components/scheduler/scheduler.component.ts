import { isPlatformBrowser } from '@angular/common';
import {
  Component,
  Input,
  AfterViewInit,
  OnDestroy,
  Inject,
  PLATFORM_ID,
  OnChanges,
  SimpleChanges,
  Output,
  EventEmitter,
} from '@angular/core';
import { parseISO, format } from 'date-fns';

@Component({
  selector: 'app-scheduler',
  templateUrl: './scheduler.component.html',
  styleUrls: ['./scheduler.component.css'],
})
export class SchedulerComponent implements AfterViewInit, OnDestroy, OnChanges {
  @Input() events: any[] = []; // Recibe los eventos desde el componente padre
  @Output() emptyDateSelected = new EventEmitter<string>();

  constructor(@Inject(PLATFORM_ID) private platformId: Object) {}

  private scheduler: any;

  ngAfterViewInit(): void {
    if (isPlatformBrowser(this.platformId)) {
      this.initScheduler();
    }
  }
  ngOnChanges(changes: SimpleChanges): void {
    if (changes['events'] && !changes['events'].firstChange) {
      console.log('changes', changes);
      this.updateScheduler();
    }
  }

  ngOnDestroy(): void {
    if (this.scheduler) {
      this.scheduler.clearAll();
      this.scheduler = null;
    }
  }

  initScheduler(): void {
    if (typeof window !== 'undefined') {
      const scheduler = (window as any).scheduler;

      if (!scheduler) {
        console.log('No se ha cargado la libreria scheduler');
      }
      scheduler.init('scheduler_here', new Date(), 'week');

      scheduler.config.first_hour = 9;
      scheduler.config.last_hour = 20;
      scheduler.config.hour_date = '%H:%i';
      scheduler.config.time_step = 30;

      scheduler.config.details_on_create = false;
      scheduler.config.details_on_dblclick = false;
      scheduler.config.drag_resize = false; // Desactiva redimensionamiento por arrastre
      scheduler.config.drag_move = false; // Desactiva movimiento por arrastre
      scheduler.config.disable_drag = true; // Desactiva arrastrar eventos
      scheduler.config.readonly = true; // Hace que el scheduler sea solo lectura

      // Sobreescribir el método para evitar mostrar el cuadro de diálogo
      scheduler.showLightbox = () => {}; // Evita la apertura del lightbox predeterminado
      scheduler.hideLightbox = () => {}; // Evita que se oculte el lightbox

      scheduler.clearAll();
      this.updateScheduler();

      scheduler.attachEvent('onEmptyClick', (date: any, e: any) => {
        this.handleEmptyDateClick(date);
      });
    }
  }

  updateScheduler(): void {
    const scheduler = (window as any).scheduler;
    if (scheduler) {
      scheduler.clearAll();
      scheduler.parse(this.events, 'json');
    }
  }

  handleEmptyDateClick(date: Date): void {
    const formattedDate = format(date, 'yyyy-MM-dd HH:mm');
    this.emptyDateSelected.emit(formattedDate); // Emite la fecha seleccionada
  }
}

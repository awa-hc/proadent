// import { isPlatformBrowser } from '@angular/common';
// import {
//   Component,
//   Input,
//   AfterViewInit,
//   OnDestroy,
//   Inject,
//   PLATFORM_ID,
//   OnChanges,
//   SimpleChanges,
// } from '@angular/core';
// import { parseISO, format } from 'date-fns';

// @Component({
//   selector: 'app-scheduler',
//   templateUrl: './scheduler.component.html',
//   styleUrls: ['./scheduler.component.css'],
// })
// export class SchedulerComponent implements AfterViewInit, OnDestroy, OnChanges {
//   @Input() events: any[] = []; // Recibe los eventos desde el componente padre

//   constructor(@Inject(PLATFORM_ID) private platformId: Object) {}

//   private scheduler: any;

//   ngAfterViewInit(): void {
//     if (isPlatformBrowser(this.platformId)) {
//       this.initScheduler();
//     }
//   }
//   ngOnChanges(changes: SimpleChanges): void {
//     if (changes['events'] && !changes['events'].firstChange) {
//       console.log('changes', changes);
//       this.updateScheduler();
//     }
//   }

//   ngOnDestroy(): void {
//     if (this.scheduler) {
//       this.scheduler.clearAll();
//       this.scheduler = null;
//     }
//   }

//   initScheduler(): void {
//     if (typeof window !== 'undefined') {
//       const scheduler = (window as any).scheduler;

//       if (!scheduler) {
//         console.log('No se ha cargado la libreria scheduler');
//       }

//       scheduler.init('scheduler_here', new Date(), 'month');
//       scheduler.clearAll();
//       this.updateScheduler();
//     }
//   }

//   updateScheduler(): void {
//     const scheduler = (window as any).scheduler;
//     if (scheduler) {
//       scheduler.clearAll();
//       scheduler.parse(this.events, 'json');
//     }
//   }
// }

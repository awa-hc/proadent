import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProfileAppointmentComponent } from './profile-appointment.component';

describe('ProfileAppointmentComponent', () => {
  let component: ProfileAppointmentComponent;
  let fixture: ComponentFixture<ProfileAppointmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProfileAppointmentComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ProfileAppointmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

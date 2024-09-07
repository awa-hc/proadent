import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProfileClinicsComponent } from './profile-clinics.component';

describe('ProfileClinicsComponent', () => {
  let component: ProfileClinicsComponent;
  let fixture: ComponentFixture<ProfileClinicsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProfileClinicsComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ProfileClinicsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

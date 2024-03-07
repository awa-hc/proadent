import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AccountreciavableComponent } from './accountreciavable.component';

describe('AccountreciavableComponent', () => {
  let component: AccountreciavableComponent;
  let fixture: ComponentFixture<AccountreciavableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AccountreciavableComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(AccountreciavableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

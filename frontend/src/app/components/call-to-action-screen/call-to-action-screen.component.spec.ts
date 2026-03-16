import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CallToActionScreenComponent } from './call-to-action-screen.component';

describe('CallToActionScreenComponent', () => {
  let component: CallToActionScreenComponent;
  let fixture: ComponentFixture<CallToActionScreenComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CallToActionScreenComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CallToActionScreenComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

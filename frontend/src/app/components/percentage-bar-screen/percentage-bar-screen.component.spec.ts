import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PercentageBarScreenComponent } from './percentage-bar-screen.component';

describe('PercentageBarScreenComponent', () => {
  let component: PercentageBarScreenComponent;
  let fixture: ComponentFixture<PercentageBarScreenComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PercentageBarScreenComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PercentageBarScreenComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

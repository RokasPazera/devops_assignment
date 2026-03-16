import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GoalScreenComponent } from './goal-screen.component';

describe('GoalScreenComponent', () => {
  let component: GoalScreenComponent;
  let fixture: ComponentFixture<GoalScreenComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GoalScreenComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GoalScreenComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

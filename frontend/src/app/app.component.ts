import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CallToActionScreenComponent } from './components/call-to-action-screen/call-to-action-screen.component';
import { ErrorScreenComponent } from './components/error-screen/error-screen.component';
import { PercentageBarScreenComponent } from './components/percentage-bar-screen/percentage-bar-screen.component';
import { GoalScreenComponent } from './components/goal-screen/goal-screen.component';
import { interval, Subscription } from 'rxjs';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,
    CallToActionScreenComponent,
    ErrorScreenComponent,
    PercentageBarScreenComponent,
    GoalScreenComponent
  ],
  templateUrl: './app.component.html'
})
export class AppComponent implements OnInit {
  currentScreen = 0;
  percentage = 0;
  goals: { name: string; targetPercentage: number }[] = [];
  error = false;
  SLIDE_DURATION_MS = 8000;

  private slideSubscription?: Subscription;

  ngOnInit(): void {
    this.loadFromAPI();
  }

  async loadFromAPI() {
    try {
      const res = await fetch('http://localhost:3000/api/percentage');
      const data = await res.json();
      this.percentage = data.percentage;

      const resGoals = await fetch('http://localhost:3000/api/goals');
      this.goals = await resGoals.json();

      this.error = false;
      this.startSlideShow();
    } catch (err) {
      console.error('Error fetching API data');
      this.error = true;
    }
  }

  startSlideShow() {
    if (this.slideSubscription) this.slideSubscription.unsubscribe();

    this.slideSubscription = interval(this.SLIDE_DURATION_MS).subscribe(() => {
      this.currentScreen = (this.currentScreen + 1) % 3;
    });
  }
}

import { Component, Input, AfterViewInit, ViewChild, ViewChildren, QueryList, ElementRef, OnChanges, SimpleChanges } from '@angular/core';
import { CommonModule } from '@angular/common';
import { gsap } from 'gsap';

@Component({
  selector: 'app-goal-screen',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './goal-screen.component.html'
})
export class GoalScreenComponent implements AfterViewInit, OnChanges {
  @Input() percentage = 0;
  @Input() goals: { name: string; targetPercentage: number }[] = [];

  @ViewChild('title') title!: ElementRef<HTMLHeadingElement>;
  @ViewChild('goalList') goalList!: ElementRef<HTMLDivElement>;
  @ViewChildren('cards') cards!: QueryList<ElementRef<HTMLDivElement>>;

  private tl?: gsap.core.Timeline;

  ngAfterViewInit() {
    this.createAnimation();
    this.tl?.play();
  }

  ngOnChanges(changes: SimpleChanges) {
    if ((changes['percentage'] || changes['goals']) && !changes['percentage']?.firstChange) {
      if (this.tl) this.tl.kill();
      this.createAnimation();
      this.tl?.play();
    }
  }

  private createAnimation() {
    const cardElements = this.cards.toArray().map(c => c.nativeElement);

    this.tl = gsap.timeline({ defaults: { ease: 'power2.out' }, paused: true });

    // Animate title
    this.tl.fromTo(
      this.title.nativeElement,
      { opacity: 0, y: 50 },
      { opacity: 1, y: 0, duration: 1 }
    );

    // Animate cards
    if (cardElements.length > 0) {
      this.tl.fromTo(
        cardElements,
        { opacity: 0, y: 50 },
        { opacity: 1, y: 0, stagger: 0.2, duration: 0.3 }
      );
    } else {
      console.warn('No goals were found, so card animation will not run!');
    }
  }
}

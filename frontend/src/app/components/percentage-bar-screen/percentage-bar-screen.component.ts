import { Component, Input, OnChanges, SimpleChanges, ElementRef, ViewChild, AfterViewInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { gsap } from 'gsap';

@Component({
  selector: 'app-percentage-bar-screen',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './percentage-bar-screen.component.html'
})
export class PercentageBarScreenComponent implements AfterViewInit, OnChanges {
  @Input() percentage = 0;

  @ViewChild('bar') bar!: ElementRef<HTMLDivElement>;
  @ViewChild('barText') barText!: ElementRef<HTMLSpanElement>;
  @ViewChild('title') title!: ElementRef<HTMLHeadingElement>;

  private tl?: gsap.core.Timeline;

  ngAfterViewInit() {
    this.createAnimation();
    this.tl?.play();
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes['percentage'] && !changes['percentage'].firstChange) {
      if (this.tl) this.tl.kill();
      this.createAnimation();
      this.tl?.play();
    }
  }

  private createAnimation() {
    this.tl = gsap.timeline({ defaults: { ease: 'power2.out' }, paused: true });

    // Animate heading
    this.tl.fromTo(
      this.title.nativeElement,
      { opacity: 0, y: 50 },
      { opacity: 1, y: 0, duration: 1 }
    );

    // Animate bar width
    this.tl.fromTo(
      this.bar.nativeElement,
      { width: '0%' },
      {
        width: this.percentage + '%',
        duration: 2,
        delay: 0.5,
        onUpdate: () => {
          this.barText.nativeElement.innerText = Math.round(this.tl!.progress() * this.percentage) + '%';
        }
      },
      '-=0.5'
    );
  }
}

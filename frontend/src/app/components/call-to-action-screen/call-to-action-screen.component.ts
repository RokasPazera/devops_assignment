import { Component, ElementRef, AfterViewInit, ViewChild } from '@angular/core';
import { gsap } from 'gsap';

@Component({
  selector: 'app-call-to-action-screen',
  templateUrl: './call-to-action-screen.component.html'
})
export class CallToActionScreenComponent implements AfterViewInit {
  @ViewChild('screen') screen!: ElementRef<HTMLElement>;
  titles: HTMLElement[] = [];

  ngAfterViewInit() {
    // Collect all h1 elements
    this.titles = Array.from(this.screen.nativeElement.querySelectorAll('h1'));
    this.createAnimation().play();
  }

  createAnimation() {
    const tl = gsap.timeline({ defaults: { duration: 1, ease: 'power2.out' }, paused: true });
    tl.fromTo(this.titles, { opacity: 0, y: 50 }, { opacity: 1, y: 0, stagger: 0.5 });
    return {
      play: () => tl.play(),
      reset: () => tl.pause(0)
    };
  }
}

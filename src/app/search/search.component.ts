import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent {
  search(query: string) {
    
  }

  number: number = 1;

  increment() {
    this.number++;
  }

  decrement() {
    if (this.number > 1) {
      this.number--;
    }
  }

  min = 0;
  max = 100;
  step = 10;
  value = 50;
  ticks = [];

  constructor() {
    this.generateTicks();
  }

  generateTicks() {
    const numTicks = (this.max - this.min) / this.step + 1;
    const tickSize = 100 / (numTicks - 1);
    for (let i = 0; i < numTicks; i++) {
      const tickValue = this.min + i * this.step;
      const tickLeft = i === 0 ? 0 : (i === numTicks - 1 ? 100 : i * tickSize);
    }
  }

  onChange(event: any) {
    console.log(this.value);
  }

  breakfastClicked = false;
  lunchClicked = false;
  dinnerClicked = false;
  snackClicked = false;
  dessertClicked = false;

  mealType(buttonId: string) {
    const button = document.getElementById(buttonId);
    if (button) {
      if (button.classList.contains("clicked")) {
        button.classList.remove("clicked");
        switch (buttonId) {
          case "Breakfast":
            this.breakfastClicked = false;
            break;
          case "Lunch":
            this.lunchClicked = false;
            break;
          case "Dinner":
            this.dinnerClicked = false;
            break;
          case "Snack":
            this.snackClicked = false;
            break;
          case "Dessert":
            this.dessertClicked = false;
            break;
        }
      } else {
        button.classList.add("clicked");
        switch (buttonId) {
          case "Breakfast":
            this.breakfastClicked = true;
            break;
          case "Lunch":
            this.lunchClicked = true;
            break;
          case "Dinner":
            this.dinnerClicked = true;
            break;
          case "Snack":
            this.snackClicked = true;
            break;
          case "Dessert":
            this.dessertClicked = true;
            break;
        }
      }
    }
  }

  ngOnInit() {
    const slider = document.querySelector('.slider') as HTMLElement;
    const range = slider.querySelector('.slider-range') as HTMLInputElement;
    const tickMarks = slider.querySelectorAll('.tick-mark') as NodeListOf<HTMLElement>;

    range.addEventListener('input', () => {
      const value = parseInt(range.value);
      const position = value * slider.clientWidth;
      const ticks = Array.from(tickMarks);

      ticks.forEach((tick) => {
        tick.classList.toggle('active', parseInt(tick.innerText) === value);
      });

      range.style.setProperty('--thumb-position', `${position}px`);
    });

    range.dispatchEvent(new Event('input'));
  }

  OntheGoClicked = false;
  MinClicked = false;
  HrClicked = false;
  HrsClicked = false;

  mealTime(buttonId: string) {
    const button = document.getElementById(buttonId);
    if (button) {
      if (button.classList.contains("clicked")) {
        button.classList.remove("clicked");
        switch (buttonId) {
          case "On the Go":
            this.OntheGoClicked = false;
            break;
          case "~30 min":
            this.MinClicked = false;
            break;
          case "~1 Hr":
            this.HrClicked = false;
            break;
          case "~2 Hrs":
            this.HrsClicked = false;
            break;
        }
      } else {
        button.classList.add("clicked");
        switch (buttonId) {
          case "On the Go":
            this.OntheGoClicked = true;
            break;
          case "~30 min":
            this.MinClicked = true;
            break;
          case "~1 Hr":
            this.HrClicked = true;
            break;
          case "~2 Hrs":
            this.HrsClicked = true;
            break;
        }
      }
    }
  }

}

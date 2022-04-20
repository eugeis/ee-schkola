import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-personname',
  templateUrl: './personname.component.html',
  styleUrls: ['./personname.component.css']
})

export class PersonNameComponent implements OnInit {

  first: string
  last: string

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.first = (<HTMLInputElement>document.getElementById('first')).value;
    this.last = (<HTMLInputElement>document.getElementById('last')).value;
  }
  deleteElement() {
    this.first = '';
    this.last = '';
  }
  printElement() {
    console.log('First Value: ' + this.first)
    console.log('Last Value: ' + this.last)
  }

}




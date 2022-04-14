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

}




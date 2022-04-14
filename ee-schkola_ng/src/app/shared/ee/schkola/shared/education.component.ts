import { Component, OnInit } from '@angular/core';
import { Graduation } from '../person/PersonApiBase';

@Component({
  selector: 'app-education',
  templateUrl: './education.component.html',
  styleUrls: ['./education.component.css']
})
export class EducationComponent implements OnInit {

  graduation: Graduation
  other: string
  profession: string

  constructor() { }

  ngOnInit(): void {
  }

}




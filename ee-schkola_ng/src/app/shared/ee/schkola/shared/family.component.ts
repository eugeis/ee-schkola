import { Component, OnInit } from '@angular/core';
import { MaritalState } from '../person/PersonApiBase';
import { PersonName } from '../person/PersonApiBase';

@Component({
  selector: 'app-family',
  templateUrl: './family.component.html',
  styleUrls: ['./family.component.css']
})
export class FamilyComponent implements OnInit {

  maritalState: MaritalState
  childrenCount: number
  partner: PersonName

  constructor() { }

  ngOnInit(): void {
  }

}




import { Component, OnInit } from '@angular/core';
import { MaritalState } from '../../schkola/person/PersonApiBase';
import { PersonName } from '../../schkola/person/PersonApiBase';

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
  inputElement() {
    this.maritalState = (<HTMLInputElement>document.getElementById('maritalState')).value;
    this.childrenCount = Number((<HTMLInputElement>document.getElementById('childrenCount')).value);
    this.partner = (<HTMLInputElement>document.getElementById('partner')).value;
  }
  deleteElement() {
    this.maritalState = new MaritalState();
    this.childrenCount = 0;
    this.partner = new PersonName();
  }
  printElement() {
    console.log('MaritalState Value: ' + this.maritalState)
    console.log('ChildrenCount Value: ' + this.childrenCount)
    console.log('Partner Value: ' + this.partner)
  }

}




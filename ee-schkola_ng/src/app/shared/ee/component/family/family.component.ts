import { Component, OnInit } from '@angular/core';
import { MaritalState } from '../../schkola/person/PersonApiBase';
import { PersonName } from '../../schkola/person/PersonApiBase';

@Component({
  selector: 'app-family',
  templateUrl: './family.component.html',
  styleUrls: ['./family.component.scss']
})

export class FamilyComponent implements OnInit {

  maritalState: MaritalState
  childrenCount: Number
  partner: PersonName = new PersonName()

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.maritalState = Number((<HTMLInputElement>document.getElementById('maritalState')).value);
    this.childrenCount = Number((<HTMLInputElement>document.getElementById('childrenCount')).value);
    this.partner.first = (<HTMLInputElement>document.getElementById('partnerFirst')).value;
    this.partner.last = (<HTMLInputElement>document.getElementById('partnerLast')).value;
  }
  deleteElement() {
    this.maritalState = 0;
    this.childrenCount = 0;
    this.partner.first = '';
    this.partner.last = '';
  }
  printElement() {
    console.log('MaritalState Value: ' + this.maritalState)
    console.log('ChildrenCount Value: ' + this.childrenCount)
    console.log('PartnerFirst Value: ' + this.partner.first)
    console.log('PartnerLast Value: ' + this.partner.last)
  }

}




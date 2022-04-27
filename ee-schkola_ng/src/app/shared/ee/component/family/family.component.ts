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
  index = 0;
  dataElement: any[][] = [];
  tempArray: any[][] = [];

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.maritalState = Number((<HTMLInputElement>document.getElementById('maritalState')).value);
    this.childrenCount = Number((<HTMLInputElement>document.getElementById('childrenCount')).value);
    this.partner.first = this.simplifiedHtmlInputElement('partnerFirst');
    this.partner.last = this.simplifiedHtmlInputElement('partnerLast');
    this.dataElement.push([this.maritalState, this.childrenCount, this.partner.first, this.partner.last]);

  }
  deleteElement(index: number) {
    this.dataElement[index][0] = '';
    this.dataElement[index][1] = '';
    this.dataElement[index][2] = '';
    this.dataElement[index][3] = '';
    this.tempArray = this.dataElement.filter(function(element) {return element[0] !== '' })
    this.dataElement = this.tempArray;
  }
  printElement(index: number) {
    console.log('MaritalState Value: ' + this.dataElement[index][0])
    console.log('ChildrenCount Value: ' + this.dataElement[index][1])
    console.log('PartnerFirst Value: ' + this.dataElement[index][2])
    console.log('PartnerLast Value: ' + this.dataElement[index][3])
  }
  changeIndex(input: number) {
        this.index = input;
  }
  loadElement(index: number) {
    (<HTMLInputElement>document.getElementById('maritalState')).value = this.dataElement[index][0];
    (<HTMLInputElement>document.getElementById('childrenCount')).value = this.dataElement[index][1];
    (<HTMLInputElement>document.getElementById('partnerFirst')).value = this.dataElement[index][2];
    (<HTMLInputElement>document.getElementById('partnerLast')).value = this.dataElement[index][3];
  }
  editElement(index: number) {
    this.dataElement[index][0] = this.simplifiedHtmlInputElement('maritalState');
    this.dataElement[index][1] = this.simplifiedHtmlInputElement('childrenCount');
    this.dataElement[index][2] = this.simplifiedHtmlInputElement('partnerFirst');
    this.dataElement[index][3] = this.simplifiedHtmlInputElement('partnerLast');
  }
  simplifiedHtmlInputElement(element: string) {
        return (<HTMLInputElement>document.getElementById(element)).value;
  }

}




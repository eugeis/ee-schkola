import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-churchinfo',
  templateUrl: './churchinfo.component.html',
  styleUrls: ['./churchinfo.component.scss']
})

export class ChurchInfoComponent implements OnInit {

  church: String
  member: Boolean
  services: String
  index = 0;
  dataElement: any[][] = [];
  tempArray: any[][] = [];

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.church = this.simplifiedHtmlInputElement('church');
    this.member = Boolean((<HTMLInputElement>document.getElementById('member')).value);
    this.services = this.simplifiedHtmlInputElement('services');
    this.dataElement.push([this.church, this.member, this.services]);

  }
  deleteElement(index: number) {
    this.dataElement[index][0] = '';
    this.dataElement[index][1] = '';
    this.dataElement[index][2] = '';
    this.tempArray = this.dataElement.filter(function(element) {return element[0] !== '' })
    this.dataElement = this.tempArray;
  }
  printElement(index: number) {
    console.log('Church Value: ' + this.dataElement[index][0])
    console.log('Member Value: ' + this.dataElement[index][1])
    console.log('Services Value: ' + this.dataElement[index][2])
  }
  changeIndex(input: number) {
        this.index = input;
  }
  loadElement(index: number) {
    (<HTMLInputElement>document.getElementById('church')).value = this.dataElement[index][0];
    (<HTMLInputElement>document.getElementById('member')).value = this.dataElement[index][1];
    (<HTMLInputElement>document.getElementById('services')).value = this.dataElement[index][2];
  }
  editElement(index: number) {
    this.dataElement[index][0] = this.simplifiedHtmlInputElement('church');
    this.dataElement[index][1] = this.simplifiedHtmlInputElement('member');
    this.dataElement[index][2] = this.simplifiedHtmlInputElement('services');
  }
  simplifiedHtmlInputElement(element: string) {
        return (<HTMLInputElement>document.getElementById(element)).value;
  }

}




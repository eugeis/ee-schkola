import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-personname',
  templateUrl: './personname.component.html',
  styleUrls: ['./personname.component.scss']
})

export class PersonNameComponent implements OnInit {

  first: String
  last: String
  index = 0;
  dataElement: any[][] = [];
  tempArray: any[][] = [];

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.first = this.simplifiedHtmlInputElement('first');
    this.last = this.simplifiedHtmlInputElement('last');
    this.dataElement.push([this.first, this.last]);

  }
  deleteElement(index: number) {
    this.dataElement[index][0] = '';
    this.dataElement[index][1] = '';
    this.tempArray = this.dataElement.filter(function(element) {return element[0] !== '' })
    this.dataElement = this.tempArray;
  }
  printElement(index: number) {
    console.log('First Value: ' + this.dataElement[index][0])
    console.log('Last Value: ' + this.dataElement[index][1])
  }
  changeIndex(input: number) {
        this.index = input;
  }
  loadElement(index: number) {
    (<HTMLInputElement>document.getElementById('first')).value = this.dataElement[index][0];
    (<HTMLInputElement>document.getElementById('last')).value = this.dataElement[index][1];
  }
  editElement(index: number) {
    this.dataElement[index][0] = this.simplifiedHtmlInputElement('first');
    this.dataElement[index][1] = this.simplifiedHtmlInputElement('last');
  }
  simplifiedHtmlInputElement(element: string) {
        return (<HTMLInputElement>document.getElementById(element)).value;
  }

}




import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-contact',
  templateUrl: './contact.component.html',
  styleUrls: ['./contact.component.scss']
})

export class ContactComponent implements OnInit {

  phone: String
  email: String
  cellphone: String
  index = 0;
  dataElement: any[][] = [];
  tempArray: any[][] = [];

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.phone = this.simplifiedHtmlInputElement('phone');
    this.email = this.simplifiedHtmlInputElement('email');
    this.cellphone = this.simplifiedHtmlInputElement('cellphone');
    this.dataElement.push([this.phone, this.email, this.cellphone]);

  }
  deleteElement(index: number) {
    this.dataElement[index][0] = '';
    this.dataElement[index][1] = '';
    this.dataElement[index][2] = '';
    this.tempArray = this.dataElement.filter(function(element) {return element[0] !== '' })
    this.dataElement = this.tempArray;
  }
  printElement(index: number) {
    console.log('Phone Value: ' + this.dataElement[index][0])
    console.log('Email Value: ' + this.dataElement[index][1])
    console.log('Cellphone Value: ' + this.dataElement[index][2])
  }
  changeIndex(input: number) {
        this.index = input;
  }
  loadElement(index: number) {
    (<HTMLInputElement>document.getElementById('phone')).value = this.dataElement[index][0];
    (<HTMLInputElement>document.getElementById('email')).value = this.dataElement[index][1];
    (<HTMLInputElement>document.getElementById('cellphone')).value = this.dataElement[index][2];
  }
  editElement(index: number) {
    this.dataElement[index][0] = this.simplifiedHtmlInputElement('phone');
    this.dataElement[index][1] = this.simplifiedHtmlInputElement('email');
    this.dataElement[index][2] = this.simplifiedHtmlInputElement('cellphone');
  }
  simplifiedHtmlInputElement(element: string) {
        return (<HTMLInputElement>document.getElementById(element)).value;
  }

}




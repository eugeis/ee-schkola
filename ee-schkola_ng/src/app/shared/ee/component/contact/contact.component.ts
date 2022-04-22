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

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.phone = (<HTMLInputElement>document.getElementById('phone')).value;
    this.email = (<HTMLInputElement>document.getElementById('email')).value;
    this.cellphone = (<HTMLInputElement>document.getElementById('cellphone')).value;
  }
  deleteElement() {
    this.phone = '';
    this.email = '';
    this.cellphone = '';
  }
  printElement() {
    console.log('Phone Value: ' + this.phone)
    console.log('Email Value: ' + this.email)
    console.log('Cellphone Value: ' + this.cellphone)
  }

}




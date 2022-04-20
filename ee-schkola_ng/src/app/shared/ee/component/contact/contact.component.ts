import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-contact',
  templateUrl: './contact.component.html',
  styleUrls: ['./contact.component.css']
})

export class ContactComponent implements OnInit {

  phone: string
  email: string
  cellphone: string

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




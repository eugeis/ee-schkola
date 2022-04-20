import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-address',
  templateUrl: './address.component.html',
  styleUrls: ['./address.component.css']
})

export class AddressComponent implements OnInit {

  street: string
  suite: string
  city: string
  code: string
  country: string

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.street = (<HTMLInputElement>document.getElementById('street')).value;
    this.suite = (<HTMLInputElement>document.getElementById('suite')).value;
    this.city = (<HTMLInputElement>document.getElementById('city')).value;
    this.code = (<HTMLInputElement>document.getElementById('code')).value;
    this.country = (<HTMLInputElement>document.getElementById('country')).value;
  }
  deleteElement() {
    this.street = '';
    this.suite = '';
    this.city = '';
    this.code = '';
    this.country = '';
  }
  printElement() {
    console.log('Street Value: ' + this.street)
    console.log('Suite Value: ' + this.suite)
    console.log('City Value: ' + this.city)
    console.log('Code Value: ' + this.code)
    console.log('Country Value: ' + this.country)
  }

}




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

}




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

}




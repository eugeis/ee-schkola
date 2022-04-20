import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-churchinfo',
  templateUrl: './churchinfo.component.html',
  styleUrls: ['./churchinfo.component.css']
})

export class ChurchInfoComponent implements OnInit {

  church: string
  member: boolean
  services: string

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.church = (<HTMLInputElement>document.getElementById('church')).value;
    this.member = Boolean((<HTMLInputElement>document.getElementById('member')).value);
    this.services = (<HTMLInputElement>document.getElementById('services')).value;
  }
  deleteElement() {
    this.church = '';
    this.member = undefined;
    this.services = '';
  }
  printElement() {
    console.log('Church Value: ' + this.church)
    console.log('Member Value: ' + this.member)
    console.log('Services Value: ' + this.services)
  }

}




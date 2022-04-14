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

}




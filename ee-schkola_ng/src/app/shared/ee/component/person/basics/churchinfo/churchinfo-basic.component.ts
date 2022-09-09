import {ChurchInfo} from '../person/PersonApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-churchinfo',
  templateUrl: './churchinfo.component.html',
  styleUrls: ['./churchinfo.component.scss'],
})

export class ChurchInfoComponent implements OnInit {

    @Input() churchinfo: ChurchInfo = new ChurchInfo();

    ngOnInit(): void {
        if (this.churchinfo === undefined) {
            this.churchinfo = {church: '', member: false, services: ''};
        }
    }
}





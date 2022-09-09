import {PersonName} from '../person/PersonApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-personname',
  templateUrl: './personname.component.html',
  styleUrls: ['./personname.component.scss'],
})

export class PersonNameComponent implements OnInit {

    @Input() personname: PersonName = new PersonName();

    ngOnInit(): void {
        if (this.personname === undefined) {
            this.personname = {first: '', last: ''};
        }
    }
}





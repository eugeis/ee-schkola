import {PersonName} from '@schkola/person/PersonApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-personname',
  templateUrl: './personname-basic.component.html',
  styleUrls: ['./personname-basic.component.scss'],
})

export class PersonNameComponent implements OnInit {

    @Input() personname: PersonName;
    
    ngOnInit() {
        if (this.personname === undefined) {
            this.personname = new PersonName();
        }
        
    }
    
}





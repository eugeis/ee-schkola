import {Family, PersonName} from '@schkola/person/PersonApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-family',
  templateUrl: './family-basic.component.html',
  styleUrls: ['./family-basic.component.scss'],
})

export class FamilyComponent implements OnInit {

    @Input() family: Family;
    
    ngOnInit() {
        if (this.family === undefined) {
            this.family = new Family();
        }
        if (this.family.partner === undefined) {
            this.family.partner = new PersonName();
        }
    }
    
}





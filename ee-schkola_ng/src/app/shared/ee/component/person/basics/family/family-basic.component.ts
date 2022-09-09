import {Family} from '../person/PersonApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-family',
  templateUrl: './family.component.html',
  styleUrls: ['./family.component.scss'],
})

export class FamilyComponent implements OnInit {

    @Input() family: Family = new Family();

    ngOnInit(): void {
        if (this.family === undefined) {
            this.family = {maritalstate: 0, childrencount: 0, partner: {first: '', last: ''}};
        }
    }
}





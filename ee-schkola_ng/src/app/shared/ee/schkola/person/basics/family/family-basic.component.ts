import {Family} from '@schkola/person/PersonApiBase'

import {Component, Input} from '@angular/core';

@Component({
  selector: 'app-family',
  templateUrl: './family-basic.component.html',
  styleUrls: ['./family-basic.component.scss'],
})

export class FamilyComponent {

    @Input() family: Family;
    
}





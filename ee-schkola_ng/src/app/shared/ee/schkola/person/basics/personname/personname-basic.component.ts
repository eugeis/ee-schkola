import {PersonName} from '@schkola/person/PersonApiBase'

import {Component, Input} from '@angular/core';

@Component({
  selector: 'app-personname',
  templateUrl: './personname-basic.component.html',
  styleUrls: ['./personname-basic.component.scss'],
})

export class PersonNameComponent {

    @Input() personname: PersonName;
    
}




